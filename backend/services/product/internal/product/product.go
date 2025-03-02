package product

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/pkg/product_store"
	"github.com/manaaan/ekolivs-oms/pkg/zettle"
	"github.com/manaaan/ekolivs-oms/product/api"
)

type Service struct {
	zettleService *zettle.Service
	storeService  *product_store.Store
}

func New(firestoreClient *firestore.Client) (*Service, error) {
	zettleService, err := zettle.New(zettle.ServiceNewParams{
		ClientId: env.Required("ZETTLE_CLIENT_ID"),
		ApiKey:   env.Required("ZETTLE_API_KEY"),
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		zettleService: zettleService,
		storeService: &product_store.Store{
			FirestoreClient: firestoreClient,
		},
	}, nil
}

func (s Service) GetProducts(ctx context.Context) ([]*api.Product, error) {
	zettleProducts, err := s.zettleService.GetProducts()
	if err != nil {
		return nil, err
	}

	storeProducts, err := s.storeService.GetProducts(ctx)
	fmt.Println(storeProducts)

	products := []*api.Product{}
	for _, zettleProduct := range zettleProducts {
		for _, variant := range zettleProduct.Variants {
			var imageUrl *string
			if zettleProduct.Presentation != nil {
				imageUrl = zettleProduct.Presentation.ImageUrl
			}

			products = append(products, &api.Product{
				ID:            variant.Uuid.String(),
				Name:          buildProductName(zettleProduct, variant),
				Sku:           variant.Sku,
				Barcode:       variant.Barcode,
				Price:         convertToPrice(variant.Price),
				CostPrice:     convertToPrice(variant.CostPrice),
				ImageUrl:      imageUrl,
				VatPercentage: zettleProduct.VatPercentage,
				Status:        api.Status_ACTIVE,
				UnitType:      convertToUnitType(zettleProduct.UnitName),
				CreatedAt:     zettleProduct.Created,
				UpdatedAt:     zettleProduct.Updated,
			})
		}
	}

	return products, nil
}

func convertToPrice(zettlePrice *zettle.Price) *api.Price {
	if zettlePrice == nil {
		return nil
	}

	return &api.Price{
		Amount:     zettlePrice.Amount,
		CurrencyID: string(zettlePrice.CurrencyId),
	}
}

func convertToUnitType(unitName *string) api.UnitType {
	if unitName == nil {
		return api.UnitType_PIECES
	}

	switch strings.ToLower(*unitName) {
	case "st":
		return api.UnitType_PIECES
	case "g":
		return api.UnitType_GRAMS
	case "kg":
		return api.UnitType_KILOGRAMS
	case "ml":
		return api.UnitType_MILLILITER
	case "l":
		return api.UnitType_LITER
	default:
		return api.UnitType_PIECES
	}
}

func buildProductName(zettleProduct zettle.ProductResponse, variant zettle.VariantDTO) string {
	if variant.Name == nil {
		return zettleProduct.Name
	}
	return fmt.Sprintf("%s - %s", zettleProduct.Name, *variant.Name)
}
