package model

import (
	"LaodamingMVC/database"
	"fmt"
	"sync"
)
//用户表结构
type UserInfo struct {
	ID int64 `gorm:"primaryKey"`
	Name string `gorm:"default:''"`
}

func GetUserInfo(ID int64,userInfo *UserInfo)error{

	return database.MysqlDb.Debug().Find(userInfo,"id = ?",ID).Error
}

func AddUserInfo(){
	user_info := UserInfo{Name:"劳达明"}
	wait := new(sync.WaitGroup)
	wait.Add(100000)
	for i:=0 ; i< 100000;i++{
		go func(w *sync.WaitGroup){
			err :=database.MysqlDb.Create(&user_info).Error
			if err != nil{
				fmt.Println(database.MysqlDb.GetErrors())
			}
			w.Done()
		}(wait)
	}
	fmt.Println("success to add 100000 row data")
}