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
	//运行框架
	r.Run(":9501")
}

