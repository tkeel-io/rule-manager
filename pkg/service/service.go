package service

import (
	"github.com/tkeel-io/kit/log"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
)

func printInputDebug(title string, in interface{}) {
	log.Debug(title, xutils.Encode2String(in))
}
