package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	_ "order/global"
	"order/global/logger"
	pb "order/pb/proto"
	"order/service"
)

var (
	// 默认 8080 端口，可以在 run main.go 的时候 —p 指定为其他端口
	port = flag.Int("p", 8080, "The server port")
)

func init() {
	flag.Parse()
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		logger.Error(err)
	}

	grpcServer := grpc.NewServer()
	logger.Infof("grpc server is running at 127.0.0.1:%d ...", *port)

	// GreeterService 是我要从service文件夹导入进来的方法
	pb.RegisterGreeterServiceServer(grpcServer, &service.GreeterService{})

	pb.RegisterOrderServiceServer(grpcServer, &service.OrderService{})

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve:%v", err)
	}
}
