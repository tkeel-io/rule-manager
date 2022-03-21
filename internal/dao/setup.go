package dao

import (
	"fmt"
	"github.com/tkeel-io/rule-manager/config"
	"strings"
	"sync"

	"github.com/tkeel-io/core-broker/pkg/core"
	"github.com/tkeel-io/kit/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_once      sync.Once
	db         *gorm.DB
	CoreClient *core.Client
)

func SetCoreClientUp() (err error) {
	CoreClient, err = core.NewCoreClient()
	return
}

func Setup() error {
	// Try to create DB first.
	connectionInfo, dbName := parseConnectionAndDBName(config.DSN)
	noDBConn, err := gorm.Open(mysql.Open(connectionInfo), nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = noDBConn.Exec(createDB(dbName)).Error; err != nil {
		log.Fatal(err)
	}

	// Open the DB
	db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db.AutoMigrate(
		&Rule{},
	)
}

func DB() *gorm.DB {
	if db == nil {
		_once.Do(func() {
			if err := Setup(); err != nil {
				log.Error("Setup DB Error: ", err)
				return
			}
		})
	}
	return db
}

func parseConnectionAndDBName(dsn string) (connection, dbName string) {
	slashIndex := strings.LastIndex(dsn, "/")
	connectionInfo := dsn[:slashIndex+1]
	dbSettings := dsn[slashIndex+1:]
	questionIndex := strings.Index(dbSettings, "?")
	if questionIndex == -1 {
		questionIndex = len(dbSettings)
	}
	dbName = dbSettings[:questionIndex]
	query := dbSettings[questionIndex:]
	connection = connectionInfo + query
	return
}

const createSQLTemplate = "CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;"

func createDB(dbName string) string {
	return fmt.Sprintf(createSQLTemplate, dbName)
}
