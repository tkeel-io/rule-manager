package constant

type EventType = uint8

const (
	EVENT_RULE_CREATE EventType = iota + 1
	EVENT_RULE_DELETE
	EVENT_RULE_UPDATE
	EVENT_ACTION_CREATE
	EVENT_ACTION_UPDATE
	EVENT_ACTION_DELETE

	EVENT_UNKOWN
)

//ActionType
const (
	ActionType_Republish  = "republish"
	ActionType_Kafka      = "kafka"
	ActionType_Bucket     = "bucket"
	ActionType_Chronus    = "chronus"
	ActionType_MYSQL      = "mysql"
	ActionType_POSTGRESQL = "postgresql"
	ActionType_REDIS      = "redis"
)

var ActionTypes = []string{
	ActionType_Republish,
	ActionType_Kafka,
	ActionType_Bucket,
	ActionType_Chronus,
	ActionType_MYSQL,
	ActionType_POSTGRESQL,
	ActionType_REDIS,
}

const (
	TopicTypeProperty = "property"
	TopicTypeEvent    = "event"
	TopicTypeAll      = "+"
	TopicTypeRaw      = "raw"
)

var TopicTypes = []string{
	TopicTypeProperty,
	TopicTypeEvent,
	TopicTypeAll,
	TopicTypeRaw,
}

//---------------------errors------------------

//---------event name-----------------------
const (
	EventName           = "name-event-"
	EventDefaultName    = EventName + "default"
	EventNameRuleStatus = EventName + "rule-status"
	EventNameRuleActive = EventName + "rule-active"
)

//rule status
const (
	//action status
	ActionStatusBan       = "ban"
	ActionStatusStop      = "stop"
	ActionStatusStating   = "stating"
	ActionStatusStopping  = "stopping"
	ActionStatusRunning   = "running"
	ActionStatusError     = "error"
	ActionStatusException = "exception"
	// RuleStart            = status.RuleStart            //= "RuleStart"
	// RuleStartError       = status.RuleStartError       //= "RuleStartError"
	// RuleActionStart      = status.RuleActionStart      //= "RuleActionStart"
	// RuleActionStartError = status.RuleActionStartError //= "RuleActionStartError"
	// RuleStarted          = status.RuleStarted          //= "RuleStarted"
	// RuleStoped           = status.RuleStoped           //= "RuleStoped"
	// RuleActionError      = status.RuleActionError      //= "RuleActionError"
	// RuleActionFail       = status.RuleActionFail       //= "RuleActionFail"

	//rule在metadata创建， 等待异步消息
	RuleStart = "RuleStart"
	//rule启动失败，rule未启动
	RuleStartError       = "RuleStartError"
	RuleActionStart      = "RuleActionStart"
	RuleActionStartError = "RuleActionStartError"
	RuleStarted          = "RuleStarted"
	RuleStoped           = "RuleStoped"
	RuleActionError      = "RuleActionError"
	RuleActionFail       = "RuleActionFail"
	RuleActionStarted    = "RuleActionStarted"
)

const (
	RuleStatusBan = iota + 1
	RuleStatusStop
	RuleStatusRunning
	RuleStatusException
	RuleStatusError
	RuleStatusStating
	RuleStatusStopping
)

//由用户发出的指令状态
const (
	CommandStatusRuleStart = RuleStatusStating
	CommandStatusRuleStop  = RuleStatusStop
)

const ErrorPrefix = "RuleErr-"

//action的configuration里的字段
const MappingInfoKey = "mapping"
