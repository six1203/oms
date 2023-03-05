package model

type PlatformShop struct {
	Common
	PlatformShopId   string `gorm:"type:varchar(255);comment:平台门店ID;index:idx_platform_shop_id"`
	PlatformShopName string `gorm:"type:varchar(255);comment:平台门店名称"`
	PlatformType     int8   `gorm:"comment:平台类型;default:0"`
	DeliveryType     int8   `gorm:"comment:配送类型;default:0"`
	ShipmentMethod   int8   `gorm:"comment:配送方式;default:0"`
}

func (s *PlatformShop) TableName() string {
	return "platform_shop"
}
