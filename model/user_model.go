package model

import (
	"LaodamingMVC/database"
	"fmt"
	"sync"
)
//用户表结构
type UserInfo struct {
	ID int64 `gorm:"column:ID"`
	Name string `gorm:"column:Name"`
}

func GetUserInfo(ID int64) *UserInfo{
	Db:= database.GetMysqlDb()
	user_info := &UserInfo{}
	Db.Find(user_info,"ID = ?",ID)
	return user_info
}

func AddUserInfo(){
	Db:= database.GetMysqlDb()
	user_info := UserInfo{Name:"劳达明"}
	wait := new(sync.WaitGroup)
	wait.Add(100000)
	for i:=0 ; i< 100000;i++{
		go func(w *sync.WaitGroup){
			err :=Db.Create(&user_info).Error
			if err != nil{
				fmt.Println(Db.GetErrors())
			}
			w.Done()
		}(wait)
	}
	fmt.Println("success to add 100000 row data")
}