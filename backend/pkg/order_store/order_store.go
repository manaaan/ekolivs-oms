package order_store

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/order/order_api"
	"google.golang.org/api/iterator"
)

const collection = "orders"

type Store struct {
	FirestoreClient *firestore.Client
}

// TODO: add filters to query, which require firestore indexes
func (s Store) GetOrders(ctx context.Context) ([]*order_api.Order, error) {
	orders := []*order_api.Order{}
	// `Name` is uppercase as in firestore, as we can't define the firestore structure in our .proto specs
	iter := s.FirestoreClient.Collection(collection).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		order, err := mapFirestoreSnapToProtoOrder(dsnap)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (s Store) GetOrder(ctx context.Context, id string) (*order_api.Order, error) {
	dsnap, err := s.FirestoreClient.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	return mapFirestoreSnapToProtoOrder(dsnap)
}

// Overwrites the order document in firestore completely
func (s Store) CreateOrUpdateOrder(ctx context.Context, data *order_api.Order) (*order_api.Order, error) {
	if len(data.ID) == 0 {
		data.ID = s.FirestoreClient.Collection(collection).NewDoc().ID
		data.CreationDate = time.Now().Format(time.RFC3339)
	}
	if _, err := s.FirestoreClient.Collection(collection).Doc(data.ID).Set(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Remove order document in firestore
func (s Store) DeleteOrder(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(collection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}

func mapFirestoreSnapToProtoOrder(dsnap *firestore.DocumentSnapshot) (*order_api.Order, error) {
	var order order_api.Order
	if err := dsnap.DataTo(&order); err != nil {
		return nil, err
	}
	return &order, nil
}
