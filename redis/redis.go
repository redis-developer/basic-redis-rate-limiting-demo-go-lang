package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
}

func (r Redis) Set(key, value string, expire time.Duration) error {
	return r.client.Set(key, value, expire).Err()
}

func (r Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r Redis) Inc(key string) error {
	return r.client.Incr(key).Err()
}

func (r Redis) Close() error {
	return r.client.Close()
}

func NewRedis(config Config) *Redis {

	opt := &redis.Options{
		Addr:     config.Addr(),
		Password: config.Password(),
	}

	client := redis.NewClient(opt)

	r := &Redis{
		client: client,
	}

	return r
}
