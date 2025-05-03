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

func (s Store) GetDemands(ctx context.Context) ([]*api.Demand, error) {
	var demands []*api.Demand
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
		var demand api.Demand
		if err := dsnap.DataTo(&demand); err != nil {
			return nil, err
		}
		demands = append(demands, &demand)
	}
	return demands, nil
}

// TODO: create connected positions
// Mocked response
func (s Store) CreateDemand(ctx context.Context, create *api.CreateDemandReq) (*api.Demand, error) {
	return &api.Demand{
		ID:          "",
		Positions:   create.Positions,
		Status:      0,
		FulfilledAt: nil,
		CreatedAt:   "",
	}, nil
}

func (s Store) DeleteDemand(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(collection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}

// TODO: Should this stay here or in a demand_postition_store.go?
func (s Store) GetDemandPositions(ctx context.Context, productId string) ([]*api.Position, error) {
	var positions []*api.Position
	iter := s.FirestoreClient.Collection("positions").Where("ProductId", "==", productId).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		var prod api.Position
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		positions = append(positions, &prod)
	}
	return positions, nil
}

func (s Store) CreateDemandPosition(ctx context.Context, position *api.Position) (*api.Position, error) {
	if _, _, err := s.FirestoreClient.Collection("positions").Add(ctx, position); err != nil {
		return nil, err
	}

	return position, nil
}
