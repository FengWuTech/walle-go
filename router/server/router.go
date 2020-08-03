package server

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"walle-go/models/server_model"
	"walle-go/pkg/app"
)

type Server struct {}

var API = Server{}

func (s Server) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	item := server_model.Item(id)
	appG.ResponseSuccess(item)
}

type PostForm struct {
	Name string `form:"name" binding:"required"`
	Host string `form:"host"`
	User string `form:"user"`
	Port int    `form:"port"`
}

func (form *PostForm) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name": form.Name,
		"host": form.Host,
		"user": form.User,
		"port": form.Port,
		"status": 1,
	}
}

func (s Server) Post(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	id := server_model.Add(form.ToDict())
	if id == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := server_model.Item(*id)
	appG.ResponseSuccess(item)
}

func (s Server) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	list, count := server_model.List(pageNum, pageSize, name)
	appG.ResponseList(list, count, true)
}

func (s Server) Put(c *gin.Context) {
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
	err = server_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := server_model.Item(id)
	appG.ResponseSuccess(item)
}

func (s Server) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	server_model.Remove(id)
	appG.ResponseSuccess(nil)
}

