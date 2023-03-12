package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
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

// 注册服务
func registerServices(server *grpc.Server) {
	pb.RegisterGreeterServiceServer(server, &service.GreeterService{})

	pb.RegisterOrderServiceServer(server, &service.OrderService{})
}

// FIXME 日志拦截器, 可以考虑直接使用第三方的日志拦截器
func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Unary request method: %s, req:%s", info.FullMethod, req)
	resp, err := handler(ctx, req)
	isOk := true
	if err != nil {
		isOk = false
	}
	log.Printf("Unary response method: %s, is_ok: %v", info.FullMethod, isOk)
	return resp, err
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		logger.Error(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingUnaryInterceptor))
	logger.Infof("grpc server is running at 127.0.0.1:%d ...", *port)
	registerServices(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve:%v", err)
	}
}
