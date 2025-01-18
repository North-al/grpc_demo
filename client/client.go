package main

import (
	"context"
	"log"

	"github.com/North-al/grpc_demo/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient err: %v", err)
	}

	defer c.Close()

	productClient := service.NewProductClient(c)
	stockResp, err := productClient.GetProductStock(context.Background(), &service.ProductStockRequest{ProductId: "1"})
	if err != nil {
		log.Fatalf("productClient.GetProductStock err: %v", err)
	}

	log.Printf("stock ProductId: %s, stock: %d", stockResp.ProductId, stockResp.Stock)
}
