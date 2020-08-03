package fileutil

//一个简化版的文件系统watcher

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// WatcherEventType 事件类型
type WatcherEventType string

const (
	//WatcherEventCreate 新创建文件
	WatcherEventCreate WatcherEventType = "create"

	//WatcherEventChange 文件内容、名称、属性等变化
	WatcherEventChange = "change"

	//WatcherEventDelete 文件被删除，包括名字修改，路径调整
	//如 文件名 由  a 调整 为b，会产生2个event，分别为：a-删除，b-创建
	WatcherEventDelete = "delete"
)

// WatcherEvent 对外暴露的event对象
type WatcherEvent struct {
	Type WatcherEventType

	//文件名,相对路径,相对于rootpath的文件名，eg: ral/services/a.toml
	Name string

	//文件的完整路径名，绝对路径,eg: /home/work/appName/conf/ral/services/a.toml
	Path string
}

// String 序列化
func (e *WatcherEvent) String() string {
	return fmt.Sprintf("file=%s, event=%s", e.Name, e.Type)
}

// WatcherEventHandler 注册的回调函数类型
type WatcherEventHandler func(event *WatcherEvent) error

type handleType struct {
	handleFn WatcherEventHandler
	index    uint64
}

// Watcher watcher对象
type Watcher struct {
	rootDir  string
	fwatcher *fsnotify.Watcher

	//注册的对应路径和回调的数据 eg：{"abc/*":func(e *Event)error{},"def/e/*":func(e *Event)error{}}
	handles      map[string][]*handleType
	handlesTotal uint64
	handlesMax   uint64

	rw sync.RWMutex

	eventFired int64
	eventDone  int64
}

// error4log 序列化err对象，以方便打印日志
func error4log(err error) string {
	if err == nil {
		return "success"
	}
	return fmt.Sprintf("faild, err=%s", err)
}

// NewWatcher 创建一个新的watcher对象
// rootDir为待watch的根目录
func NewWatcher(rootDir string, handlesMax uint64) *Watcher {
	w := &Watcher{
		rootDir:    rootDir,
		handles:    make(map[string][]*handleType),
		handlesMax: handlesMax,
	}

	w.watchRoot()

	return w
}

