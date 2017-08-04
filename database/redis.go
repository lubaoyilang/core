package database

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (this *Redis) New(addr, password string, db int) {
	this.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func (this *Redis) Set(key string, value interface{}, sessionTime string) error {
	var err error
	_sessionTime, err := time.ParseDuration(sessionTime)
	err = this.client.Set(key, value, _sessionTime).Err()
	return err
}
func (this *Redis) Get(key string) (interface{}, error) {
	return this.client.Get(key).Result()
}
