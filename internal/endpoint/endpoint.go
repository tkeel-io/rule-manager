package endpoint

import (
	"sync"
)

var metadataServiceName = "mdmp-rule-manage"
var authOnce sync.Once

//initialize all endpoint.
func Init() {
	//	initRule()
	// InitRedis()
	initMetaEnd()
}
