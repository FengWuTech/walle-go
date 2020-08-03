package wsio //walle socket io

import (
	socketio "github.com/googollee/go-socket.io"
	"strconv"
	"time"
	"walle-go/deployer"
	"walle-go/logger"
	"walle-go/models"
	"walle-go/models/record_model"
	"walle-go/models/task_model"
	"walle-go/pkg/timeutil"
)

type WebSocketIO struct {
	Server    *socketio.Server
	Namespace string
}

func (w *WebSocketIO) Setup(server *socketio.Server) {

	w.Server = server

	//server.OnConnect("/", w.onOpen)
	server.OnEvent(w.Namespace, "open", w.onOpen)
	server.OnEvent(w.Namespace, "deploy", w.onDeploy)
	server.OnEvent(w.Namespace, "branches", w.onBranches)
	server.OnEvent(w.Namespace, "tags", w.onTags)
	server.OnEvent(w.Namespace, "commits", w.onCommits)
	server.OnEvent(w.Namespace, "ping", w.onPing)
	server.OnEvent(w.Namespace, "logs", w.onLogs)

}

func parseRoomId(s socketio.Conn) (int, string) {
	room := s.Context().(map[string]interface{})["room_id"]
	roomIdStr := room.(string)
	roomId, _ := strconv.Atoi(roomIdStr)
	return roomId, roomIdStr
}

// 42/walle,["open",{"project_id":3}]
// 传过来的是一个list, 元素1是指令，对应func；元素2是message，一个map.
func (w *WebSocketIO) onOpen(s socketio.Conn, mInfo map[string]interface{}) {

	logger.Debugf("on open: %v", mInfo)
	roomId := 0
	isTask := false

	if _, ok := mInfo["task"]; ok {
		roomId, _ = strconv.Atoi(mInfo["task"].(string))
		isTask = true
	}

	if _, ok := mInfo["project_id"]; ok {
		roomId = int(mInfo["project_id"].(float64))
	}
	logger.Debugf("open room: %d", roomId)
	roomIdStr := strconv.Itoa(roomId)
	s.SetContext(map[string]interface{}{"room_id": roomIdStr})

	//加入房屋
	logger.Debugf("join room: %d", roomId)
	s.Join(roomIdStr)

	if isTask {
		logger.Debugf("get task: %d", roomId)
		taskInfo := task_model.Item(roomId)

		logger.Debugf("send construct to room: %d", roomId)
		//发送消息
		w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "construct", map[string]interface{}{"event": "connect", "data": taskInfo})
	}
}

func (w *WebSocketIO) onDeploy(s socketio.Conn, mInfo map[string]interface{}) {
	logger.Debugf("on deploy: %v", mInfo)
	roomId, _ := parseRoomId(s)

	d := deployer.New(roomId)
	// 解决循环引用，先把socket作为参数传进去
	d.Server = w.Server
	d.InitRepo()

	go func() {
		err := d.WalleDeploy()
		if err != nil {
			d.End(false)
		}
	}()

	//发送消息
	//w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "console", map[string]interface{}{"event": "forbidden", "data": nil})
}

func (w *WebSocketIO) onBranches(s socketio.Conn) {
	logger.Debugf("on branches")
	roomId, roomIdStr := parseRoomId(s)

	d := deployer.NewProject(roomId)
	d.InitRepo()

	branches := d.Repo.Branches()
	//发送消息
	w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "branches", map[string]interface{}{"event": "branches", "data": branches})
}

func (w *WebSocketIO) onTags(s socketio.Conn, mInfo map[string]interface{}) {
	logger.Debugf("on tags: %v", mInfo)
	_, roomIdStr := parseRoomId(s)
	tags := []string{"v0.1", "v0.2", "v0.3"}
	//发送消息
	w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "tags", map[string]interface{}{"event": "tags", "data": tags})
}

func (w *WebSocketIO) onCommits(s socketio.Conn, mInfo map[string]interface{}) {
	roomId, roomIdStr := parseRoomId(s)
	branch := mInfo["branch"].(string)
	logger.Debugf("list branch %s", branch)

	d := deployer.NewProject(roomId)
	d.InitRepo()

	commits := d.Repo.Commits(branch)
	//发送消息
	w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "commits", map[string]interface{}{"event": "commits", "data": commits})
}

func (w *WebSocketIO) onPing(s socketio.Conn, msg string) {
	_, roomIdStr := parseRoomId(s)
	ctime := timeutil.Format("%Y-%m-%d %H:%M:%S", time.Now().Local())
	//发送消息
	w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "pong", map[string]interface{}{"event": "ping:pong", "data": map[string]interface{}{"time": ctime}})
	logger.Info(msg)
}

func (w *WebSocketIO) onLogs(s socketio.Conn, mInfo map[string]interface{}) {
	logger.Debugf("on logs: %v", mInfo)
	_, roomIdStr := parseRoomId(s)
	var (
		task    task_model.Tasks
		records []record_model.Records
	)
	models.GetDB().Where("id = ?", roomIdStr).First(&task)
	if task.Id == nil {
		return
	}
	if *task.Status != deployer.TaskStatusDoing && *task.Status != deployer.TaskStatusSuccess && *task.Status != deployer.TaskStatusFail {
		w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "console", map[string]interface{}{"event": "console", "data": ""})
		return
	}
	models.GetDB().Where("task_id = ?", roomIdStr).Find(&records)
	for _, record := range records {
		log := map[string]interface{}{
			"host":       record.Host,
			"user_model": record.User,
			"cmd":        record.Command,
			"status":     record.Status,
			"stage":      record.Stage,
			"sequence":   record.Sequence,
			"success":    record.Success,
			"error":      record.Error,
		}
		if *record.Stage == deployer.StageEnd {
			msg := "部署成功"
			cmd := "success"
			if *task.Status == deployer.TaskStatusFail {
				msg = "部署失败"
				cmd = "fail"
			}
			w.Server.BroadcastToRoom(w.Namespace, roomIdStr, cmd, map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": msg, "host": record.Host}})
		} else {
			w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "console", map[string]interface{}{"event": "console", "data": log})
		}
	}
	if *task.Status == deployer.TaskStatusSuccess {
		msg := "success"
		w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "success", map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": msg}})
	}
	if *task.Status == deployer.TaskStatusFail {
		msg := "fail"
		w.Server.BroadcastToRoom(w.Namespace, roomIdStr, "fail", map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": msg}})
	}
}
