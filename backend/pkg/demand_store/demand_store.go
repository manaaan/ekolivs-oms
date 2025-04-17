package demand_store

import (
	"context"
	"errors"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/demand/api"
)

const collection = "demands"

type Store struct {
	FirestoreClient *firestore.Client
}

type StorePosition struct {
	api.Position
}

type StoreDemand struct {
	api.Demand
}

func (s Store) GetDemands(ctx context.Context) ([]*StoreDemand, error) {
	var demands []*StoreDemand
	iter := s.FirestoreClient.Collection(collection).OrderBy("ID", firestore.Asc).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		var prod StoreDemand
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		demands = append(demands, &prod)
	}
	return demands, nil
}

func (s Store) CreateDemand(ctx context.Context, demand *StoreDemand) (*StoreDemand, error) {
	if _, err := s.FirestoreClient.Collection(collection).Doc(demand.ID).Set(ctx, demand); err != nil {
		return nil, err
	}
	return demand, nil
}

func (s Store) DeleteDemand(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(collection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}
