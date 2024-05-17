package server

import (
	"context"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/demand/internal/demand"
)

type Server struct {
	api.UnimplementedDemandServiceServer
	DemandService *demand.Service
}

func (s Server) CreateDemand(ctx context.Context, req *api.CreateDemandReq) (*api.DemandRes, error) {
	_, err := s.DemandService.CreateDemand(req)
	if err != nil {
		slog.Error("Unable to create demand")
		return nil, err
	}

	return &api.DemandRes{}, nil
}
