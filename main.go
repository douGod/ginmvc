package main

import (
	"LaodamingMVC/router"
	"github.com/gin-gonic/gin"
)

func main(){
	//发布模式
	gin.SetMode(gin.ReleaseMode)
	//注册路由
	r := router.SetupRouter()
	//挂起tcp服务端
	//go network.SetUpTcpServer()
	//挂起tcp客户端
	//go network.SetUpTcpClient()
	//运行框架
	r.Run(":9501")
}

