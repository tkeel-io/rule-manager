package dao

import (
	"context"
	"fmt"
	"testing"

	daoutils "github.com/tkeel-io/rule-manager/internal/dao/utils"
	"git.internal.yunify.com/manage/common/db"
)

func TestMain(m *testing.M) {
	db.InitPG("192.168.14.102", "5432", "yunify", "zhu88jie", "iot")
	//check pg table
	err := db.CheckTables([]interface{}{
		(*Action)(nil),
		(*Rule)(nil),
	})
	if err != nil {
		panic(err)
	}
	Init()
	m.Run()
}

func TestRuleQuery(t *testing.T) {

	rule := Rule{}

	cond := daoutils.RuleQueryReq{
		UserId: "usr-r9Tis0yJ",
	}

	rules, err := rule.Query(context.Background(), &cond)
	fmt.Println(rules)
	fmt.Println(err)

}
