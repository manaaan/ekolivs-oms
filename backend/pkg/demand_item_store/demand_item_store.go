package demand_item_store

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_store"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/api/iterator"
)

const Collection = "demandItems"

type Store struct {
	FirestoreClient *firestore.Client
}

type StoreDemandItem struct {
	demand_api.Item
}

func (s Store) GetItems(ctx context.Context, demand *demand_api.Demand) ([]*demand_api.Item, error) {
	var items []*demand_api.Item

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
		var prod demand_api.Item
		if err := dsnap.DataTo(&prod); err != nil {
			return nil, err
		}
		items = append(items, &prod)
	}

	return items, nil
}

func (s Store) CreateOrUpdateDemandItem(ctx context.Context, demand *demand_api.Demand, item *demand_api.Item) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demand.ID, item)
	if _, err := dr.Set(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s Store) CreateOrUpdateDemandItemWithTx(tx *firestore.Transaction, demandId string, item *demand_api.Item) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demandId, item)
	if err := tx.Set(dr, item); err != nil {
		return nil, err
	}

	return item, nil
}

func prepToCreateOrUpdateDemandItem(firestoreClient *firestore.Client, demandId string, item *demand_api.Item) *firestore.DocumentRef {
	if len(item.ID) == 0 {
		item.ID = firestoreClient.Collection(Collection).NewDoc().ID
	}

	dr := firestoreClient.Collection(demand_store.Collection).Doc(demandId).Collection(Collection).Doc(item.ID)
	return dr
}

func (s Store) DeleteDemandItem(ctx context.Context, demand *demand_api.Demand, id string) error {
	if _, err := s.FirestoreClient.Collection(demand_store.Collection).Doc(demand.ID).Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
