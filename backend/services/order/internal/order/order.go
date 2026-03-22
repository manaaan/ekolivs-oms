package order

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/orderstore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/specs/order_api"
)

type Service struct {
	Store *orderstore.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		Store: &orderstore.Store{
			FirestoreClient: firestoreClient,
		},
	}
}

func (s Service) GetOrders(ctx context.Context, req *order_api.OrdersReq) ([]*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	orders, err := s.Store.GetOrders(ctx, req)
	if err != nil {
		log.Error("failed to get orders from order store", "error", err)
		return nil, err
	}

	return orders, nil
}

func (s Service) GetOrderByID(ctx context.Context, id string) (*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	order, err := s.Store.GetOrder(ctx, id)
	if err != nil {
		log.Error("failed to get order from order store", "error", err, "orderID", id)
		return nil, err
	}

	return order, nil
}

func (s Service) CreateOrder(ctx context.Context, create *order_api.Order) (*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	order, err := s.Store.CreateOrUpdateOrder(ctx, create)
	if err != nil {
		log.Error("failed to create order", "error", err)
		return nil, err
	}

	return order, nil
}
