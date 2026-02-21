package demandstore

import (
	"context"
	"errors"
	"time"

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

func (s ItemStore) CreateOrUpdateDemandItem(ctx context.Context, demand *demand_api.Demand, item *demand_api.Item, position uint32) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demand.ID, item, position)
	if _, err := dr.Set(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s ItemStore) CreateOrUpdateDemandItemWithTx(tx *firestore.Transaction, demandID string, item *demand_api.Item, position uint32) (*demand_api.Item, error) {
	dr := prepToCreateOrUpdateDemandItem(s.FirestoreClient, demandID, item, position)
	if err := tx.Set(dr, item); err != nil {
		return nil, err
	}

	return item, nil
}

func prepToCreateOrUpdateDemandItem(firestoreClient *firestore.Client, demandID string, item *demand_api.Item, position uint32) *firestore.DocumentRef {
	if item.ID == "" {
		item.ID = firestoreClient.Collection(Collection).Doc(demandID).Collection(ItemCollection).NewDoc().ID
		item.CreationDate = time.Now().Format(time.RFC3339)
		item.DemandID = demandID
		item.Status = demand_api.Status_RECEIVED
		item.Position = position
	}

	dr := firestoreClient.Collection(Collection).Doc(demandID).Collection(ItemCollection).Doc(item.ID)
	return dr
}

func (s ItemStore) DeleteDemandItem(ctx context.Context, demandID, itemID string) error {
	if _, err := s.FirestoreClient.Collection(Collection).Doc(demandID).Collection(ItemCollection).Doc(itemID).Delete(ctx); err != nil {
		return err
	}
	return nil
}
