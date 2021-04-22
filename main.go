package main

import (
	"LaodamingMVC/network"
	"LaodamingMVC/router"
	"github.com/gin-gonic/gin"
)

func main(){
    var r *gin.Engine
	//发布模式
	gin.SetMode(gin.ReleaseMode)
	//注册路由
	router.SetupRouter(r)
	//挂起tcp服务端
	go network.SetUpTcpServer()
	//挂起tcp客户端
	go network.SetUpTcpClient()
	//运行框架
	r.Run(":9501")
}

