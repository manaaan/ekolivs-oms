package demand

import (
	"context"
	"log/slog"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_position_store"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_store"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
)

type Service struct {
	firestoreClient *firestore.Client
	demandStore     *demand_store.Store
	positionStore   *demand_position_store.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		firestoreClient: firestoreClient,
		demandStore: &demand_store.Store{
			FirestoreClient: firestoreClient,
		},
		positionStore: &demand_position_store.Store{
			FirestoreClient: firestoreClient,
		},
	}
}

func (s Service) GetDemands(ctx context.Context, req *demand_api.DemandsReq) ([]*demand_api.Demand, error) {
	// TODO: Introduce concurrency to improve response times
	demands, err := s.demandStore.GetDemands(ctx, req)
	if err != nil {
		slog.Error("failed to get demands from demand store", "error", err)
		return nil, err
	}

	for _, demand := range demands {
		positions, err := s.positionStore.GetPositions(ctx, demand)
		if err != nil {
			slog.Error("failed to get positions from position store", "error", err)
			return nil, err
		}

		demand.Positions = positions
	}

	return demands, nil
}

func (s Service) CreateOrUpdateDemand(ctx context.Context, data *demand_api.Demand) (*demand_api.Demand, error) {
	err := s.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		demand, err := s.demandStore.CreateOrUpdateDemandWithTx(tx, data)
		if err != nil {
			slog.Error("failed to create or update demand", "error", err)
			return err
		}

		for _, position := range data.Positions {
			_, err := s.positionStore.CreateOrUpdatePositionWithTx(tx, demand.ID, position)
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

// TODO: Add DeleteDemandPosition, add UpdateDemandPosition
// NOTE: Not exposed yet, as we don't fully delete the Demand with Positions
func (s Service) DeleteDemand(ctx context.Context, id string) error {
	return nil
	// err := s.demandStore.DeleteDemand(ctx, id)
	// if err != nil {
	// 	slog.Error("failed to delete demand", "error", err)
	// 	return err
	// }
	// return nil
}
