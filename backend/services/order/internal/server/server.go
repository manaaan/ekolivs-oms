package server

import (
	"context"
	"errors"

	"github.com/manaaan/ekolivs-oms/backend/pkg/errkit"
	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/services/order/internal/order"
	"github.com/manaaan/ekolivs-oms/backend/services/order/order_api"
)

type Server struct {
	order_api.UnimplementedOrderServiceServer
	OrderService *order.Service
}

func (s Server) CreateOrder(ctx context.Context, create *order_api.Order) (*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	order, err := s.OrderService.CreateOrder(ctx, create)
	if err != nil {
		log.Error("failed to create order", "error", err)
		return nil, err
	}

	return order, nil
}

func (s Server) GetOrderByID(ctx context.Context, req *order_api.OrderIDReq) (*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	if req == nil || len(req.ID) == 0 {
		err := &errkit.ErrBadRequest{Err: errors.New("missing order ID on request")}
		log.Warn(err.Error(), "req", req)
		return nil, err
	}
	order, err := s.OrderService.GetOrderByID(ctx, req.ID)
	if err != nil {
		log.Error("failed to get order by ID", "error", err, "req", req)
		return nil, err
	}

	return order, nil
}

func (s Server) GetOrders(ctx context.Context, req *order_api.OrdersReq) (*order_api.OrdersRes, error) {
	log, ctx := tlog.New(ctx)
	orders, err := s.OrderService.GetOrders(ctx, req)
	if err != nil {
		log.Error("failed to get orders", "error", err)
		return nil, err
	}

	return &order_api.OrdersRes{
		Orders: orders,
	}, nil
}
