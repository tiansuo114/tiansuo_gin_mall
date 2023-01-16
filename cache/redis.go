package cache

import (
	"github.com/go-redis/redis"
	"strconv"
)

type RedisMessage struct {
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
}

var RedisClient *redis.Client

func (r *RedisMessage) BuildRedis() error {
	db, _ := strconv.ParseUint(r.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: r.RedisAddr,
		//Password:
		DB: int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	RedisClient = client
	return nil
}
