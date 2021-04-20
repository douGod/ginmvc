package control

import (
	"LaodamingMVC/model"
	"LaodamingMVC/network"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"strconv"
	"sync"
)
var Once sync.Once
var store = base64Captcha.DefaultMemStore
var driver *base64Captcha.DriverString
func GetUser(c *gin.Context){
	UserID := c.Param("ID")
	ID,_:= strconv.ParseInt(UserID,10,64)
	user_info := model.GetUserInfo(ID)
	c.HTML(200,"user_info.html",user_info)
}
func AddUser(c *gin.Context){
	model.AddUserInfo()
}
//测试输出验证码
func Test(c *gin.Context){
	item := flushcode()
	item.WriteTo(c.Writer)
}
func flushcode()base64Captcha.Item{
	drivers := GetDriver().ConvertFonts()
	cs := base64Captcha.NewCaptcha(drivers, store)
	_, content, answer := cs.Driver.GenerateIdQuestionAnswer()
	id := "captcha:yufei"
	item, _ := cs.Driver.DrawCaptcha(content)
	cs.Store.Set(id, answer)
	fmt.Println(answer)
	return item
}
func VerifyCode(c *gin.Context){
	drivers := GetDriver().ConvertFonts()
	cs := base64Captcha.NewCaptcha(drivers, store)
	Code := c.Param("Code")
	id := "captcha:yufei"
	if cs.Store.Verify(id,Code,false){
		flushcode()
		c.String(200,"success")
	}else{
		flushcode()
		c.String(200,"false")
	}



}
//获取验证码驱动
func GetDriver() *base64Captcha.DriverString{
    Once.Do(func(){
		driver = new(base64Captcha.DriverString)
		driver.Height = 50
		driver.Width = 200
		driver.NoiseCount = 5
		driver.ShowLineOptions= base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine
		driver.Length = 6
		driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	})
	return driver
}
func SendMsg(c *gin.Context){
	network.SendMessage([]byte("hellow world!!"))

}

func ShowWeChat(c *gin.Context){
	c.HTML(200,"ws.html",nil)
}