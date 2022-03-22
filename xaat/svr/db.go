package svr

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbOnce sync.Once
var db *gorm.DB

func dbGet() *gorm.DB {
	var err error
	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true",
			DbConf.User, DbConf.Password, DbConf.Host, DbConf.Port, "")
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		E2P(err)
	})
	return db
}
