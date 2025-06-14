package server

import (
	"context"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/demand/internal/demand"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	api.UnimplementedDemandServiceServer
	DemandService *demand.Service
}

func (s Server) GetDemands(ctx context.Context, _ *emptypb.Empty) (*api.DemandsRes, error) {
	demands, err := s.DemandService.GetDemands(ctx)
	if err != nil {
		slog.Error("Unable to get demands")
		return nil, err
	}

	return &api.DemandsRes{Demands: demands}, nil
}

func (s Server) CreateDemand(ctx context.Context, data *api.Demand) (*api.Demand, error) {
	d, err := s.DemandService.CreateOrUpdateDemand(ctx, data)
	if err != nil {
		slog.Error("failed to update product", "error", err)
		return nil, err
	}

	return d, nil
}

func (s Server) DeleteDemand(ctx context.Context, id *api.IdReq) (*emptypb.Empty, error) {
	err := s.DemandService.DeleteDemand(ctx, id.Id)
	if err != nil {
		slog.Error("failed to delete demand", "error", err)
		return nil, err
	}
	return nil, nil
}
