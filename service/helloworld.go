package service

import (
	"context"
	"log"
	pb "order/pb/proto"
	"time"
)

type GreeterService struct{}

func (h *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	var datetimeFormat = "2006-01-02 15:04:05"

	log.Printf("name: %v ctx: %v content:%v datetime:%v\n", req.GetName(), ctx, req, time.Now().Format(datetimeFormat))

	reply := &pb.HelloResponse{Message: "hello"}

	return reply, nil
}
