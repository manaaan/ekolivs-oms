package server

import (
	"context"
	"errors"

	"github.com/manaaan/ekolivs-oms/backend/pkg/errkit"
	"github.com/manaaan/ekolivs-oms/backend/pkg/tlog"
	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/demand"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	demand_api.UnimplementedDemandServiceServer
	DemandService *demand.Service
}

func (s Server) GetDemands(ctx context.Context, req *demand_api.DemandsReq) (*demand_api.DemandsRes, error) {
	log, ctx := tlog.New(ctx)
	demands, err := s.DemandService.GetDemands(ctx, req)
	if err != nil {
		log.Error("Unable to get demands", "error", err)
		return nil, err
	}

	return &demand_api.DemandsRes{Demands: demands}, nil
}

func (s Server) GetDemand(ctx context.Context, idReq *demand_api.IdReq) (*demand_api.Demand, error) {
	if idReq == nil || len(idReq.Id) < 1 {
		return nil, errkit.BuildGRPCStatusErr(ctx, &errkit.ErrBadRequest{Err: errors.New("need ID to get demand")})
	}
	log, ctx := tlog.New(ctx)
	demandClient, err := s.DemandService.GetDemand(ctx, idReq)
	if err != nil {
		log.Error("unable to get demand", "error", err)
	}
	return demandClient, nil
}

func (s Server) CreateDemand(ctx context.Context, req *demand_api.CreateDemandReq) (*demand_api.Demand, error) {
	log, ctx := tlog.New(ctx)
	if req == nil || req.Items == nil || len(req.Items) == 0 {
		return nil, errkit.BuildGRPCStatusErr(ctx, &errkit.ErrBadRequest{Err: errors.New("missing required input: items")})
	} else if len(req.Member) == 0 {
		log.Info(req.Member)
		return nil, errkit.BuildGRPCStatusErr(ctx, &errkit.ErrBadRequest{Err: errors.New("missing required input: member")})
	}

	data := &demand_api.Demand{
		Items:  req.Items,
		Member: req.Member,
	}

	d, err := s.DemandService.CreateOrUpdateDemand(ctx, data)
	if err != nil {
		log.Error("failed to update demand", "error", err)
		return nil, err
	}

	return d, nil
}

func (s Server) DeleteDemand(ctx context.Context, id *demand_api.IdReq) (*emptypb.Empty, error) {
	log, ctx := tlog.New(ctx)
	err := s.DemandService.DeleteDemand(ctx, id.Id)
	if err != nil {
		log.Error("failed to delete demand", "error", err)
		return nil, err
	}
	return nil, nil
}
