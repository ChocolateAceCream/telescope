package utils

import (
	"context"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	redisConfig := singleton.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		singleton.Logger.Error("redis connect ping failed", zap.Error(err))
		singleton.Redis = nil
	} else {
		singleton.Logger.Info("redis connect ping success, response is: ", zap.String("pong", pong))
		singleton.Redis = client
	}
}
