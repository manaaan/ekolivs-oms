package zettle

type Service struct{}

func New() (*Service, error) {
	return &Service{}, nil
}

func (Service) GetProducts() (ProductResponse, error) {
	return ProductResponse{}, nil
}
