package product

import (
	"fmt"

	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/pkg/zettle"
	"github.com/manaaan/ekolivs-oms/product/api"
)

type Service struct {
	zettleService *zettle.Service
	storeService  *Store
}

func New() (*Service, error) {
	zettleService, err := zettle.New(zettle.ServiceNewParams{
		ClientId: env.Required("ZETTLE_ORG_UUID"),
		ApiKey:   env.Required("ZETTLE_API_KEY"),
	})
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
	fmt.Println(zettleProducts)

	storeProducts, err := s.storeService.GetProducts()
	fmt.Println(storeProducts)

	// transform to api.Product slice
	return []*api.Product{}, nil
}
