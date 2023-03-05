package service

import (
	"context"
	"log"
	pb "order/pb/github.com/six1203/order/order"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func (order *OrderService) GetOrderDetailById(ctx *context.Context, req *pb.GetOrderDetailByIdRequest) (*pb.GetOrderDetailByIdResponse, error) {
	log.Fatalf("req:%v", req)

	return nil, nil
}
