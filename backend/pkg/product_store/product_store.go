package product_store

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/product/api"
	"google.golang.org/api/iterator"
)

const (
	collection = "products"
)

type Store struct {
	FirestoreClient *firestore.Client
}

type StoreProduct struct {
	api.Product
	Supplier string `json:"supplier" firestore:"supplier,omitempty"`
	Source   string `json:"source" firestore:"source,omitempty"`
}

// TODO: add filters to query, which require firestore indexes
func (s Store) GetProducts(ctx context.Context) ([]*StoreProduct, error) {
	products := []*StoreProduct{}
	// TODO: Further sorting by price? Would require firestore indexes
	iter := s.FirestoreClient.Collection(collection).OrderBy("name", firestore.Asc).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var prod StoreProduct
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}
	return products, nil
}

func (s Store) GetProduct(ctx context.Context, id string) (*StoreProduct, error) {
	dsnap, err := s.FirestoreClient.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var prod StoreProduct
	if err := dsnap.DataTo(&prod); err != nil {
		return nil, err
	}
	return &prod, nil
}

// Overwrites the product document in firestore completely
func (s Store) CreateOrUpdateProduct(ctx context.Context, data *StoreProduct) (*StoreProduct, error) {
	if _, err := s.FirestoreClient.Collection(collection).Doc(data.ID).Set(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Remove product document in firestore
func (s Store) DeleteProduct(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(collection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}
