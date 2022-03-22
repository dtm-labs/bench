package svr

import (
	"bytes"
	"io/ioutil"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
)

const SvrUrl = "http://localhost:8080"
const DtmServer = "http://localhost:36789/api/dtmsvr"

var DbConf = &dtmcli.DBConf{
	Driver:   "mysql",
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "",
}

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

var XaClient *dtmcli.XaClient

func Main() {
	logger.Infof("starting bench server")
	app := GetGinApp()
	addOrderRoutes(app)
	addProductRoutes(app)
	addAggregateRoutes(app)
	var err error
	XaClient, err = dtmcli.NewXaClient(DtmServer, *DbConf, SvrUrl+"/api/xa", func(path string, xa *dtmcli.XaClient) {
		app.POST(path, func(c *gin.Context) {
			r := xa.HandleCallback(c.Query("gid"), c.Query("branch_id"), c.Query("op"))
			if err, ok := r.(error); ok {
				c.JSON(500, gin.H{"error": err.Error()})
			} else {
				c.JSON(200, r)
			}
		})
	})
	logger.FatalIfError(err)
	logger.Infof("listening on %s", ":8080")
	logger.InitLog("debug")
	app.Run(":8080")
}
