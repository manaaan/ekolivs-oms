package demand

import "github.com/manaaan/ekolivs-oms/demand/api"

type Service struct{}

func New() (*Service, error) {

	return &Service{}, nil
}

func (s Service) CreateDemand(input *api.CreateDemandReq) (api.DemandRes, error) {
	return api.DemandRes{}, nil
}
