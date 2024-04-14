package zettle

type Service struct{}

type product struct{}

func New() (*Service, error) {
	return &Service{}, nil
}

func (Service) GetProducts() (product, error) {
	return product{}, nil
}
