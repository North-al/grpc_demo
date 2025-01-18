package service

import "context"

// var ProductService = &productService{}

type ProductService struct {
	UnimplementedProductServer
}

func (s *ProductService) GetProductStock(ctx context.Context, in *ProductStockRequest) (*ProductStockResponse, error) {
	stock := s.GetStockById(in.ProductId)

	return &ProductStockResponse{Stock: stock, ProductId: in.ProductId}, nil
}

func (s *ProductService) GetStockById(id string) int32 {
	// 维护一个库存表
	stock := map[string]int32{
		"1": 100,
		"2": 200,
		"3": 300,
	}

	return stock[id]
}
