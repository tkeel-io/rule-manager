package endpoint

import (
	"time"

	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/rule-manager/config"
	redis "github.com/tkeel-io/rule-manager/internal/dao"
)

const redisClientLogTitle = "[RedisClient]"

type RedisEndpoint struct{}

func InitRedis() {
	redis.InitRedis(config.GetConfig().Redis.Addr, config.GetConfig().Redis.Password, config.GetConfig().Redis.DB)
}

func GetRedisEndpoint() (r *RedisEndpoint) {
	return &RedisEndpoint{}
}

func (r *RedisEndpoint) Set(key, value string, expiration time.Duration) error {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:":   key,
		"value:": value,
	})
	return redis.GetRedis().Set(key, value, expiration).Err()
}

func (r *RedisEndpoint) Get(key string) (string, error) {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:": key,
	})
	return redis.GetRedis().Get(key).Result()
}

func (r *RedisEndpoint) HSet(key, field string, values interface{}) error {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:":  key,
		"field": field,
		//"value:": values,
	})
	return redis.GetRedis().HSet(key, field, values).Err()
}

func (r *RedisEndpoint) HGet(key, field string) (string, error) {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:":  key,
		"field": field,
	})
	return redis.GetRedis().HGet(key, field).Result()
}

func (r *RedisEndpoint) HGetAll(key string) (map[string]string, error) {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:": key,
	})
	return redis.GetRedis().HGetAll(key).Result()
}

func (r *RedisEndpoint) HDel(key string, fields ...string) error {
	log.Debug(redisClientLogTitle, map[string]interface{}{
		"key:":  key,
		"field": fields,
	})
	return redis.GetRedis().HDel(key, fields...).Err()
}
