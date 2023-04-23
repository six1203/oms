package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"order/global"
	"order/model/system"
	pb "order/pb/proto"
	"strconv"
)

type PlatformShopService struct {
	pb.UnimplementedPlatformShopServiceServer
}

func (ps *PlatformShopService) GetPlatformShopById(ctx context.Context, req *pb.GetPlatformShopByIdRequest) (*pb.GetPlatformShopByIdResponse, error) {

	db := global.GetDB()

	var platformShop system.PlatformShop
	result := db.First(&platformShop, req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 未找到
		return nil, fmt.Errorf("对象不存在：%d", req.Id)
	}
	int64Value, err := strconv.ParseInt(platformShop.PlatformShopId, 10, 64)
	if err != nil {
		// 处理错误
	}
	ShopId := int32(int64Value)

	return &pb.GetPlatformShopByIdResponse{
		PlatformShop: &pb.PlatformShop{
			Id:               int32(platformShop.Id),
			ShopId:           ShopId,
			PlatformShopName: platformShop.PlatformShopName,
			PlatformShopId:   platformShop.PlatformShopId,
			DeliveryType:     pb.DeliveryType(platformShop.DeliveryType),
			ShipmentMethod:   pb.PlatformShopShipmentMethod(platformShop.ShipmentMethod),
		},
	}, nil
}
