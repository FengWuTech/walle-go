package project

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"walle-go/models"
	"walle-go/models/project_model"
	"walle-go/pkg/app"
)

type Project struct {}

var API = Project{}

func (p Project) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	envId := com.StrTo(c.Param("id")).MustInt()
	item := project_model.Item(envId)
	appG.ResponseSuccess(item)
}

type PostForm struct {
	Name           string `form:"name" binding:"required"`
	SpaceId        int    `form:"space_id"`
	EnvironmentId  int    `form:"environment_id"`
	Status         int    `form:"status"`
	Excludes       string `form:"excludes"`
	IsInclude      int    `form:"is_include"`
	Master         string `form:"master"`
	ServerIds      string `form:"server_ids"`
	KeepVersionNum int    `form:"keep_version_num"`
	TargetRoot     string `form:"target_root"`
	TargetReleases string `form:"target_releases"`
	TaskVars       string `form:"task_vars"`
	PrevDeploy     string `form:"prev_deploy"`
	PostDeploy     string `form:"post_deploy"`
	PrevRelease    string `form:"prev_release"`
	PostRelease    string `form:"post_release"`
	RepoUrl        string `form:"repo_url"`
	RepoUsername   string `form:"repo_username"`
	RepoPassword   string `form:"repo_password"`
	RepoMode       string `form:"repo_mode"`
	NoticeType     string `form:"notice_type"`
	NoticeHook     string `form:"notice_hook"`
	TaskAudit      int    `form:"task_audit"`
}

func (form *PostForm) ToDict() map[string]interface{} {
	status := form.Status
	if status == 0 {
		status = models.StatusDefault
	}
	spaceId := 1
	return map[string]interface{}{
		"name": form.Name,
		"user_id": 1,
		"status": status,
		"master": form.Master,
		"environment_id": form.EnvironmentId,
		"space_id": spaceId,
		"excludes": form.Excludes,
		"is_include": form.IsInclude,
		"server_ids": form.ServerIds,
		"keep_version_num": form.KeepVersionNum,
		"target_root": form.TargetRoot,
		"target_releases": form.TargetReleases,
		"task_vars": form.TaskVars,
		"prev_deploy": form.PrevDeploy,
		"post_deploy": form.PostDeploy,
		"prev_release": form.PrevRelease,
		"post_release": form.PostRelease,
		"repo_url": form.RepoUrl,
		"repo_username": form.RepoUsername,
		"repo_password": form.RepoPassword,
		"repo_mode": form.RepoMode,
		"notice_type": form.NoticeType,
		"notice_hook": form.NoticeHook,
		"task_audit": form.TaskAudit,
	}
}

func create(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	envId := project_model.Add(form.ToDict())
	if envId == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := project_model.Item(*envId)
	appG.ResponseSuccess(item)
}

func members(c *gin.Context) {
	// TODO: 替换真实功能
	var (
		appG = app.Gin{C: c}
	)
	appG.ResponseSuccess(nil)
	return
}

func _copy(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt()
	item := project_model.Item(id)
	delete(item, "id")
	item["name"] = item["name"].(string) + "-copy"
	newId := project_model.Add(item)
	if newId == nil {
		appG.ResponseError("创建出错")
		return
	}
	newItem := project_model.Item(*newId)
	appG.ResponseSuccess(newItem)
}

func detection(c *gin.Context) {
	// TODO: 替换真实功能
	var (
		appG = app.Gin{C: c}
	)
	appG.ResponseWithMsgAndCode(http.StatusOK, 0, "配置检测通过，恭喜：）开始你的上线之旅吧", nil)
	return
}

func (p Project) Post(c *gin.Context) {
	if c.Param("action") == "" {
		create(c)
		return
	}
	switch c.Param("action")  {
	case "members":
		members(c)
	case "copy":
		_copy(c)
	case "detection":
		detection(c)
	}
}

func (p Project) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	environmentId := com.StrTo(c.DefaultQuery("environment_id", "0")).MustInt()
	spaceId := 1
	list, count := project_model.List(pageNum, pageSize, name, environmentId, spaceId)
	appG.ResponseList(list, count, true)
}

func (p Project) Put(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	id := com.StrTo(c.Param("id")).MustInt()
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	err = project_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := project_model.Item(id)
	appG.ResponseSuccess(item)
}

func (p Project) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	project_model.Remove(id)
	appG.ResponseSuccess(nil)
}

