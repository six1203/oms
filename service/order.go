package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"order/global"
	"order/global/logger"
	"order/model"
	pb "order/pb/proto"
	"order/tools"
	"strconv"
	"strings"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func (o *OrderService) GetOrderDetailById(ctx context.Context, req *pb.GetOrderDetailByIdRequest) (*pb.GetOrderDetailByIdResponse, error) {

	db := global.GetDB()

	var order model.Order
	result := db.First(&order, 1)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("对象不存在：%d", req.OrderId)
	}

	// FIXME: 这个struct没有写完，太多了，懒得写
	pbOrder := pb.Order{
		PlatformOrderId: order.PlatformOrderId,
		PlatformType:    pb.PlatformType(order.PlatformType),
		MainStatus:      pb.OrderMainStatus(order.MainStatus),
		MainStatusDesc:  order.MainStatusDesc,
	}
	return &pb.GetOrderDetailByIdResponse{Order: &pbOrder}, nil
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	logger.Infof("收到创建订单的请求：%v", req)

	db := global.GetDB()

	var ps model.PlatformShop
	result := db.Where("platform_shop_id = ?", req.PlatformShopId).First(&ps)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("门店不存在：%s", req.PlatformShopId)
	}

	regeocode, err := tools.Geocode(req.DetailAddress)
	if err != nil {
		return nil, err
	}

	platformOrderId := tools.GenerateId()

	platformUserId := tools.GenerateId()

	coordinate := strings.Split(regeocode.Location, ",")

	order := model.Order{
		PlatformOrderId:  strconv.FormatInt(platformOrderId, 10),
		PlatformShopPk:   ps.Id,
		PlatformShopId:   ps.PlatformShopId,
		PlatformShopName: ps.PlatformShopName,
		PlatformType:     ps.PlatformType,
		MainStatus:       model.ORDER_MAIN_STATUS_WAIT_CONFIRM,
		// FIXME 通过mapping表来获取
		MainStatusDesc:      "待接单",
		CreateTime:          tools.GetNowTime(),
		ConfirmDeadline:     tools.GetNowTimeAddMinute(5),
		UpdateTime:          tools.GetUnixEpoch(),
		FinishTime:          tools.GetUnixEpoch(),
		CancelTime:          tools.GetUnixEpoch(),
		ExpectedArrivalTime: tools.GetNowTimeAddMinute(60),
		Total:               req.Total,
		UserPaid:            req.UserPaid,
		DiscountAmount:      req.DiscountAmount,
		PlatformUserId:      strconv.FormatInt(platformUserId, 10),
		Receiver:            req.Receiver,
		RealMobile:          req.RealMobile,
		DetailAddress:       req.DetailAddress,
		FullAddress:         req.DetailAddress,
		Province:            regeocode.Province,
		City:                regeocode.City,
		Town:                regeocode.District,
		Longitude:           coordinate[0],
		Latitude:            coordinate[1],
		UserRemark:          req.UserRemark,
	}

	db.Create(&order)

	pbOrder := pb.Order{
		PlatformOrderId: order.PlatformOrderId,
		PlatformType:    pb.PlatformType(order.PlatformType),
		MainStatus:      pb.OrderMainStatus(order.MainStatus),
		MainStatusDesc:  order.MainStatusDesc,
	}
	return &pb.CreateOrderResponse{Order: &pbOrder}, nil
}
