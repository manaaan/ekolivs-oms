package demand_store

import (
	"context"
	"errors"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/demand/api"
)

const Collection = "demands"

type Store struct {
	FirestoreClient *firestore.Client
}

func (s Store) GetDemands(ctx context.Context) ([]*api.Demand, error) {
	var demands []*api.Demand

	iter := s.FirestoreClient.Collection(Collection).Documents(ctx)
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

func (s Store) CreateOrUpdateDemand(ctx context.Context, data *api.Demand) (*api.Demand, error) {
	if _, err := s.FirestoreClient.Collection(Collection).Doc(data.ID).Set(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s Store) CreateOrUpdateDemandWithTx(tx *firestore.Transaction, data *api.Demand) (*api.Demand, error) {
	if len(data.ID) == 0 {
		data.ID = s.FirestoreClient.Collection(Collection).NewDoc().ID
		data.CreatedAt = time.Now().Format(time.RFC3339)
	}

	// remove positions since they are created as a subcollection of demand
	myCopy := proto.Clone(data).(*api.Demand)
	myCopy.Positions = nil

	dr := s.FirestoreClient.Collection(Collection).Doc(myCopy.ID)

	if err := tx.Set(dr, myCopy); err != nil {
		return nil, err
	}

	return data, nil
}

func (s Store) DeleteDemand(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
