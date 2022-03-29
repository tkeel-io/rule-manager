package dao

import (
	"sync"

	"github.com/tkeel-io/kit/log"

	"github.com/go-redis/redis"
)

var client *redis.Client
var once sync.Once

func InitRedis(addr, password string, db int) {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,       // use default DB
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Errorf("connect to redis fail addr:%s,db:%d", addr, db)
			panic(err)
		}
		//log.InfoWithFields("connect to redis success addr:%s,db:%d", log.Fields{"redis": pong})
		log.Infof("connect to redis success addr:%s,db:%d", addr, db)
	})
}

func GetRedis() *redis.Client {
	return client
}
