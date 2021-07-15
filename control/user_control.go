package control

import (
	"LaodamingMVC/model"
	"LaodamingMVC/network"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)
func GetUser(c *gin.Context){
	UserID := c.Param("ID")
	ID,_:= strconv.ParseInt(UserID,10,64)
	var userInfo model.UserInfo
	err := model.GetUserInfo(ID,&userInfo)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	c.JSON(200,userInfo)
}
func AddUser(c *gin.Context){
	model.AddUserInfo()
}

func SendMsg(c *gin.Context){
	network.SendMessage([]byte("hellow world!!"))

}

func ShowWeChat(c *gin.Context){
	c.HTML(200,"ws.html",nil)
}