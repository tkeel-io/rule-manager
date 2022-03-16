package daoutils

import (
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
)

/*
	这个package用于绕过service，对dao进行操作。实现一些特定的操作，并归档于这里。
*/

func inSlice(s []string, elem string) bool {
	for _, e := range s {
		if e == elem {
			return true
		}
	}
	return false
}

func CommitTransaction(tx *pg.Tx, err error, LogTag string, fields log.Fields) error {
	if err != nil {
		log.ErrorWithFields(LogTag, log.Fields{
			"error": err,
		})
		er := tx.Rollback()
		if er != nil {
			log.ErrorWithFields(LogTag, log.Fields{
				"desc":  "rollback error.",
				"error": er,
			})
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		log.ErrorWithFields(LogTag, log.Fields{
			"desc":  "commit error.",
			"error": err,
		})
	}
	log.InfoWithFields(LogTag, fields)

	return err
}
