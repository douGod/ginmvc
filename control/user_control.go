package control

import (
	"LaodamingMVC/model"
	"LaodamingMVC/network"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context){
	UserID := c.Param("ID")
	ID,_:= strconv.ParseInt(UserID,10,64)
	user_info := model.GetUserInfo(ID)
	c.HTML(200,"user_info.html",user_info)
}
func AddUser(c *gin.Context){
	model.AddUserInfo()
}
func Test(c *gin.Context){
	go network.SendMessage([]byte(c.Param("Message")))
}

func ShowWeChat(c *gin.Context){
	c.HTML(200,"ws.html",nil)
}