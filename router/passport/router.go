package passport

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"walle-go/models/user_model"
	"walle-go/pkg/app"
	"walle-go/pkg/token"
)

type LoginForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password"  binding:"required"`
}

func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form LoginForm
	)
	err := c.ShouldBind(&form)
	if err != nil {
		appG.ResponseError(err.Error())
		return
	}
	user := user_model.GetByEmail(form.Email)
	if user != nil && user.VerifyPassword(form.Password) {
		url := location.Get(c)
		if url != nil && url.Host != "" {
			tokenStr := token.GenerateToken(*user.Id)
			c.SetCookie(token.COOKIE_NAME, tokenStr, 8640000, "/", url.Host, false, false)
			appG.ResponseSuccess(map[string]interface{}{"user": user.ToMap(), "token": tokenStr})
			return
		}
	}
	appG.ResponseError("账号密码错误")
	return
}

func Logout(c *gin.Context) {
	var appG = app.Gin{C: c}
	url := location.Get(c)
	if url != nil && url.Host != "" {
		c.SetCookie(token.COOKIE_NAME, "logout", 8640000, "/", url.Host, false, false)
	}
	appG.ResponseSuccess(nil)
}
