package utils

type ActionUpdateReq struct {
	Id              string
	UserId          string
	RuleId          *string
	SinkId          *string
	Name            *string
	Status          *string
	ConfigStatus    *bool
	Configuration   map[string]interface{}
	ActionType      *string
	ErrorActionFlag *bool
}

type ActionDeleteReq struct {
	Id     string
	UserId string
	RuleId string
}

type ActionQueryReq struct {
	Id              *string
	Ids             []string
	UserId          string
	RuleId          *string
	SinkId          *string
	Name            *string
	ConfigStatus    *bool
	ActionType      *string
	ErrorActionFlag *bool
	Page            *Pager
	SearchKey       *string
	FlagQueryBan    bool
}
