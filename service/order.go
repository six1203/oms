package service

import (
	"context"
	"log"
	pb "order/pb/proto"
)

type OrderService struct {
}

func (order *OrderService) GetOrderDetailById(ctx *context.Context, req *pb.GetOrderDetailByIdRequest) (*pb.GetOrderDetailByIdResponse, error) {
	log.Fatalf("req:%v", req)

	return nil, nil
}
