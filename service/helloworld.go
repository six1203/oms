package service

import (
	"context"
	"log"
	pb "order/pb/github.com/six1203/order/helloworld"
	"time"
)

type GreeterService struct {
	pb.UnimplementedGreeterServiceServer
}

func (h *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	var datetimeFormat = "2006-01-02 15:04:05"

	var reqName = req.GetName()

	log.Printf("name: %v ctx: %v content:%v datetime:%v\n", reqName, ctx, req, time.Now().Format(datetimeFormat))

	reply := &pb.HelloResponse{Message: "hello " + reqName}

	return reply, nil
}
