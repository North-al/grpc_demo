package main

import (
	"context"
	"log"
	"net"

	"github.com/North-al/grpc_demo/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {

	var authInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log.Println("auth Interceptor")

		// 模拟认证
		incomingContext, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
		}

		// 获取认证信息
		auth, ok := incomingContext["authorization"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
		}

		if auth[0] != "Bearer token" {
			return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
		}

		// 继续处理请求
		return handler(ctx, req)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))

	// 注册服务
	service.RegisterProductServer(server, &service.ProductService{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}

}
