package svr

import (
	"database/sql"

	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func addOrderRoutes(app *gin.Engine) {
	app.POST("/api/createOrder", func(c *gin.Context) {
		err := dtmcli.XaLocalTransaction(c.Request.URL.Query(), *DbConf, func(db *sql.DB, xa *dtmcli.Xa) error {
			var order []*SoMaster
			err := c.BindJSON(&order)
			E2P(err)
			gdb, err := gorm.Open(mysql.New(mysql.Config{
				Conn: db,
			}), &gorm.Config{SkipDefaultTransaction: true})
			E2P(err)
			CreateSO(gdb, order)
			return nil
		})
		E2P(err)
		c.JSON(200, gin.H{"status": "ok"})
	})
}
