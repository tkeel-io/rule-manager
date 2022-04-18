package service

import (
	"testing"
)

func Test_getTargetFromExt(t *testing.T) {
	ext := `{"mapping": {"conn": {"user": "root", "database": "testverify", "endpoints": ["root:a3fks=ixmeb82a@tcp(10.96.166.237:3306)/testverify?charset=utf8&parseTime=True&loc=Local"], "sink_type": "mysql"}, "maps": [], "table": "runoob_tbl"}}`
	target, tableName, user := getTargetFromExt(ext)
	t.Log(target)
	t.Log(tableName)
	t.Log(user)
}
