package session

import (
	"thelark.cn/cms/core/database"
)

var redis *database.Redis
var maxLifeTime string

func init() {
	redis = new(database.Redis)
}

func New(addr, password, db, maxLifeTime interface{}) {
	redis.New(addr.(string), password.(string), db.(int))
	maxLifeTime = maxLifeTime.(string)
}

func Set(key string, value interface{}) error {
	return redis.Set(key, value, maxLifeTime)
}

func Get(key string) (interface{}, error) {
	return redis.Get(key)
}
