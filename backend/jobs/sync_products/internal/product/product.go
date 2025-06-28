package product

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/pkg/product_store"
	"github.com/manaaan/ekolivs-oms/backend/pkg/zettle"
	"github.com/manaaan/ekolivs-oms/backend/services/product/api"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	zettleService *zettle.Service
	storeService  *product_store.Store
}

// Initialization function for call in `main` pkg that panics on failure
func Init(firestoreClient *firestore.Client) *Service {
	productService, err := New(firestoreClient)
	if err != nil {
		log.Fatalf("Unable to initialize product service")
		return nil
	}
	return productService
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

// Sync products from Zettle to our DB
func (s Service) SyncProducts(ctx context.Context) error {
	zettleProducts, err := s.zettleService.GetProducts()
	if err != nil {
		return err
	}

	g, fetchCtx := errgroup.WithContext(ctx)

	for _, zettleProduct := range zettleProducts {
		for i, variant := range zettleProduct.Variants {
			if i > 1 {
				continue
			}
			g.Go(func() error {
				var imageUrl *string
				if zettleProduct.Presentation != nil {
					imageUrl = zettleProduct.Presentation.ImageUrl
				}

				product := &product_store.StoreProduct{
					Product: api.Product{
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
					},
					Supplier: product_store.GetSupplierForProduct(zettleProduct.Name),
					Source:   "zettle",
				}

				_, err := s.storeService.CreateOrUpdateProduct(fetchCtx, product)
				if err != nil {
					return err
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		slog.Error("Failed to sync products from zettle", "error", err)
		return err
	}
	slog.Info("Synched products from zettle", "noOfProducts", len(zettleProducts))

	return nil
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
