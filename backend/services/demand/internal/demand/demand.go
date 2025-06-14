package demand

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/pkg/demand_postition_store"
	"github.com/manaaan/ekolivs-oms/pkg/demand_store"
	"log/slog"
)

type Service struct {
	firestoreClient *firestore.Client
	demandStore     *demand_store.Store
	positionStore   *demand_postition_store.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		firestoreClient: firestoreClient,
		demandStore: &demand_store.Store{
			FirestoreClient: firestoreClient,
		},
		positionStore: &demand_postition_store.Store{
			FirestoreClient: firestoreClient,
		},
	}
}

func (s Service) GetDemands(ctx context.Context) ([]*api.Demand, error) {
	demands, err := s.demandStore.GetDemands(ctx)
	if err != nil {
		slog.Error("failed to get demands from demand store", "error", err)
		return nil, err
	}

	var positions []*api.Position
	for _, demand := range demands {
		positions, err = s.positionStore.GetPositions(ctx, demand)
		if err != nil {
			slog.Error("failed to get positions from position store", "error", err)
			return nil, err
		}

		demand.Positions = positions
	}

	return demands, nil
}

func (s Service) CreateOrUpdateDemand(ctx context.Context, data *api.Demand) (*api.Demand, error) {
	demandRef := s.firestoreClient.Collection("demands").Doc(data.ID)
	positionsCollection := demandRef.Collection("positions")
	positionRefs := make(map[string]*firestore.DocumentRef)
	for _, position := range data.Positions {
		positionRefs[position.ID] = positionsCollection.Doc(position.ID)
	}

	err := s.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		_, err := s.demandStore.CreateOrUpdateDemandWithTx(tx, demandRef, data)
		if err != nil {
			slog.Error("failed to create or update demand", "error", err)
			return err
		}

		for _, position := range data.Positions {
			_, err := s.positionStore.CreateOrUpdatePositionWithTx(tx, positionRefs[position.ID], position)
			if err != nil {
				slog.Error("failed to create or update position", "error", err)
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

// TODO: delete positions
func (s Service) DeleteDemand(ctx context.Context, id string) error {
	err := s.demandStore.DeleteDemand(ctx, id)
	if err != nil {
		slog.Error("failed to delete demand", "error", err)
		return err
	}
	return nil
}
