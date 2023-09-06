package redis

import (
	"context"

	gredis "github.com/redis/go-redis/v9"
)

type RedisDb struct {
	context context.Context
	client  *gredis.Client
}

func NewRedisClient() *RedisDb {
	ctx := context.Background()
	cli := gredis.NewClient(&gredis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisDb{context: ctx, client: cli}
}

func (rdb *RedisDb) Get(key string) (any, error) {
	val, err := rdb.client.Get(rdb.context, key).Result()
	return val, err
}

func (rdb *RedisDb) Set(key string, value any) error {
	_, err := rdb.client.Set(rdb.context, key, value, 0).Result()
	return err
}
