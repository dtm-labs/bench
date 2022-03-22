package svr

import (
	"bytes"
	"io/ioutil"

	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
)

const SvrUrl = "http://localhost:8080"

func E2P(e error) {
	if e != nil {
		panic(e)
	}
}

// GetGinApp init and return gin
func GetGinApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(func(c *gin.Context) {
		body := ""
		if c.Request.Body != nil {
			rb, err := c.GetRawData()
			E2P(err)
			if len(rb) > 0 {
				body = string(rb)
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			}
		}
		logger.Debugf("begin %s %s body: %s", c.Request.Method, c.Request.URL, body)
		c.Next()
	})
	app.Any("/api/ping", func(c *gin.Context) { c.JSON(200, map[string]interface{}{"msg": "pong"}) })
	return app
}

func Main() {
	logger.Infof("starting bench server")
	app := GetGinApp()
	addOrderRoutes(app)
	addProductRoutes(app)
	addAggregateRoutes(app)
	logger.Infof("listening on %s", ":8080")
	logger.InitLog("warn")
	app.Run(":8080")
}
