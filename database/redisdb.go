package database

import (
	"github.com/go-redis/redis"
	"sync"
)
var RedisClient *redis.Client
var onceRedis sync.Once
func init(){
	onceRedis.Do(func(){
		RedisClient = redis.NewClient(&redis.Options{
			Addr:"127.0.0.1:6379",
			Password:"",
			DB:0,
			PoolSize:10,//最大连接数
			MinIdleConns:2,//最小连接数
			IdleTimeout:1,//多余连接1分钟后释放
		})
	})
	if _,err:=RedisClient.Ping().Result();err != nil{
		panic(err.Error())
	}
}

