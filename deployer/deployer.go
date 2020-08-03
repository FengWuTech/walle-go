package deployer

import (
	"errors"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"strconv"
	"strings"
	"time"
	"walle-go/git"
	"walle-go/logger"
	"walle-go/models"
	"walle-go/models/project_model"
	"walle-go/models/record_model"
	"walle-go/models/server_model"
	"walle-go/models/task_model"
	"walle-go/waller"
)

const (
	StagePrevDeploy = "prev_deploy"
	StageDeploy     = "deploy"
	StagePostDeploy = "post_deploy"

	StagePrevRelease = "prev_release"
	StageRelease     = "release"
	StagePostRelease = "post_release"
	StageEnd         = "end"
	IsRollback       = 1

	CodeBase = "/tmp/walle/codebase/"

	TaskStatusDoing   = 3
	TaskStatusSuccess = 4
	TaskStatusFail    = 5
)

type Server struct {
	Host string
	Port int
	User string
}

type Deployer struct {
	ProjectId          string
	TaskId             string
	Stage              string
	Sequence           int
	Version            string
	LastReleaseVersion string
	ReleaseVersion     string
	DirCodebaseProject string
	DirReleaseVersion  string
	Env                map[string]string
	Room               string
	Repo               *git.Repo
	Localhost          *waller.Waller
	Task               *task_model.Tasks
	Project            *project_model.Projects
	Server             *socketio.Server
}

func globalEnv(d *Deployer) map[string]string {
	now := time.Now()
	return map[string]string{
		"WEBROOT": *d.Project.TargetRoot,
		"VERSION": d.ReleaseVersion,
		"CURRENT_RELEASE": *d.Project.TargetReleases,
		"BRANCH": *d.Task.Branch,
		"TAG": *d.Task.Tag,
		"COMMIT_ID": *d.Task.CommitId,
		"PROJECT_NAME": *d.Project.Name,
		"PROJECT_ID": d.ProjectId,
		"TASK_NAME": *d.Task.Name,
		"TASK_ID": d.TaskId,
		"DEPLOY_USER": *d.Task.UserName,
		"DEPLOY_TIME": now.Format("20060102-150405"),
	}
}

func New(taskId int) *Deployer {
	now := time.Now()
	var (
		task    task_model.Tasks
		project project_model.Projects
	)
	models.GetDB().Where("id = ?", taskId).First(&task)
	models.GetDB().Where("id = ?", *task.ProjectId).First(&project)
	releaseVersion := fmt.Sprintf("%d_%d_%s", *task.ProjectId, taskId, now.Format("20060102_150405"))
	if *task.IsRollback == IsRollback {
		releaseVersion = *task.LinkId
	}
	return &Deployer{
		ProjectId:      fmt.Sprintf("%d", *task.ProjectId),
		TaskId:         fmt.Sprintf("%d", *task.Id),
		Stage:          "",
		Sequence:       0,
		Env:            make(map[string]string),
		Version:        now.Format("20060102150405"),
		ReleaseVersion: releaseVersion,
		Task:           &task,
		Project:        &project,
		Room:           fmt.Sprintf("%d", *task.Id),
		Repo:           &git.Repo{Path: fmt.Sprintf("%s%d", CodeBase, *project.Id),},
		Localhost: &waller.Waller{
			Dir:  CodeBase,
			Env:  nil,
			Host: "127.0.0.1",
			User: "walle",
			Room: fmt.Sprintf("%d", *task.Id),
		},
	}
}

func NewProject(projectId int) *Deployer {
	var (
		project project_model.Projects
	)
	models.GetDB().Where("id = ?", projectId).First(&project)
	return &Deployer{
		Room:    fmt.Sprintf("%d", projectId),
		Project: &project,
		Repo:    &git.Repo{Path: fmt.Sprintf("%s%d", CodeBase, *project.Id),},
	}
}

func (d *Deployer) queryServers() []Server {
	var servers []Server
	var projectServers []server_model.Servers
	serverIds := strings.Split(*d.Project.ServerIds, ",")
	models.GetDB().Where("id in (?)", serverIds).Find(&projectServers)
	for _, server := range projectServers {
		servers = append(servers, Server{
			Host: *server.Host,
			Port: *server.Port,
			User: *server.User,
		})
	}
	return servers
}

