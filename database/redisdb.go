package database

import (
	"github.com/go-redis/redis"
	"sync"
)
var redisConn *redis.Client
var onceRedis sync.Once
func connectRedis() error {
	onceRedis.Do(func(){
		redisConn = redis.NewClient(&redis.Options{
			Addr:"127.0.0.1:6379",
			Password:"",
			DB:0,
			PoolSize:10,//最大连接数
			MinIdleConns:2,//最小连接数
			IdleTimeout:1,//多余连接1分钟后释放
		})
	})
	if _,err:=redisConn.Ping().Result();err != nil{
		return err
	}
	return nil
}
func GetRedisDb() (redis.Client,error){
	if redisConn == nil{
		if err := connectRedis();err != nil{
			return redis.Client{}, err
		}
	}
	return *redisConn,nil
}
