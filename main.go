package main

import (
	"LaodamingMVC/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)
func TestJwt(){
	file ,_ := os.Open("C:\\Users\\Administrator\\Desktop\\a.txt")
	defer file.Close()
	fileData := make([]byte,512)
	n,_ := file.Read(fileData)
	fmt.Println(string(fileData[:n]))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(fileData[:n])
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tokenString)
	token ,err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return fileData[:n],nil
	})
	if token.Valid{
		fmt.Println(token.Claims.(jwt.MapClaims))
	}
}
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

