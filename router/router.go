package router

import (
	"LaodamingMVC/control"
	"LaodamingMVC/network"
	"github.com/gin-gonic/gin"
)
func noRouter(c *gin.Context){
	c.String(404,"无效的访问路由")
}
func SetupRouter() gin.Engine{
	r := gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("./view/*")
	//设置访问路由
	r.GET("/GetUser/:ID",control.GetUser)
	r.GET("/AddUser",control.AddUser)
	r.GET("/SendMsg",control.SendMsg)
	//聊天页面路由
	r.GET("/",control.ShowWeChat)
	//websocket连接升级请求
	r.GET("/ws",network.WeChat)
	r.NoRoute(noRouter)
	return *r
}
