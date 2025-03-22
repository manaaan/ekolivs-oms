package server

import (
	"context"
	"errors"
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
	products, err := s.ProductService.GetProducts(ctx)
	if err != nil {
		slog.Error("Unable to get products")
		return nil, err
	}

	return &api.ProductsRes{Products: products}, nil
}

func (s Server) UpdateProduct(ctx context.Context, prod *api.Product) (*api.Product, error) {
	p, err := s.ProductService.UpdateProduct(ctx, prod)
	if err != nil {
		slog.Error("failed to update product", "error", err)
		return nil, err
	}

	return p, nil
}

func (s Server) GetProductByID(ctx context.Context, req *api.ProductIDReq) (*api.Product, error) {
	if req == nil || len(req.GetID()) == 0 {
		return nil, errors.New("invalid request, missing product ID")
	}

	p, err := s.ProductService.GetProductByID(ctx, req.GetID())
	if err != nil {
		slog.Error("failed to update product", "error", err)
		return nil, err
	}

	return p, nil
}
