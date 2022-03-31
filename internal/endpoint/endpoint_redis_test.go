package endpoint

import (
	"testing"
	"time"

	redis "github.com/tkeel-io/rule-manager/internal/dao"
)

func TestRedisEndpoint_Set(t *testing.T) {
	redis.InitRedis("127.0.0.1:6379", "qingcloud2019", 4)
	rc := GetRedisEndpoint()
	rc.Set("key1", "ket2", time.Hour)
}
