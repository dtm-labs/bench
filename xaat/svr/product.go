package svr

import "github.com/gin-gonic/gin"

func addProductRoutes(app *gin.Engine) {
	app.POST("/api/allocateInventory", func(c *gin.Context) {
		var invReq []*AllocateInventoryReq
		err := c.BindJSON(&invReq)
		E2P(err)
		AllocateInventory(dbGet(), invReq)
		c.JSON(200, gin.H{"status": "ok"})
	})
}
