package svr

import "github.com/gin-gonic/gin"

func addOrderRoutes(app *gin.Engine) {
	app.POST("/api/createOrder", func(c *gin.Context) {
		var order []*SoMaster
		err := c.BindJSON(&order)
		E2P(err)
		CreateSO(dbGet(), order)
		c.JSON(200, gin.H{"status": "ok"})
	})
}
