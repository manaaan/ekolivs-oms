package demand_store

import (
	"context"
	"errors"
	"time"

	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"

	"cloud.google.com/go/firestore"
)

const Collection = "demands"

type Store struct {
	FirestoreClient *firestore.Client
}

func (s Store) GetDemands(ctx context.Context, req *demand_api.DemandsReq) ([]*demand_api.Demand, error) {
	var demands []*demand_api.Demand
	query := s.FirestoreClient.Collection(Collection).Query
	if req != nil && req.Member != nil {
		query = query.Where("Member", "==", *req.Member)
	}
	// `Name` is uppercase as in firestore, as we can't define the firestore structure in our .proto specs
	iter := query.OrderBy("CreationDate", firestore.Desc).Documents(ctx)
	defer iter.Stop()
	for {
		dsnap, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		var demand demand_api.Demand
		if err := dsnap.DataTo(&demand); err != nil {
			return nil, err
		}
		demands = append(demands, &demand)
	}

	return demands, nil
}

func (s Store) CreateOrUpdateDemand(ctx context.Context, data *demand_api.Demand) (*demand_api.Demand, error) {
	dr, copyData := prepToCreateOrUpdateDemand(s.FirestoreClient, data)
	if _, err := dr.Set(ctx, copyData); err != nil {
		return nil, err
	}
	return data, nil
}

func (s Store) CreateOrUpdateDemandWithTx(tx *firestore.Transaction, data *demand_api.Demand) (*demand_api.Demand, error) {
	dr, copyData := prepToCreateOrUpdateDemand(s.FirestoreClient, data)
	if err := tx.Set(dr, copyData); err != nil {
		return nil, err
	}

	return data, nil
}

func prepToCreateOrUpdateDemand(firestoreClient *firestore.Client, data *demand_api.Demand) (*firestore.DocumentRef, *demand_api.Demand) {
	if len(data.ID) == 0 {
		data.ID = firestoreClient.Collection(Collection).NewDoc().ID
		data.CreationDate = time.Now().Format(time.RFC3339)
	}

	// remove items from demand object (JSON) since they are created as a subcollection of demand
	// and should not be stored twice in firestore
	myCopy := proto.Clone(data).(*demand_api.Demand)
	myCopy.Items = nil

	dr := firestoreClient.Collection(Collection).Doc(myCopy.ID)
	return dr, myCopy
}

func (s Store) DeleteDemand(ctx context.Context, id string) error {
	if _, err := s.FirestoreClient.Collection(Collection).Doc(id).Delete(ctx); err != nil {
		return err
	}

	return nil
}
