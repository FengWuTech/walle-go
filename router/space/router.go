package space

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"walle-go/models/space_model"
	"walle-go/pkg/app"
)

type Space struct{}

var API = Space{}

func (s Space) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	item := space_model.Item(id)
	appG.ResponseSuccess(item)
	return
}

func (s Space) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	list, count := space_model.List(pageNum, pageSize, name)
	appG.ResponseList(list, count, true)
	return
}

type PostForm struct {
	Name   string `form:"name" binding:"required"`
	UserId int    `form:"user_id"`
	Status int    `form:"status"`
}

func (form *PostForm) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":    form.Name,
		"user_id": form.UserId,
		"status":  1,
	}
}

func (s Space) Post(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	id := space_model.Add(form.ToDict())
	if id == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := space_model.Item(*id)
	appG.ResponseSuccess(item)
	return
}

func (s Space) Put(c *gin.Context) {
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
	err = space_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := space_model.Item(id)
	appG.ResponseSuccess(item)
	return
}

func (s Space) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	space_model.Remove(id)
	appG.ResponseSuccess(nil)
	return
}
