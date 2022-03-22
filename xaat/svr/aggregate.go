package svr

import (
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/lithammer/shortuuid"
)

var soMasters = []SoMaster{
	{
		BuyerUserSysno:       10001,
		SellerCompanyCode:    "SC001",
		ReceiveDivisionSysno: 110105,
		ReceiveAddress:       "朝阳区长安街001号",
		ReceiveZip:           "000001",
		ReceiveContact:       "斯密达",
		ReceiveContactPhone:  "18728828296",
		StockSysno:           1,
		PaymentType:          1,
		SoAmt:                430.5,
		Status:               10,
		Appid:                "dk-order",
		SoItems: []*SoItem{
			{
				ProductSysno:  1,
				ProductName:   "刺力王",
				CostPrice:     200,
				OriginalPrice: 232,
				DealPrice:     215.25,
				Quantity:      2,
			},
		},
	},
}

func addAggregateRoutes(app *gin.Engine) {
	app.GET("/api/benchSuccess", func(c *gin.Context) {
		gid := shortuuid.New()
		err := XaClient.XaGlobalTransaction(gid, func(xa *dtmcli.Xa) (*resty.Response, error) {
			_, err := xa.CallBranch(soMasters, SvrUrl+"/api/createOrder")
			E2P(err)
			_, err = xa.CallBranch([]AllocateInventoryReq{{
				ProductSysNo: 1,
				Qty:          2,
			}}, SvrUrl+"/api/allocateInventory")
			E2P(err)
			return nil, nil
		})
		E2P(err)
		c.JSON(200, gin.H{"status": "ok"})
	})
}
