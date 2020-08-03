package general

import (
	"github.com/gin-gonic/gin"
	"walle-go/models/menu_model"
	"walle-go/models/user_model"
	"walle-go/pkg/app"
)

func Menu(c *gin.Context) {
	var appG = app.Gin{C: c}
	uid := app.GetUid(c)
	spaceInfo := map[string]interface{}{
		"id":   1,
		"name": "plm",
		"role": "OWNER",
	}
	spaceList := []interface{}{spaceInfo,}
	space := map[string]interface{}{
		"current":   spaceInfo,
		"available": spaceList,
	}
	ret := map[string]interface{}{
		"user":  user_model.Item(uid),
		"menu":  menu_model.ToList(),
		"space": space,
	}
	appG.ResponseSuccess(ret)
}