// watchRoot 对根目录以及子目录进行watch
func (w *Watcher) watchRoot() error {
	if w.fwatcher != nil {
		w.fwatcher.Close()
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	w.fwatcher = watcher
	for {
		//若对应的文件目录此时不存在，则等待
		//比如将当前目录rm掉了，copy 或者mv 了一个目录过来
		if Exists(w.rootDir) {
			err := w.fwatcher.Add(w.rootDir)
			log.Printf("[fileutil_watcher] root_dir added watch %s, root_dir= %s\n", error4log(err), w.rootDir)
			if err != nil {
				return err
			}
			break
		}
		log.Printf("[fileutil_watcher] root_dir not exists, waiting for add watch, root_dir= %s\n", w.rootDir)
		time.Sleep(1 * time.Second)
	}

	go func() {
		for {
			select {
			case event := <-w.fwatcher.Events:
				go w.fireEvent(event)
			case err := <-w.fwatcher.Errors:
				log.Printf("[fileutil_watcher] file watcher catch err: %s", err.Error())
			}
		}
	}()

	//子目录必须逐个监听才可以
	err = w.addWatchDir(w.rootDir)

	return err
}

func (w *Watcher) addWatchDir(dir string) error {
	return filepath.Walk(dir, func(fpath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		return w.fwatcher.Add(fpath)
	})
}

func (w *Watcher) fireEvent(event fsnotify.Event) {
	fname := filepath.Clean(event.Name)
	relName, _ := filepath.Rel(w.rootDir, fname)

	myEvent := &WatcherEvent{
		Name: relName,
		Path: fname,
	}

	//为了简化使用，将文件的内容变化、属性变化统一归一为change事件

	isRemove := event.Op&fsnotify.Remove == fsnotify.Remove

	if event.Op&fsnotify.Create == fsnotify.Create {
		myEvent.Type = WatcherEventCreate
	} else if event.Op&fsnotify.Write == fsnotify.Write {
		myEvent.Type = WatcherEventChange
	} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
		myEvent.Type = WatcherEventChange
	} else if isRemove {
		myEvent.Type = WatcherEventDelete
	} else if event.Op&fsnotify.Rename == fsnotify.Rename {
		myEvent.Type = WatcherEventDelete
	} else {
		log.Printf("[fileutil_watcher] other event ignore: %s\n", event)
		return
	}
	var err error

	switch myEvent.Type {
	case WatcherEventDelete:
		//目录被删除了，或者mv到其他地方、rename成别的名字了
		if fname == w.rootDir {
			log.Printf("[fileutil_watcher] root dir was deleted, path=%s\n", fname)
			w.watchRoot()
		} else if !isRemove {
			err = w.fwatcher.Remove(fname)
			log.Printf("[fileutil_watcher] file was deleted, path= %s, remove watch %s\n", fname, error4log(err))
		}
	case WatcherEventCreate:
		//新创建的文件必须再次加入
		err = w.fwatcher.Add(fname)
		log.Printf("[fileutil_watcher] a new file was founded, path= %s, addwatch it %s\n", fname, error4log(err))
	}

	w.rw.RLock()
	defer w.rw.RUnlock()

	var handlesFired []*handleType

	for pattern, handles := range w.handles {
		if w.isMatch(pattern, relName) {
			for _, handler := range handles {
				handlesFired = append(handlesFired, handler)
			}
		}
	}

	if len(handlesFired) > 0 {
		//对注册的hander排序后执行，保证先注册的一定能先执行
		sort.Slice(handlesFired, func(i, j int) bool {
			return handlesFired[i].index > handlesFired[j].index
		})

		for _, handles := range handlesFired {
			go handles.handleFn(myEvent)
		}
	}

	log.Printf("[fileutil_watcher] fired file event ( %v ): handles total= %d\n", myEvent, len(handlesFired))
}

func (w *Watcher) isMatch(pattern string, fpath string) bool {
	if pattern == fpath {
		return true
	}
	if len(pattern) > len(fpath) {
		return false
	}

	if strings.HasSuffix(pattern, "*") {
		patternPre := strings.TrimRight(pattern, "*")
		if strings.HasPrefix(fpath, patternPre) {
			return true
		}
	}
	match, _ := filepath.Match(pattern, fpath)
	return match
}

// Close 关闭watcher
func (w *Watcher) Close() error {
	if w.fwatcher != nil {
		return w.fwatcher.Close()
	}
	return nil
}

// RegisterFileWatcher 注册文件变更事件观察者
// pattern 是可以包括  * 的路径，如  db/*.toml 或者 db/demo.toml
func (w *Watcher) RegisterFileWatcher(pattern string, handleFn WatcherEventHandler) (err error) {
	if handleFn == nil {
		return fmt.Errorf("regirster handleFn is nil")
	}

	//用于兼容输入参数,忽略第一个/
	pattern = strings.TrimLeft(pattern, "/")

	//测试一下pattern是否是满足要求的
	_, err = filepath.Match(pattern, "test")
	if err != nil {
		return err
	}

	w.rw.Lock()
	defer w.rw.Unlock()
	_, has := w.handles[pattern]
	if !has {
		w.handles[pattern] = make([]*handleType, 0, 5)
	}
	w.handlesTotal++

	//正常情况下不应该出现
	if w.handlesTotal >= w.handlesMax {
		panic(fmt.Sprintf("too many filewatcher handlers (%d)", w.handlesMax))
	}

	handler := &handleType{
		handleFn: handleFn,
		index:    w.handlesTotal,
	}
	w.handles[pattern] = append(w.handles[pattern], handler)
	return nil
}
