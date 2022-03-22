package svr

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbOnce sync.Once
var db *gorm.DB

func dbGet() *gorm.DB {
	var err error
	dbOnce.Do(func() {
		db, err = gorm.Open(mysql.Open("root:@tcp(localhost:3306)/dtm_bench?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true"), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		E2P(err)
	})
	return db
}
