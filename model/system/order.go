package system

import (
	"fmt"
	"gorm.io/datatypes"
	"order/model/base"
	"time"
)

type OrderMainStatus int

const (
	OrderMainStatusUnknown OrderMainStatus = 0
	// OrderMainStatusUnpaid 未支付
	OrderMainStatusUnpaid OrderMainStatus = 5
	// OrderMainStatusWaitConfirm 待接单
	OrderMainStatusWaitConfirm OrderMainStatus = 10
	// OrderMainStatusConfirmed 已接单
	OrderMainStatusConfirmed OrderMainStatus = 20
	// OrderMainStatusDelivering 配送中
	OrderMainStatusDelivering OrderMainStatus = 30
	// OrderMainStatusDelivered 配送完成
	OrderMainStatusDelivered OrderMainStatus = 40
	// OrderMainStatusCanceling 取消中
	OrderMainStatusCanceling OrderMainStatus = 50
	// OrderMainStatusCanceled 已取消
	OrderMainStatusCanceled OrderMainStatus = 60
	// OrderMainStatusFinished 已完成
	OrderMainStatusFinished OrderMainStatus = 70
)

func (orderStatus OrderMainStatus) CnName() string {
	switch orderStatus {
	case OrderMainStatusUnknown:
		return "未知"
	case OrderMainStatusUnpaid:
		return "等待支付"
	case OrderMainStatusWaitConfirm:
		return "待确认"
	case OrderMainStatusConfirmed:
		return "已确认"
	case OrderMainStatusDelivering:
		return "配送中"
	case OrderMainStatusDelivered:
		return "已送达"
	case OrderMainStatusCanceling:
		return "取消中"
	case OrderMainStatusCanceled:
		return "已取消"
	case OrderMainStatusFinished:
		return "已完成"
	default:
		return fmt.Sprintf("未知状态 %d", orderStatus)
	}
}

type Order struct {
	base.Model
	PlatformOrderId             string          `gorm:"type:varchar(255);not null;default:''"`
	PlatformShopPk              int64           `gorm:"not null;default:0"`
	PlatformShopId              string          `gorm:"type:varchar(255);not null;default:''"`
	PlatformShopName            string          `gorm:"type:varchar(255);not null;default:''"`
	PlatformType                int8            `gorm:"not null;default:0"`
	CredentialType              int8            `gorm:"not null;default:0"`
	MainStatus                  OrderMainStatus `gorm:"not null;default:0"`
	MainStatusDesc              string          `gorm:"type:varchar(255);not null;default:''"`
	CreateTime                  time.Time       `gorm:"type:datetime;not null;index:create_time;default:current_timestamp"`
	ConfirmDeadline             time.Time       `gorm:"type:datetime;not null;index:confirm_deadline;default:1970-01-01 00:00:00"`
	FinishTime                  time.Time       `gorm:"type:datetime;not null;index:finish_time;default:1970-01-01 00:00:00"`
	CancelTime                  time.Time       `gorm:"type:datetime;not null;index:cancel_time;default:1970-01-01 00:00:00"`
	CancelReason                string          `gorm:"type:varchar(1024);not null;default:''"`
	UpdateTime                  time.Time       `gorm:"type:datetime;not null;index:update_time;default:1970-01-01 00:00:00"`
	ExpectedArrivalTime         time.Time       `gorm:"type:datetime;not null;index:expected_arrival_time;default:1970-01-01 00:00:00"`
	IsPreOrder                  bool            `gorm:"not null;default:false"`
	Total                       int32           `gorm:"not null;default:0"`
	UserPaid                    int32           `gorm:"not null;default:0"`
	DiscountAmount              int32           `gorm:"not null;default:0"`
	PostInsuranceAmount         int32           `gorm:"not null;default:0"`
	EstimatedIncome             int32           `gorm:"not null;default:0"`
	PlatformCommission          int32           `gorm:"not null;default:0"`
	PlatformSubsidy             int32           `gorm:"not null;default:0"`
	MerchantSubsidy             int32           `gorm:"not null;default:0"`
	DeliveryFee                 int32           `gorm:"not null;default:0"`
	OrderDeliveryFee            int32           `gorm:"not null;default:0"`
	PackingFee                  int32           `gorm:"not null;default:0"`
	OrderPackingFee             int32           `gorm:"not null;default:0"`
	PlatformUserId              string          `gorm:"type:varchar(255);not null;default:''"`
	Receiver                    string          `gorm:"type:varchar(255);not null;default:''"`
	RealMobile                  string          `gorm:"type:varchar(255);not null;default:''"`
	PrivacyNumber               string          `gorm:"type:varchar(255);not null;default:''"`
	MobileSuffix                string          `gorm:"type:varchar(255);not null;default:''"`
	Province                    string          `gorm:"type:varchar(255);not null;default:''"`
	City                        string          `gorm:"type:varchar(255);not null;default:''"`
	Town                        string          `gorm:"type:varchar(255);not null;default:''"`
	DetailAddress               string          `gorm:"type:varchar(1024);not null;default:''"`
	FullAddress                 string          `gorm:"type:varchar(1024);not null;default:''"`
	Longitude                   string          `gorm:"type:varchar(255);not null;default:''"`
	Latitude                    string          `gorm:"type:varchar(255);not null;default:''"`
	UserRemark                  string          `gorm:"type:varchar(1024);not null;default:''"`
	MerchantRemark              string          `gorm:"type:varchar(1024);not null;default:''"`
	IsAbnormal                  bool            `gorm:"not null;default:false"`
	AbnormalReason              datatypes.JSON  `gorm:"json"`
	PlatformShopDeliveryType    int8            `gorm:"not null;default:0"`
	PlatformShopShipmentMethod  int8            `gorm:"not null;default:0"`
	MerchantId                  int64           `gorm:"not null;default:0"`
	Tags                        datatypes.JSON  `gorm:"json"`
	IsSupportMultiPackLogistics bool            `gorm:"not null;default:false"`
	LogisticsAgentCode          string          `gorm:"type:varchar(255);not null;default:''"`
	UserPickupMethod            int8            `gorm:"not null;default:0"`
	UserPickupCode              string          `gorm:"type:varchar(255);not null;default:''"`
}

// TableName 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (o Order) TableName() string {
	//绑定MYSQL表名为order
	return "order"
}
