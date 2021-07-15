package database

import (
	"LaodamingMVC/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)
//操作数据库指针
var MysqlDb *gorm.DB
var once sync.Once
func init(){
	var err error
	once.Do(func(){
		var Dsn = config.C("DB_USER")+":"+ config.C("DB_PWD")+"@("+config.C("DB_HOST")+")/"+config.C("DB_NAME")+"?charset=utf8&parseTime=true"
		if MysqlDb,err = gorm.Open(config.C("DB_DRIVER"),Dsn);err != nil{
			panic(err.Error())
		}
		//myDb.LogMode(true)//logdebug打印sql
		fmt.Println("success to connect db")
		//数据表设置为单数
		MysqlDb.SingularTable(true)
		//空闲时连接数
		MysqlDb.DB().SetMaxIdleConns(10)
		//最大连接数
		MysqlDb.DB().SetMaxOpenConns(100)
	})
	if err = MysqlDb.DB().Ping();err != nil{
		panic(err.Error())
	}
}
