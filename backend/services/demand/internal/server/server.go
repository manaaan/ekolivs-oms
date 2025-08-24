package server

import (
	"context"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/demand"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	demand_api.UnimplementedDemandServiceServer
	DemandService *demand.Service
}

func (s Server) GetDemands(ctx context.Context, req *demand_api.DemandsReq) (*demand_api.DemandsRes, error) {
	demands, err := s.DemandService.GetDemands(ctx, req)
	if err != nil {
		slog.Error("Unable to get demands")
		return nil, err
	}

	return &demand_api.DemandsRes{Demands: demands}, nil
}

func (s Server) CreateDemand(ctx context.Context, req *demand_api.CreateDemandReq) (*demand_api.Demand, error) {
	if req == nil || req.Positions == nil || len(req.Positions) == 0 {
		return nil, nil
	}

	data := &demand_api.Demand{
		Positions: req.Positions,
	}

	d, err := s.DemandService.CreateOrUpdateDemand(ctx, data)
	if err != nil {
		slog.Error("failed to update demand", "error", err)
		return nil, err
	}

	return d, nil
}

func (s Server) DeleteDemand(ctx context.Context, id *demand_api.IdReq) (*emptypb.Empty, error) {
	err := s.DemandService.DeleteDemand(ctx, id.Id)
	if err != nil {
		slog.Error("failed to delete demand", "error", err)
		return nil, err
	}
	return nil, nil
}
