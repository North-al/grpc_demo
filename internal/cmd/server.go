package main

import (
	"log"
	"net"

	"github.com/North-al/grpc_demo/internal/service"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

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
