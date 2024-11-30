package redisconn

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Redis *RedisClient

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func InitRedis() *RedisClient {
	// TODO: コンテキストでデッドラインなど様々なパラメータを設定
	ctx := context.Background()
	// TODO: 環境変数で設定する
	client := redis.NewClient(&redis.Options{
		Addr:     "food_shuffle_redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	err = client.FlushAll(ctx).Err()
	if err != nil {
		panic(err)
	}

	Redis = &RedisClient{
		client: client,
		ctx:    ctx,
	}
	return Redis
}
