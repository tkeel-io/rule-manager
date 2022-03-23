package action_sink

/*

sink:
	1. 可以连接。
	2. 可以存储映射
	3. 可以获取映射

*/

// type ISink interface {
// 	Connect(context.Context) error
// 	Type() string
// 	ID() string
// 	Ready() bool
// 	Close() error
// }

// func pullSink(sinkType, key string) ([]byte, error) {
// 	var (
// 		err   error
// 		value string
// 	)
// 	switch sinkType {
// 	case "chronus":
// 		value, err = redis.GetRedis().HGet(sink_chronus.SinkChronusHashTableKey, key).Result()
// 		if nil != err {
// 			log.ErrorWithFields("[SinkPull]", log.Fields{
// 				"desc":  "pull sink from redis failed.",
// 				"error": err,
// 			})
// 		}
// 		return []byte(value), err
// 	default:
// 		err = errors.New("sink type not support")
// 		log.ErrorWithFields("[SinkPull]", log.Fields{
// 			"desc":  "pull sink from redis failed, sink type not support.",
// 			"error": err,
// 		})
// 		return nil, err
// 	}
// }
