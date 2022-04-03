package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
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
		visitorAPI.GET("path",server.FetchPath)
		connectionAPI := visitorAPI.Group("/connections")
		{
			connectionAPI.GET("all",server.GetConnections)
			connectionAPI.GET("map",server.GetConnectionsByMapId)
		}

		nodeAPI := visitorAPI.Group("/nodes")
		{
			nodeAPI.GET("all",server.GetNodes)
			nodeAPI.GET("map",server.GetNodesByMapId)
			nodeAPI.GET("building",server.GetNodesByBuildingName)
		}
	}
	//授权用户, 需要登陆
	adminAPI := v1.Group("/admin")
	{
		addAPI := adminAPI.Group("/add")
		{ 
			addAPI.POST("node",server.AddNode)
			addAPI.POST("connection",server.AddConnection)
			addAPI.POST("staircase",server.AddStaircase)
			addAPI.POST("map",server.AddMap)
			//addAPI.POST("emergent",server.AddEmergent)
		}

		deleteAPI := adminAPI.Group("/delete")
		{
			//deleteAPI.DELETE("node",server.DeleteNode)
			deleteAPI.POST("connection",server.DeleteConnection)
			deleteAPI.POST("both",server.Delete)
			deleteAPI.POST("map",server.DeleteMap)
		}

		indexAPI := adminAPI.Group("/index")
		{
			indexAPI.GET("nodeId",server.GetNodeId)
		}

		mapAPI := adminAPI.Group("/map")
		{
			//fixme:1 给我nodeID 给你 mapID
			mapAPI.GET("all",server.FetchMaps)
			mapAPI.GET("name",server.FetchMapNames)
			mapAPI.GET("filter/id",server.FetchMapByIdFilter)
			mapAPI.GET("filter/name",server.FetchMapByNameFilter)
			mapAPI.GET("nodeId",server.FetchMapIdByNodeId)
			mapAPI.GET("building",server.FecthBuildingNameByNodeId)
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
	route.Use(Cors())
	//https config

	route.Use(TlsHandler())

	return router(route)
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Max-Age", "36000")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, content-type, Accept,Authorization,authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Header("Content-Type", "application/json;charset=utf-8")
			//c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}
		c.Next()
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}