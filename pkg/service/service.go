package service

import (
	"git.internal.yunify.com/manage/common/log"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
)

func printInputDebug(title string, in interface{}) {

	log.DebugWithFields(title, log.Fields{
		"inputs:": xutils.Encode2String(in),
	})
}
