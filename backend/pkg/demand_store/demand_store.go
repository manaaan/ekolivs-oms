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

// TODO: join positions with demands
func (s Store) GetDemands(ctx context.Context) ([]*api.Demand, error) {
	var demands []*api.Demand
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
		var prod api.Demand
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		demands = append(demands, &prod)
	}
	return demands, nil
}

// TODO: create connected positions
func (s Store) CreateDemand(ctx context.Context, create *api.CreateDemandReq) (*api.Demand, error) {
	//if _, err := s.FirestoreClient.Collection(collection)
	//	return nil, err
	//}
	//return demand, nil
	return &api.Demand{
		ID:          "",
		Positions:   create.Positions,
		Status:      0,
		FulfilledAt: nil,
		CreatedAt:   "",
	}, nil
}

// TODO: delete connected positions
func (s Store) DeleteDemand(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(collection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}
