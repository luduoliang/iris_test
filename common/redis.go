package common

import (
	"github.com/go-redis/redis"
	"iris_test/config"
	"time"
)

var (
	Redis = NewRedisClient()
)

//返回redis，单例
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.Addr", "127.0.0.1:6379"),
		Password: config.GetString("redis.Password", ""), // no password set
		DB: int(config.GetInt64("redis.DB", 0)),   // 因为系统是64位的，所以默认的 int 型是 int64
	})
}

//获取
func GetRedis(key string) *redis.StringCmd {
	return Redis.Get(key)
}

//设置
func SetRedis(key string, value interface{}, expirse time.Duration) {
	Redis.Set(key, value, expirse)
}

//删除
func DelRedis(key string) {
	Redis.Del(key)
}

//清空
func ClearRedis() {
	Redis.FlushAll()
}
