package product

type Store struct{}

type storeProduct struct{}

func (s Store) GetProducts() ([]storeProduct, error) {
	return []storeProduct{}, nil
}
