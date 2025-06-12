package server

import (
	"context"

	"github.com/manaaan/ekolivs-oms/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/product/api"
	"github.com/manaaan/ekolivs-oms/product/internal/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	api.UnimplementedProductServiceServer
	ProductService *product.Service
}

func (s Server) GetProducts(ctx context.Context, e *emptypb.Empty) (*api.ProductsRes, error) {
	log, ctx := tlog.New(ctx)
	products, err := s.ProductService.GetProducts(ctx)
	if err != nil {
		log.Error("Unable to get products")
		return nil, err
	}

	return &api.ProductsRes{Products: products}, nil
}

func (s Server) UpdateProduct(ctx context.Context, prod *api.Product) (*api.Product, error) {
	log, ctx := tlog.New(ctx)
	p, err := s.ProductService.UpdateProduct(ctx, prod)
	if err != nil {
		log.Error("failed to update product", "error", err)
		return nil, err
	}

	return p, nil
}

func (s Server) GetProductByID(ctx context.Context, req *api.ProductIDReq) (*api.Product, error) {
	log, ctx := tlog.New(ctx)
	if len(req.GetID()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request, missing product ID")
	}

	p, err := s.ProductService.GetProductByID(ctx, req.GetID())
	if err != nil {
		// TODO: improve error handling
		log.Error("failed to fetch product", "error", err, "id", req.GetID())
		return nil, status.Error(codes.NotFound, "failed to fetch product")
	}

	return p, nil
}
