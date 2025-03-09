package initializer

import (
	"context"
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	global.Redis = rdb
	testRedis()
	global.Logger.Info("Connect Redis Success")
}

var ctx = context.Background()

func testRedis() {
	rdb := global.Redis
	err := rdb.Set(ctx, "count", 12, 0).Err()
	if err != nil {
		panic(fmt.Sprintf("Redis Client Error %v", err))
	}

	val, err := rdb.Get(ctx, "count").Result()

	if err != nil {
		panic(fmt.Sprintf("Redis Client Error %v", err))
	}

	fmt.Printf("Value %v", val)
}
