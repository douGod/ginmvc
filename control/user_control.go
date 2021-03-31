package control

import (
	"LaodamingMVC/model"
	"fmt"
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
	var arr = []byte{10,5,4,7,6,18,1}
	var length = len(arr)
	//冒泡排序
	for  i:= 0;i < length - 1;i++{
		for j:= 1; j < length - i - 1;j++{
			if arr[j] > arr[j - 1]{
				arr[j],arr[j - 1] = arr[j - 1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func ShowWeChat(c *gin.Context){
	c.HTML(200,"ws.html",nil)
}