package database

import (
	"LaodamingMVC/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
//操作数据库指针
var myDb *gorm.DB
func init(){
	connectDb()
}
func connectDb() bool{
	var err error
	var Dsn = config.C("DB_USER")+":"+ config.C("DB_PWD")+"@("+config.C("DB_HOST")+")/"+config.C("DB_NAME")+"?charset=utf8&parseTime=true"
	if myDb,err = gorm.Open(config.C("DB_DRIVER"),Dsn);err != nil{
		fmt.Println(err)
		return false
	}
	//myDb.LogMode(true)//logdebug打印sql
	fmt.Println("success to connect db")
	//数据表设置为单数
	myDb.SingularTable(true)
	//空闲时连接数
	myDb.DB().SetMaxIdleConns(10)
	//最大连接数
	myDb.DB().SetMaxOpenConns(100)
	return true
}
func GetDb()*gorm.DB{
	return myDb
}