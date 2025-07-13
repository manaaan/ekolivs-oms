package demand_position_store

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/pkg/demand_store"
	"google.golang.org/api/iterator"
)

const Collection = "demandPositions"

type Store struct {
	FirestoreClient *firestore.Client
}

type StorePosition struct {
	api.Position
}

func (s Store) GetPositions(ctx context.Context, demand *api.Demand) ([]*api.Position, error) {
	var positions []*api.Position

	iter := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Documents(ctx)
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

func (s Store) CreateOrUpdatePosition(ctx context.Context, demand *api.Demand, position *api.Position) (*api.Position, error) {
	if _, _, err := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Add(ctx, position); err != nil {
		return nil, err
	}

	return position, nil
}

func (s Store) CreateOrUpdatePositionWithTx(tx *firestore.Transaction, demandId string, position *api.Position) (*api.Position, error) {
	if len(position.ID) == 0 {
		position.ID = s.FirestoreClient.Collection(Collection).NewDoc().ID
	}

	dr := s.FirestoreClient.Collection(demand_store.Collection).Doc(demandId).Collection(Collection).Doc(position.ID)

	if err := tx.Set(dr, position); err != nil {
		return nil, err
	}

	return position, nil
}

func (s Store) DeletePosition(ctx context.Context, demand *api.Demand, id string) error {
	if _, err := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
