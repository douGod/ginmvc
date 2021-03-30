package config

import "sync"

var mapConfig map[string] string
var once sync.Once
func init(){
	once.Do(func(){
		mapConfig = make(map[string] string)
		//数据库配置信息
		mapConfig["DB_HOST"] = "127.0.0.1:3306"
		mapConfig["DB_USER"] = "root"
		mapConfig["DB_PWD"] = "123"
		mapConfig["DB_NAME"] = "lytest"
		mapConfig["DB_DRIVER"] = "mysql"
	})

}
func C(key string) string{
	if val,ok := mapConfig[key];!ok{
		return ""
	}else{
		return val
	}
}