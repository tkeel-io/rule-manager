package utils

import (
	"github.com/go-pg/pg/orm"
)

type Pager struct {
	Limit   *int32
	Offset  *int32
	SortKey *string
	Reverse *bool
}

func WherePage(query *orm.Query, page *Pager) {
	if nil != page {
		reverse := false
		//reverve....
		if nil != page.Reverse {
			//not support.
			reverse = *page.Reverse
		}
		if nil != page.Limit {
			query.Limit(int(*page.Limit))
		}
		if nil != page.SortKey {
			//query.Order(*page.SortKey)
			if reverse {
				query.Order(*page.SortKey + " ASC")
			} else {
				query.Order(*page.SortKey + " DESC")
			}
		}
		if nil != page.Offset {
			query.Offset(int(*page.Offset))
		}
	}
}
