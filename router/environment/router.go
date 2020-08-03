package environment

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"walle-go/models/environment_model"
	"walle-go/pkg/app"
)

type Environment struct {}

var API = Environment{}

func (e Environment) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	item := environment_model.Item(id)
	appG.ResponseSuccess(item)
}

type PostForm struct {
	EnvName string `form:"env_name" binding:"required"`
	Status  int    `form:"status"`
	SpaceId int    `form:"space_id"`
	EnvId   int    `form:"env_id"`
}

func (form *PostForm) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name": form.EnvName,
		"space_id": 1,
		"status": 1,
	}
}

func (e Environment) Post(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	id := environment_model.Add(form.ToDict())
	if id == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := environment_model.Item(*id)
	appG.ResponseSuccess(item)
}

func (e Environment) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	spaceId := 1
	list, count := environment_model.List(pageNum, pageSize, name, spaceId)
	appG.ResponseList(list, count, true)
}

func (e Environment) Put(c *gin.Context) {
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
	err = environment_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := environment_model.Item(id)
	appG.ResponseSuccess(item)
}

func (e Environment) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	environment_model.Remove(id)
	appG.ResponseSuccess(nil)
}
