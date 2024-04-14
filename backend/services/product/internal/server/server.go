package server

import (
	"context"

	"github.com/manaaan/ekolivs-oms/product/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct{}

func (Server) GetProducts(context.Context, *emptypb.Empty) (*api.ProductsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
