package demand

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/demand_store"
	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
)

type Service struct {
	demandStore *demand_store.Store
}

func New(firestoreClient *firestore.Client) *Service {
	return &Service{
		demandStore: demand_store.New(firestoreClient),
	}
}

func (s Service) GetDemands(ctx context.Context, req *demand_api.DemandsReq) ([]*demand_api.Demand, error) {
	log, ctx := tlog.New(ctx)
	demands, err := s.demandStore.GetDemands(ctx, req)
	if err != nil {
		log.Error("failed to get demands from demand store", "error", err)
		return nil, err
	}

	return demands, nil
}

func (s Service) GetDemand(ctx context.Context, idReq *demand_api.IdReq) (*demand_api.Demand, error) {
	log, ctx := tlog.New(ctx)
	demand, err := s.demandStore.GetDemand(ctx, idReq.Id)
	if err != nil {
		log.Error("failed to get demand from demand store", "error", err)
		return nil, err
	}

	return demand, nil
}

func (s Service) CreateOrUpdateDemand(ctx context.Context, data *demand_api.Demand) (*demand_api.Demand, error) {
	log, ctx := tlog.New(ctx)
	demand, err := s.demandStore.CreateOrUpdateDemand(ctx, data)
	if err != nil {
		log.Error("failed to create or update demand", "error", err)
		return nil, err
	}

	return demand, nil
}

// TODO: Add DeleteDemandItem, add UpdateDemandItem
// NOTE: Not exposed yet, as we don't fully delete the Demand with Items
func (s Service) DeleteDemand(ctx context.Context, id string) error {
	return nil
	// err := s.demandStore.DeleteDemand(ctx, id)
	// if err != nil {
	// 	slog.Error("failed to delete demand", "error", err)
	// 	return err
	// }
	// return nil
}
