package task

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"walle-go/models/task_model"
	"walle-go/pkg/app"
)

type Task struct {}

var API = Task{}

func (t Task) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	item := task_model.Item(id)
	appG.ResponseSuccess(item)
}

type PostForm struct {
	Name      string `form:"name" binding:"required"`
	ProjectId int    `form:"project_id"`
	Servers   string `form:"servers"`
	CommitId  string `form:"commit_id"`
	Tag string `form:"tag"`
	Branch string `form:"branch"`
}

func (form *PostForm) ToDict() map[string]interface{} {
	uid := 1
	username := "test"
	return map[string]interface{}{
		"name": form.Name,
		"user_id": uid,
		"user_name": username,
		"project_id": form.ProjectId,
		"action": 0,
		"status": 0,
		"link_id": "",
		"ex_link_id": "",
		"servers": form.Servers,
		"commit_id": form.CommitId,
		"tag": form.Tag,
		"branch": form.Branch,
		"is_rollback": 0,
	}
}

func (t Task) Post(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	id := task_model.Add(form.ToDict())
	if id == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := task_model.Item(*id)
	appG.ResponseSuccess(item)
}

func (t Task) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	list, count := task_model.List(pageNum, pageSize, name)
	appG.ResponseList(list, count, true)
}

func (t Task) Put(c *gin.Context) {
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
	err = task_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := task_model.Item(id)
	appG.ResponseSuccess(item)
}

func (t Task) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	task_model.Remove(id)
	appG.ResponseSuccess(nil)
}


