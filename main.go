package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "order/pb/proto"
	"order/service"
	"order/tools/logger"
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
	log.Printf("grpc server is running on 127.0.0.1:%d ...", *port)

	// GreeterService 是我要从service文件夹导入进来的方法
	pb.RegisterGreeterServiceServer(grpcServer, &service.GreeterService{})

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve:%v", err)
	}
}
