package model

import (
	"gorm.io/datatypes"
	"time"
)

type Order struct {
	Common
	PlatformOrderId             string         `gorm:"type:varchar(255);not null;default:''"`
	PlatformShopPk              int64          `gorm:"not null;default:0"`
	PlatformShopId              string         `gorm:"type:varchar(255);not null;default:''"`
	PlatformShopName            string         `gorm:"type:varchar(255);not null;default:''"`
	PlatformType                int8           `gorm:"not null;default:0"`
	CredentialType              int8           `gorm:"not null;default:0"`
	MainStatus                  int8           `gorm:"not null;default:0"`
	MainStatusDesc              string         `gorm:"type:varchar(255);not null;default:''"`
	CreateTime                  time.Time      `gorm:"not null;index:create_time"`
	ConfirmDeadline             time.Time      `gorm:"not null;index:confirm_deadline"`
	FinishTime                  time.Time      `gorm:"not null;index:finish_time"`
	CancelTime                  time.Time      `gorm:"not null;index:cancel_time"`
	CancelReason                string         `gorm:"type:text;not null"`
	UpdateTime                  time.Time      `gorm:"not null;index:update_time"`
	ExpectedArrivalTime         time.Time      `gorm:"not null;index:expected_arrival_time"`
	IsPreOrder                  bool           `gorm:"not null;default:false"`
	Total                       int32          `gorm:"not null;default:0"`
	UserPaid                    int32          `gorm:"not null;default:0"`
	DiscountAmount              int32          `gorm:"not null;default:0"`
	PostInsuranceAmount         int32          `gorm:"not null;default:0"`
	EstimatedIncome             int32          `gorm:"not null;default:0"`
	PlatformCommission          int32          `gorm:"not null;default:0"`
	PlatformSubsidy             int32          `gorm:"not null;default:0"`
	MerchantSubsidy             int32          `gorm:"not null;default:0"`
	DeliveryFee                 int32          `gorm:"not null;default:0"`
	OrderDeliveryFee            int32          `gorm:"not null;default:0"`
	PackingFee                  int32          `gorm:"not null;default:0"`
	OrderPackingFee             int32          `gorm:"not null;default:0"`
	PlatformUserId              string         `gorm:"type:varchar(255);not null;default:''"`
	Receiver                    string         `gorm:"type:varchar(255);not null;default:''"`
	RealMobile                  string         `gorm:"type:varchar(255);not null;default:''"`
	PrivacyNumber               string         `gorm:"type:varchar(255);not null;default:''"`
	MobileSuffix                string         `gorm:"type:varchar(255);not null;default:''"`
	Province                    string         `gorm:"type:varchar(255);not null;default:''"`
	City                        string         `gorm:"type:varchar(255);not null;default:''"`
	Town                        string         `gorm:"type:varchar(255);not null;default:''"`
	DetailAddress               string         `gorm:"type:text;not null"`
	FullAddress                 string         `gorm:"type:text;not null"`
	Longitude                   string         `gorm:"type:varchar(255);not null;default:''"`
	Latitude                    string         `gorm:"type:varchar(255);not null;default:''"`
	UserRemark                  string         `gorm:"type:text;not null"`
	MerchantRemark              string         `gorm:"type:text;not null"`
	IsAbnormal                  bool           `gorm:"not null;default:false"`
	AbnormalReason              datatypes.JSON `gorm:"json"`
	PlatformShopDeliveryType    int8           `gorm:"not null;default:0"`
	PlatformShopShipmentMethod  int8           `gorm:"not null;default:0"`
	MerchantId                  int64          `gorm:"not null;default:0"`
	Tags                        datatypes.JSON `gorm:"json"`
	IsSupportMultiPackLogistics bool           `gorm:"not null;default:false"`
	LogisticsAgentCode          string         `gorm:"type:varchar(255);not null;default:''"`
	UserPickupMethod            int8           `gorm:"not null;default:0"`
	UserPickupCode              string         `gorm:"type:varchar(255);not null;default:''"`
}

// TableName 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (o Order) TableName() string {
	//绑定MYSQL表名为order
	return "order"
}
