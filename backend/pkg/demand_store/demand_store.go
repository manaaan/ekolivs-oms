package demand_store

import (
	"context"
	"errors"
	"time"

	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"

	"cloud.google.com/go/firestore"
)

const Collection = "demands"

type Store struct {
	FirestoreClient *firestore.Client
	ItemsStore      *ItemStore
}

func New(firestoreClient *firestore.Client) *Store {
	store := &Store{
		FirestoreClient: firestoreClient,
	}
	store.ItemsStore = &ItemStore{
		FirestoreClient: firestoreClient,
	}
	return store
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
		demand, err := mapFirestoreSnapToProtoDemand(dsnap)
		if err != nil {
			return nil, err
		}
		demands = append(demands, demand)
	}

	if req.ExtendItems != nil && *req.ExtendItems {
		g, fetchCtx := errgroup.WithContext(ctx)

		for _, demand := range demands {
			g.Go(func() error {
				items, err := s.getItems(fetchCtx, demand)
				if err != nil {
					return err
				}
				// TODO: is this safe? Need to verify in test
				demand.Items = items
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			return nil, err
		}

	}

	return demands, nil
}

func (s Store) GetDemand(ctx context.Context, id string) (*demand_api.Demand, error) {
	docRef := s.FirestoreClient.Collection(Collection).Doc(id)
	dsnap, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}
	demand, err := mapFirestoreSnapToProtoDemand(dsnap)
	if err != nil {
		return nil, err
	}
	items, err := s.getItems(ctx, demand)
	if err != nil {
		return nil, err
	}
	demand.Items = items
	return demand, nil
}

func mapFirestoreSnapToProtoDemand(dsnap *firestore.DocumentSnapshot) (*demand_api.Demand, error) {
	var demand demand_api.Demand
	if err := dsnap.DataTo(&demand); err != nil {
		return nil, err
	}
	return &demand, nil
}

func (s Store) getItems(ctx context.Context, demand *demand_api.Demand) ([]*demand_api.Item, error) {
	log, ctx := tlog.New(ctx)
	items, err := s.ItemsStore.GetItems(ctx, demand)
	if err != nil {
		log.Error("failed to get items from demand item store", "error", err)
		return nil, err
	}

	return items, nil
}

func (s Store) CreateOrUpdateDemand(ctx context.Context, data *demand_api.Demand) (*demand_api.Demand, error) {
	log, ctx := tlog.New(ctx)
	err := s.FirestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		dr, copyData := prepToCreateOrUpdateDemand(s.FirestoreClient, data)
		if err := tx.Set(dr, copyData); err != nil {
			log.Error("failed to create or update demand", "error", err)
			return err
		}

		for _, item := range data.Items {
			_, err := s.ItemsStore.CreateOrUpdateDemandItemWithTx(tx, data.ID, item)
			if err != nil {
				log.Error("failed to create or update item", "error", err)
				return err
			}
		}

		return nil
	})

	if err != nil {
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
