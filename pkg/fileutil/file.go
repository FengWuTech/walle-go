package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// FileAppend just alias
	FileAppend = os.O_APPEND
)


// Exists 判断文件或目录是否存在
func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsDir(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}


// FileGetContents 从文件读取内容
func FileGetContents(filePath string) (data []byte, err error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	bf, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		return nil, err1
	}
	return bf, nil
}

// FilePutContents 往文件写入内容
func FilePutContents(filePath string, data []byte, def ...int) error {
	flags := os.O_RDWR | os.O_CREATE
	isAppend := false
	if len(def) > 0 && def[0] == FileAppend {
		isAppend = true
		flags = flags | os.O_APPEND
	}
	f, err := os.OpenFile(filePath, flags, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	if isAppend {
		stat, _ := f.Stat()
		f.WriteAt(data, stat.Size())
	} else {
		f.Truncate(0)
		f.Write(data)
	}
	return nil
}

// FilesWithSuffixInDir 查找文件夹中符合后缀要求的文件名
// NOTICE: 如果用类似的后缀，如 conf config, 函数不去重
func FilesWithSuffixInDir(dir string, suffixes []string) (res []string, err error) {
	files, err := filepath.Glob(dir + "/*")
	if err != nil {
		return nil, err
	}

	matcher := func(file string, suffixes []string) bool {
		for _, suffix := range suffixes {
			if strings.HasSuffix(file, suffix) {
				return true
			}
		}

		return false
	}

	for _, file := range files {
		if matcher(file, suffixes) {
			res = append(res, file)
		}
	}

	return
}
