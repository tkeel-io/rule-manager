package action_sink

// var sinkManager *SinkManager = NewSinkManager()

// type SinkManager struct {
// 	Sinks map[string]ISink //map[sink_id]Sink
// 	lock  sync.RWMutex
// }

// func NewSinkManager() *SinkManager {
// 	return &SinkManager{
// 		Sinks: make(map[string]ISink),
// 	}
// }

// func GetSinkManager() *SinkManager {
// 	return sinkManager
// }

// func (this *SinkManager) Register(sink ISink) {
// 	this.lock.Lock()
// 	defer this.lock.Unlock()
// 	this.register(sink)
// }
// func (this *SinkManager) register(sink ISink) {
// 	if _, ok := this.Sinks[sink.Id()]; !ok {
// 		this.Sinks[sink.Id()] = sink
// 	}
// }

// func (this *SinkManager) Get(sinkType, key string) ISink {
// 	this.lock.RLock()
// 	defer this.lock.RUnlock()
// 	if sk, ok := this.Sinks[key]; ok {
// 		return sk
// 	}

// 	switch sinkType {
// 	case constant.ActionType_Chronus:
// 		//pull from redis by sink_id
// 		data, err := pullSink(sinkType, key)
// 		if nil != err {
// 			log.ErrorWithFields("[SinkManager]", log.Fields{
// 				"desc":      "pull sink failed.",
// 				"sink_type": sinkType,
// 				"sink_id":   key,
// 				"error":     err,
// 			})
// 			return nil
// 		}
// 		var sink *sink_chronus.ChronusSink
// 		sink, err = sink_chronus.NewChronusSinkByRedis(data)
// 		if nil != err {
// 			log.ErrorWithFields("[SinkManager]", log.Fields{
// 				"desc":      "call NewChronusSinkByRedis",
// 				"sink_type": sinkType,
// 				"sink_id":   key,
// 				"error":     err,
// 			})
// 			return nil
// 		}
// 		//register
// 		this.register(sink)
// 		return sink
// 	default:
// 	}
// 	return nil
// }

//将所有redis操作放到utils，endpoint只是用来初始化redis
//以Get触发的形式取redis获取sink。 然后实例生成一个sink。
