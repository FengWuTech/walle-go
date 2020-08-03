package router

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"walle-go/pkg/token"
	"walle-go/router/environment"
	"walle-go/router/general"
	"walle-go/router/passport"
	"walle-go/router/project"
	"walle-go/router/role"
	"walle-go/router/server"
	"walle-go/router/space"
	"walle-go/router/task"
	"walle-go/router/user"
)

type APIInterface interface {
	Get(c *gin.Context)
	Put(c *gin.Context)
	Post(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

func Resource(r *gin.Engine, path string, apiInterface APIInterface) {
	api := r.Group(path)
	{
		api.GET("/:id", apiInterface.Get)
		api.GET("/", apiInterface.List)
		api.POST("/", apiInterface.Post)
		api.POST("/:id/:action", apiInterface.Post)
		api.PUT("/:id", apiInterface.Put)
		api.PUT("/:id/:action", apiInterface.Put)
		api.DELETE("/:id", apiInterface.Delete)
	}
}

func TokenParseMiddleware(c *gin.Context) {
	session, err := c.Cookie(token.COOKIE_NAME)
	if err == nil {
		uid, _ := token.ParseTokenUid(session)
		if uid > 0 {
			c.Set("uid", uid)
		}
	}
	c.Next()
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(location.Default())
	r.Use(TokenParseMiddleware)

	apiGeneral := r.Group("/api/general")
	{
		apiGeneral.GET("/menu", general.Menu) // 获取菜单信息
	}
	apiPassport := r.Group("/api/passport")
	{
		apiPassport.POST("/login", passport.Login)   // 登陆
		apiPassport.POST("/logout", passport.Logout) // 登出
	}
	Resource(r, "/api/environment", environment.API)
	Resource(r, "/api/project", project.API)
	Resource(r, "/api/server", server.API)
	Resource(r, "/api/user", user.API)
	Resource(r, "/api/space", space.API)
	Resource(r, "/api/role", role.API)
	Resource(r, "/api/task", task.API)

	return r
}
