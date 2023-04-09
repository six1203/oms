package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"order/global"
	"order/global/logger"
	"order/model/request"
	"order/model/system"
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

	var order system.Order
	result := db.First(&order, req.OrderId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("对象不存在：%d", req.OrderId)
	}

	// FIXME: 这个struct没有写完，太多了，懒得写
	pbOrder := pb.Order{
		PlatformOrderId: order.PlatformOrderId,
		PlatformType:    pb.PlatformType(order.PlatformType),
		MainStatus:      pb.OrderMainStatus(order.MainStatus),
		MainStatusDesc:  order.MainStatusDesc,
		OrderId:         order.Id,
		CreateTime:      tools.TimeToTimestamp(order.CreateTime),
	}
	return &pb.GetOrderDetailByIdResponse{Order: &pbOrder}, nil
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	logger.Infof("收到创建订单的请求：%v", req)

	db := global.GetDB()

	var ps system.PlatformShop
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

	order := system.Order{
		PlatformOrderId:     strconv.FormatInt(platformOrderId, 10),
		PlatformShopPk:      ps.Id,
		PlatformShopId:      ps.PlatformShopId,
		PlatformShopName:    ps.PlatformShopName,
		PlatformType:        ps.PlatformType,
		MainStatus:          system.ORDER_MAIN_STATUS_WAIT_CONFIRM,
		MainStatusDesc:      system.ORDER_MAIN_STATUS_WAIT_CONFIRM.CnName(),
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
		OrderId:         order.Id,
	}
	return &pb.CreateOrderResponse{Order: &pbOrder}, nil
}

func (o *OrderService) ListSimpleOrder(ctx context.Context, req *pb.ListSimpleOrderRequest) (*pb.ListSimpleOrderResponse, error) {

	validate := validator.New()
	err := validate.Struct(
		request.PaginationRequest{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	)
	if err != nil {
		return nil, err
	}

	db := global.GetDB()
	var orders []system.Order
	db.Order("create_time desc").Offset(0).Limit(100).Find(&orders)

	logger.Infof("orders", orders)

	var total int64
	db.Order("create_time desc").Count(&total)

	var simpleOrders []*pb.SimpleOrder

	for _, order := range orders {

		simpleOrder := pb.SimpleOrder{
			OrderId:         order.Id,
			PlatformOrderId: order.PlatformOrderId,
		}

		simpleOrders = append(simpleOrders, &simpleOrder)
	}

	return &pb.ListSimpleOrderResponse{
		SimpleOrders: simpleOrders,
		Total:        int32(total),
	}, nil

}
