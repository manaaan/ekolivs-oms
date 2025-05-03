package demand

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/pkg/demand_store"
	"log/slog"
)

type Service struct {
	storeService *demand_store.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		storeService: &demand_store.Store{
			FirestoreClient: firestoreClient,
		},
	}
}

// TODO: get positions
func (s Service) GetDemands(ctx context.Context) ([]*api.Demand, error) {
	demands, err := s.storeService.GetDemands(ctx)
	if err != nil {
		slog.Error("failed to get demands from demand store", "error", err)
		return nil, err
	}

	return demands, nil
}

// TODO: create positions
func (s Service) CreateDemand(ctx context.Context, create *api.CreateDemandReq) (*api.Demand, error) {
	demand, err := s.storeService.CreateDemand(ctx, create)
	if err != nil {
		slog.Error("failed to create demand", "error", err)
		return nil, err
	}

	return demand, nil
}

// TODO: delete positions
func (s Service) DeleteDemand(ctx context.Context, id string) error {
	err := s.storeService.DeleteDemand(ctx, id)
	if err != nil {
		slog.Error("failed to delete demand", "error", err)
		return err
	}
	return nil
}
