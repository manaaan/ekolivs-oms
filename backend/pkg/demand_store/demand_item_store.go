package demand_store

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/api/iterator"
)

const ItemCollection = "demandItems"

type ItemStore struct {
	FirestoreClient *firestore.Client
}

func (s ItemStore) GetItems(ctx context.Context, demand *demand_api.Demand) ([]*demand_api.Item, error) {
	var items []*demand_api.Item

	iter := s.FirestoreClient.Collection(Collection).Doc(demand.ID).Collection(ItemCollection).Documents(ctx)
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

func (s ItemStore) CreateOrUpdateDemandItem(ctx context.Context, demand *demand_api.Demand, item *demand_api.Item) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demand.ID, item)
	if _, err := dr.Set(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s ItemStore) CreateOrUpdateDemandItemWithTx(tx *firestore.Transaction, demandId string, item *demand_api.Item) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demandId, item)
	if err := tx.Set(dr, item); err != nil {
		return nil, err
	}

	return item, nil
}

func prepToCreateOrUpdateDemandItem(firestoreClient *firestore.Client, demandId string, item *demand_api.Item) *firestore.DocumentRef {
	if len(item.ID) == 0 {
		item.ID = firestoreClient.Collection(Collection).Doc(demandId).Collection(ItemCollection).NewDoc().ID
	}

	dr := firestoreClient.Collection(Collection).Doc(demandId).Collection(ItemCollection).Doc(item.ID)
	return dr
}

func (s ItemStore) DeleteDemandItem(ctx context.Context, demandId string, itemId string) error {
	if _, err := s.FirestoreClient.Collection(Collection).Doc(demandId).Collection(ItemCollection).Doc(itemId).Delete(ctx); err != nil {
		return err
	}
	return nil
}
