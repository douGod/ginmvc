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
	goods_info := &GoodsInfo{}
	database.MysqlDb.First(goods_info)
	return goods_info
}