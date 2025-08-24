package demand_position_store

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_store"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/api/iterator"
)

const Collection = "demandPositions"

type Store struct {
	FirestoreClient *firestore.Client
}

type StorePosition struct {
	demand_api.Position
}

func (s Store) GetPositions(ctx context.Context, demand *demand_api.Demand) ([]*demand_api.Position, error) {
	var positions []*demand_api.Position

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
		var prod demand_api.Position
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		positions = append(positions, &prod)
	}

	return positions, nil
}

func (s Store) CreateOrUpdatePosition(ctx context.Context, demand *demand_api.Demand, position *demand_api.Position) (*demand_api.Position, error) {
	dr := prepToCreateOrUpdatePosition(s.FirestoreClient, demand.ID, position)
	if _, err := dr.Set(ctx, position); err != nil {
		return nil, err
	}

	return position, nil
}

func (s Store) CreateOrUpdatePositionWithTx(tx *firestore.Transaction, demandId string, position *demand_api.Position) (*demand_api.Position, error) {
	dr := prepToCreateOrUpdatePosition(s.FirestoreClient, demandId, position)
	if err := tx.Set(dr, position); err != nil {
		return nil, err
	}

	return position, nil
}

func prepToCreateOrUpdatePosition(firestoreClient *firestore.Client, demandId string, position *demand_api.Position) *firestore.DocumentRef {
	if len(position.ID) == 0 {
		position.ID = firestoreClient.Collection(Collection).NewDoc().ID
	}

	dr := firestoreClient.Collection(demand_store.Collection).Doc(demandId).Collection(Collection).Doc(position.ID)
	return dr
}

func (s Store) DeletePosition(ctx context.Context, demand *demand_api.Demand, id string) error {
	if _, err := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
