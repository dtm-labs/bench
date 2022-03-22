package svr

import (
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin"
)

func addAggregateRoutes(app *gin.Engine) {
	app.GET("/api/benchSuccess", func(c *gin.Context) {
		soMasters := []SoMaster{
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
		resp, err := dtmcli.GetRestyClient().R().SetBody(soMasters).Post(SvrUrl + "/api/createOrder")
		E2P(err)
		logger.FatalfIf(resp.StatusCode() != 200, "resp.StatusCode() != 200")
		resp, err = dtmcli.GetRestyClient().R().SetBody([]AllocateInventoryReq{{
			ProductSysNo: 1,
			Qty:          2,
		}}).Post(SvrUrl + "/api/allocateInventory")
		E2P(err)
		logger.FatalfIf(resp.StatusCode() != 200, "resp.StatusCode() != 200")
		c.JSON(200, gin.H{"status": "ok"})
	})
}
