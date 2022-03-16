package utils

//用于对rules进行存储
type SelectField struct {
	//thing {property|event} identifier.
	Expr  string
	Alias string
	Type  string
}

type RuleUpdateReq struct {
	Id           string
	UserId       string
	Name         *string
	Status       *string
	RuleDesc     *string
	DataType     *uint8
	SelectText   *string
	SelectFields []*SelectField
	TopicType    *string
	ShortTopic   *string
	WhereText    *string
	Ruleql       string
	Raw          *string
}

type RuleDeleteReq struct {
	Id     string
	UserId string
}

type RuleQueryReq struct {
	Id           *string
	Ids          []string
	UserId       string
	Name         *string
	DataType     *uint8
	TopicType    *string
	ShortTopic   *string
	Page         *Pager
	SearchKey    *string
	FlagQueryBan bool
}
