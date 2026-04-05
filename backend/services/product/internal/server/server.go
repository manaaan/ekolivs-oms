package server

import (
	"context"

	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/services/product/internal/product"
	"github.com/manaaan/ekolivs-oms/backend/specs/product_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	product_api.UnimplementedProductServiceServer
	ProductService *product.Service
}

func (s Server) GetProducts(ctx context.Context, e *emptypb.Empty) (*product_api.ProductsRes, error) {
	log, ctx := tlog.New(ctx)
	log.Debug("GetProducts called")
	products, err := s.ProductService.GetProducts(ctx)
	if err != nil {
		log.Error("Unable to get products")
		return nil, err
	}

	return &product_api.ProductsRes{Products: products}, nil
}

func (s Server) UpdateProduct(ctx context.Context, prod *product_api.Product) (*product_api.Product, error) {
	log, ctx := tlog.New(ctx)
	log.Debug("UpdateProduct called", "prod", prod)
	p, err := s.ProductService.UpdateProduct(ctx, prod)
	if err != nil {
		log.Error("failed to update product", "error", err)
		return nil, err
	}

	return p, nil
}

func (s Server) GetProductByID(ctx context.Context, req *product_api.ProductIDReq) (*product_api.Product, error) {
	log, ctx := tlog.New(ctx)
	log.Debug("GetProductByID called", "req", req)
	if len(req.GetId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request, missing product ID")
	}

	p, err := s.ProductService.GetProductByID(ctx, req.GetId())
	if err != nil {
		// TODO: improve error handling
		log.Error("failed to fetch product", "error", err, "id", req.GetId())
		return nil, status.Error(codes.NotFound, "failed to fetch product")
	}

	return p, nil
}
