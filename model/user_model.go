package model

import (
	"LaodamingMVC/database"
	"fmt"
)
//用户表结构
type UserInfo struct {
	ID int64 `gorm:"column:ID"`
	Name string `gorm:"column:Name"`
}

func GetUserInfo(ID int64) *UserInfo{
	Db:= database.GetDb()
	user_info := &UserInfo{}
	Db.Find(user_info,"ID = ?",ID)
	return user_info
}

func AddUserInfo(){
	Db:= database.GetDb()
	user_info := UserInfo{Name:"劳达明"}
	err :=Db.Create(&user_info).Error
	if err != nil{
		fmt.Println(Db.GetErrors())
	}
}