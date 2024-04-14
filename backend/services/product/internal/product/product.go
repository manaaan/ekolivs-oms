package product

import (
	"github.com/manaaan/ekolivs-oms/pkg/zettle"
	"github.com/manaaan/ekolivs-oms/product/api"
)

type Service struct {
	zettleService *zettle.Service
	storeService  *Store
}

func New() (*Service, error) {
	zettleService, err := zettle.New()
	if err != nil {
		return nil, err
	}

	return &Service{
		zettleService: zettleService,
		storeService:  &Store{},
	}, nil
}

func (s Service) GetProducts() ([]*api.Product, error) {
	zettleProducts, err := s.zettleService.GetProducts()
	if err != nil {
		return nil, err
	}

	storeProducts, err := s.storeService.GetProducts()
	// transform to api.Product slice
	return []*api.Product{}, nil
}
