package control

import (
	"LaodamingMVC/database"
	"LaodamingMVC/model"
	"LaodamingMVC/network"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
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
//测试redis连接池
func TestRedis(c *gin.Context){
	redisDb := database.GetRedisDb()
	err := redisDb.SetNX("ldm","劳达明",0).Err()
	if err != nil{
		log.Fatal(err)
	}
	val,err := redisDb.Get("ldm").Result()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(val)
}
func PostData(c *gin.Context){
	type FormData struct {
		Name string `form:"Name"`
		ID int `form:"ID"`
		VerCode string `form:"VerCode"`
	}
	data := &FormData{}
	c.ShouldBind(data)
	if !verifyCode(data.VerCode){
		c.String(200,"验证码错误")
		return
	}
	c.JSON(200,data)
}
//测试输出验证码
func DrawVerCode(c *gin.Context){
	item := flushcode()
	item.WriteTo(c.Writer)
}
//刷新验证码
func flushcode()base64Captcha.Item{
	drivers := getDriver().ConvertFonts()
	cs := base64Captcha.NewCaptcha(drivers, store)
	_, content, answer := cs.Driver.GenerateIdQuestionAnswer()
	id := "laodaming"//自定义，一般是用户唯一标识
	item, _ := cs.Driver.DrawCaptcha(content)
	cs.Store.Set(id, answer)
	fmt.Println(answer)
	return item
}
func verifyCode(VerCode string) bool{
	drivers := getDriver().ConvertFonts()
	cs := base64Captcha.NewCaptcha(drivers, store)
	id := "laodaming"//自定义，一般是用户唯一标识
	if cs.Store.Verify(id,VerCode,false){
		flushcode()
		return true
	}else{
		flushcode()
		return false
	}



}
//获取验证码驱动
func getDriver() *base64Captcha.DriverString{
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