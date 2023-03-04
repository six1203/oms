package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "order/pb/proto"
	"order/service"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	log.Println("grpc server is running ...")

	// GreeterService 是我要从service文件夹导入进来的方法
	pb.RegisterGreeterServiceServer(grpcServer, &service.GreeterService{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
