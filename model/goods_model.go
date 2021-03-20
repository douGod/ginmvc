package model

import (
	"LaodamingMVC/database"
)
//用户表结构
type GoodsInfo struct {
	ID int64 `gorm:"column:ID"`
	GoodsName string `gorm:"column:GoodsName"`
}

func GetGoods() *GoodsInfo{
	Db:= database.GetDb()
	goods_info := &GoodsInfo{}
	Db.First(goods_info)
	return goods_info
}