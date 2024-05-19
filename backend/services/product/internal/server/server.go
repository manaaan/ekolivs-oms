package server

import (
	"context"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/product/api"
	"github.com/manaaan/ekolivs-oms/product/internal/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	api.UnimplementedProductServiceServer
	ProductService *product.Service
}

func (s Server) GetProducts(ctx context.Context, e *emptypb.Empty) (*api.ProductsRes, error) {
	products, err := s.ProductService.GetProducts()
	if err != nil {
		slog.Error("Unable to get products")
		return nil, err
	}

	return &api.ProductsRes{Products: products}, nil
}
