package storage

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"iris-study/src/config"
	"sync"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

var logger *zap.SugaredLogger
var db *gorm.DB = nil
var mutex = sync.Mutex{}

func GetDB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}

func init() {
	// init DB connection
	initDB()
}

func initDB() {
	mutex.Lock()
	if db != nil {
		return
	}
	defer mutex.Unlock()
	open, err := gorm.Open("postgres", "host=42.194.185.230 port=5432 user=postgres dbname=postgres password=yanghuan666 sslmode=disable")
	logger = config.GetLog()
	if err != nil {
		logger.Error("init postgres DB panic :" + err.Error())
		return
	}
	db = open
}
