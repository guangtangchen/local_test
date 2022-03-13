package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisClient *redis.Client

func initRedis(ctx context.Context) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "82.156.205.9:6379",
		Password: "821562069",
	})
	resp, err := redisClient.Ping().Result()
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	resp, err = redisClient.Get("masterkey").Result()
	if err != nil || resp != "20210221" {
		panic("redis read fail")
	}
	fmt.Println("redis connected")
}

func getRedisClient(ctx context.Context) *redis.Client {
	return redisClient.WithContext(ctx)
}

func Set(ctx context.Context, key string, value interface{}, expire time.Duration) (string, error) {
	// if success, respString = ok, respError = nil
	return getRedisClient(ctx).Set(key, value, expire).Result()
}
