package svr

import (
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SoMaster struct {
	Sysno                int64   `json:"sysNo"`
	SoId                 string  `json:"soID"`
	BuyerUserSysno       int64   `json:"buyerUserSysNo"`
	SellerCompanyCode    string  `json:"sellerCompanyCode"`
	ReceiveDivisionSysno int64   `json:"receiveDivisionSysNo"`
	ReceiveAddress       string  `json:"receiveAddress"`
	ReceiveZip           string  `json:"receiveZip"`
	ReceiveContact       string  `json:"receiveContact"`
	ReceiveContactPhone  string  `json:"receiveContactPhone"`
	StockSysno           int64   `json:"stockSysNo"`
	PaymentType          int32   `json:"paymentType"`
	SoAmt                float64 `json:"soAmt"`
	//10，创建成功，待支付；30；支付成功，待发货；50；发货成功，待收货；70，确认收货，已完成；90，下单失败；100已作废
	Status       int32      `json:"status"`
	OrderDate    time.Time  `json:"orderDate"`
	PaymentDate  *time.Time `json:"paymentDate"`
	DeliveryDate *time.Time `json:"deliveryDate"`
	ReceiveDate  *time.Time `json:"receiveDate"`
	Appid        string     `json:"appID"`
	Memo         string     `json:"memo"`
	CreateUser   *string    `json:"createUser"`
	GmtCreate    *time.Time `json:"gmtCreate"`
	ModifyUser   *string    `json:"modifyUser"`
	GmtModified  *time.Time `json:"gmtModified"`

	SoItems []*SoItem `gorm:"-"`
}

func (*SoMaster) TableName() string { return "dtm_bench.so_master" }

type SoItem struct {
	Sysno         int64   `json:"sysNo"`
	SoSysno       int64   `json:"soSysNo"`
	ProductSysno  int64   `json:"productSysNo"`
	ProductName   string  `json:"productName"`
	CostPrice     float64 `json:"costPrice"`
	OriginalPrice float64 `json:"originalPrice"`
	DealPrice     float64 `json:"dealPrice"`
	Quantity      int32   `json:"quantity"`
}

func (*SoItem) TableName() string { return "dtm_bench.so_item" }

type Inventory struct {
	Sysno           uint64
	ProductSysno    uint64
	AccountQty      int32
	AvailableQty    int32
	AllocatedQty    int32
	AdjustLockedQty int32
}

func (*Inventory) TableName() string { return "dtm_bench.inventory" }

func NextID() uint64 {
	id, _ := uuid.NewUUID()
	return uint64(id.ID())
}

func CreateSO(db *gorm.DB, soMasters []*SoMaster) {
	result := make([]uint64, 0, len(soMasters))
	for _, soMaster := range soMasters {
		soid := NextID()
		so_master := &SoMaster{
			Sysno:                gconv.Int64(soid),
			SoId:                 gconv.String(soid),
			BuyerUserSysno:       soMaster.BuyerUserSysno,
			SellerCompanyCode:    soMaster.SellerCompanyCode,
			ReceiveDivisionSysno: soMaster.ReceiveDivisionSysno,
			ReceiveAddress:       soMaster.ReceiveAddress,
			ReceiveZip:           soMaster.ReceiveZip,
			ReceiveContact:       soMaster.ReceiveContact,
			ReceiveContactPhone:  soMaster.ReceiveContactPhone,
			StockSysno:           soMaster.StockSysno,
			PaymentType:          soMaster.PaymentType,
			SoAmt:                soMaster.SoAmt,
			Status:               soMaster.Status,
			OrderDate:            time.Now(),
			Appid:                soMaster.Appid,
			Memo:                 soMaster.Memo,
		}
		err := db.Create(so_master).Error

		E2P(err)

		soItems := soMaster.SoItems
		for _, soItem := range soItems {
			soItemID := NextID()
			so_item := &SoItem{
				Sysno:         gconv.Int64(soItemID),
				SoSysno:       gconv.Int64(soid),
				ProductSysno:  soItem.ProductSysno,
				ProductName:   soItem.ProductName,
				CostPrice:     soItem.CostPrice,
				OriginalPrice: soItem.OriginalPrice,
				DealPrice:     soItem.DealPrice,
				Quantity:      soItem.Quantity,
			}
			err := db.Create(so_item).Error
			E2P(err)
		}
		result = append(result, soid)
	}
}

type AllocateInventoryReq struct {
	ProductSysNo int64
	Qty          int32
}

func AllocateInventory(db *gorm.DB, reqs []*AllocateInventoryReq) {

	for _, req := range reqs {
		err := db.Model(&Inventory{}).
			Where("product_sysno = ? and available_qty >= ?", req.ProductSysNo, req.Qty).
			UpdateColumns(map[string]interface{}{
				"available_qty": gorm.Expr("available_qty - ?", req.Qty),
				"allocated_qty": gorm.Expr("allocated_qty + ?", req.Qty),
			}).Error
		E2P(err)
	}
}
