package role

import (
	"github.com/gin-gonic/gin"
	"walle-go/pkg/app"
)

type Role struct{}

var API = Role{}

func (r Role) Get(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.ResponseError("Method not implemented")
	return
}

func (r Role) List(c *gin.Context) {
	var appG = app.Gin{C: c}
	roles := []map[string]interface{}{
		{
			"id": "SUPER", "name": "超级管理员",
		},
		{
			"id": "OWNER", "name": "空间所有者",
		},
		{
			"id": "MASTER", "name": "项目管理员",
		},
		{
			"id": "DEVELOPER", "name": "开发者",
		},
		{
			"id": "REPORTER", "name": "访客",
		},
	}
	appG.ResponseList(roles, len(roles), true)
	return
}

func (r Role) Post(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.ResponseError("Method not implemented")
	return
}

func (r Role) Put(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.ResponseError("Method not implemented")
	return
}

func (r Role) Delete(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.ResponseError("Method not implemented")
	return
}