func (d *Deployer) WalleDeploy() error {
	d.loadTaskVarsAndGlobalToEnv(*d.Project.TaskVars)
	d.Localhost.LoadEnv(d.Env)
	d.Localhost.Server = d.Server
	// start
	err := d.start()
	if err != nil {
		return err
	}

	// 非回滚，会有deploy相关操作
	if *d.Task.IsRollback != IsRollback {
		// deploy前置操作
		err = d.preDeploy()
		if err != nil {
			return err
		}

		// deploy
		err = d.deploy()
		if err != nil {
			return err
		}

		// deploy后置操作
		err = d.postDeploy()
		if err != nil {
			return err
		}
	}

	servers := d.queryServers()
	allServersSuccess := true

	for _, v := range servers {
		w := &waller.Waller{
			Dir:    "",
			Env:    nil,
			Host:   v.Host,
			User:   v.User,
			Port:   v.Port,
			Server: d.Server,
			Room:   d.Localhost.Room,
		}
		w.LoadEnv(d.Env)

		// release前置操作
		if *d.Task.IsRollback != IsRollback {
			err = d.preRelease(w)
			if err != nil {
				logger.Debugf("preRelease %s err: %v", w.Host, err)
				allServersSuccess = false
				continue
			}
		} else {
			err = d.preReleaseCustom(w)
			if err != nil {
				logger.Debugf("rollback preReleaseCustom %s err: %v", w.Host, err)
				allServersSuccess = false
				continue
			}
		}
		// release
		d.release(w)
		if err != nil {
			logger.Debugf("release %s err: %v", w.Host, err)
			allServersSuccess = false
			continue
		}

		// release后置操作
		d.postRelease(w)
		if err != nil {
			logger.Debugf("postRelease %s err: %v", w.Host, err)
			allServersSuccess = false
			continue
		}
		stage := StageEnd
		sequence := 0
		wenv := d.config()
		taskId, _ := strconv.Atoi(wenv["task_id"])
		userId, _ := strconv.Atoi(wenv["user_id"])
		exitCode := 0
		command := ""
		record := record_model.Records{
			Stage:    &stage,
			Sequence: &sequence,
			UserId:   &userId,
			TaskId:   &taskId,
			Status:   &exitCode,
			Host:     &w.Host,
			User:     &w.User,
			Command:  &command,
		}
		models.GetDB().Create(&record)
		d.Server.BroadcastToRoom("/walle", d.Room, "success", map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": fmt.Sprintf("%s 部署成功!", w.Host), "host": w.Host}})
	}
	d.End(allServersSuccess)
	return nil
}

func (d *Deployer) loadTaskVarsAndGlobalToEnv(taskVarsStr string) error {
	envs := globalEnv(d)
	for k, v := range envs {
		d.Env[k] = v
	}
	taskVars := strings.Split(taskVarsStr, "\n")
	for _, var_ := range taskVars {
		var_ = strings.TrimSpace(var_)
		if len(var_) > 0 && string(var_[0]) != "#" {
			varList := strings.Split(var_, "=")
			if len(varList) == 2 {
				d.Env[strings.TrimSpace(varList[0])] = strings.TrimSpace(varList[1])
			}
		}
	}
	return nil
}

func (d *Deployer) start() error {
	// remove all old records of this task
	models.GetDB().Where("task_id = ?", d.Task.Id).Delete(record_model.Records{})
	// change task status to doing
	models.GetDB().Model(&task_model.Tasks{}).Where("id = ?", d.Task.Id).Update("status", TaskStatusDoing)
	d.DirCodebaseProject = CodeBase + d.ProjectId
	d.DirReleaseVersion = CodeBase + d.ReleaseVersion
	return nil
}

func (d *Deployer) InitRepo() error {
	logger.Debugf("init repo: %s", *d.Project.RepoUrl)
	d.Repo.Init(*d.Project.RepoUrl, "master")
	return nil
}

func (d *Deployer) config() map[string]string {
	return map[string]string{
		"task_id": d.TaskId,
		"user_id": fmt.Sprintf("%d", *d.Task.UserId),
		"stage": d.Stage,
		"sequence": fmt.Sprintf("%d", d.Sequence),
	}
}

func (d *Deployer) toMailMessage(title string) string {
	version := ""
	if d.Task != nil && *d.Task.Tag != "" {
		version = *d.Task.Tag
	} else {
		version = fmt.Sprintf("%s/%s", *d.Task.Branch, *d.Task.CommitId)
	}
	return fmt.Sprintf(` %s %s 
	<br><br> <strong>项目</strong>：%s
	<br><br> <strong>任务</strong>：%s ( http://walle.gmtech.top/task/deploy/%d )
	<br><br> <strong>分支</strong>：%s
	<br><br><br><img src='http://walle-web.io/dingding.jpg'>`, *d.Task.UserName, title, *d.Project.Name,
	*d.Task.Name, *d.Task.Id, version)
}

func (d *Deployer) End(success bool) error {
	// send mail
	logger.Debugf("last version: %s", d.LastReleaseVersion)
	logger.Debugf("current version: %s", d.ReleaseVersion)
	status :=  TaskStatusFail
	if success {
		status = TaskStatusSuccess
	}
	models.GetDB().Model(&task_model.Tasks{}).Where("id = ?", d.Task.Id).Updates(task_model.Tasks{
		Status: &status,
		LinkId: &d.ReleaseVersion,
		ExLinkId: &d.LastReleaseVersion,
	})
	title := "上线部署成功"
	if !success {
		title = "上线部署失败"
	}
	body := d.toMailMessage(title)
	SendMail(title, body, strings.Split(*d.Project.NoticeHook, ","))
	logger.Infof("Send mail %s to %v", title, *d.Project.NoticeHook)


	if !success {
		d.Server.BroadcastToRoom("/walle", d.Room, "fail", map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": title}})
		return errors.New("部署上线失败")
	} else {
		d.Server.BroadcastToRoom("/walle", d.Room, "success", map[string]interface{}{"event": "finish", "data": map[string]interface{}{"message": title}})
	}
	return nil
}
