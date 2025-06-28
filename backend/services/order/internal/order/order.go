package order

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/order_store"
	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/services/order/order_api"
)

type Service struct {
	Store *order_store.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		Store: &order_store.Store{
			FirestoreClient: firestoreClient,
		},
	}
}

func (s Service) GetOrders(ctx context.Context) ([]*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	orders, err := s.Store.GetOrders(ctx)
	if err != nil {
		log.Error("failed to get orders from order store", "error", err)
		return nil, err
	}

	return orders, nil
}

func (s Service) GetOrderByID(ctx context.Context, ID string) (*order_api.Order, error) {
	log, ctx := tlog.New(ctx)
	order, err := s.Store.GetOrder(ctx, ID)
	if err != nil {
		log.Error("failed to get order from order store", "error", err, "orderID", ID)
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
