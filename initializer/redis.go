package initializer

import (
	"github.com/go-redis/redis/v7"
)

func NewRedis(ev *EV) (client *redis.Client, err error) {
	dsn := ev.RedisDSN
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return
	}

	return
}
