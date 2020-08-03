package user

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"walle-go/models/user_model"
	"walle-go/pkg/app"
	"walle-go/pkg/security"
)

type User struct{}

var API = User{}

func (u User) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	item := user_model.Item(id)
	appG.ResponseSuccess(item)
}

type RegistrationForm struct {
	Email    string `form:"email" binding:"required"`
	Username string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

func (form *RegistrationForm) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":    form.Email,
		"username": form.Username,
		"password": security.GeneratePasswordHash(form.Password),
	}
}

type UserUpdateForm struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

func (form *UserUpdateForm) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"username": form.Username,
		"password": security.GeneratePasswordHash(form.Password),
	}
}

func (u User) Post(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form RegistrationForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	id := user_model.Add(form.ToDict())
	if id == nil {
		appG.ResponseError("创建出错")
		return
	}
	item := user_model.Item(*id)
	appG.ResponseSuccess(item)
}

func (u User) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	pageNum := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	name := c.Query("kw")
	list, count := user_model.List(pageNum, pageSize, name)
	appG.ResponseList(list, count, true)
}

func (u User) Put(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form UserUpdateForm
	)
	id := com.StrTo(c.Param("id")).MustInt()
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	err = user_model.Update(id, form.ToDict())
	if err != nil {
		appG.ResponseError("更新出错")
		return
	}
	item := user_model.Item(id)
	appG.ResponseSuccess(item)
}

func (u User) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	user_model.Remove(id)
	appG.ResponseSuccess(nil)
}
