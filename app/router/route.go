package router

import (
	"github.com/gin-gonic/gin"
	"pro/app/http/v1/server"
	"pro/app/middle"
	"pro/app/socket"
	"pro/config"
)

func router(route *gin.Engine) *gin.Engine {

	//socket服务器
	route.GET("/ws", socket.Run)
	//route.GET("/ws/ping", socket.Ping)

	v1 := route.Group("/v1")
	//遊客操作，无需登录
	visitorAPI := v1.Group("/api")
	{
		visitorAPI.GET("index", server.Index)
		visitorAPI.GET("test", server.Test)
		visitorAPI.GET("search",server.Search)
	}
	//授权用户, 需要登陆
	adminAPI := v1.Group("/admin")
	{
		addAPI := adminAPI.Group("/add")
		{ 
			addAPI.POST("node",server.AddNode)
			addAPI.POST("connection",server.AddConnection)
		}

		deleteAPI := adminAPI.Group("/delete")
		{
			deleteAPI.DELETE("node",server.DeleteNode)
			deleteAPI.DELETE("connection",server.DeleteConnection)
		}

	}

	return route
}

func RouteInit() *gin.Engine {
	if config.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	route := gin.New()
	if config.Mode == "dev" {
		route.Use(gin.Logger())
	}
	route.Use(gin.Recovery()) // 捕捉异常
	route.Use(middle.Access)
	return router(route)
}
