package product

import (
	"context"
	"log/slog"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/pkg/product_store"
	"github.com/manaaan/ekolivs-oms/pkg/zettle"
	"github.com/manaaan/ekolivs-oms/product/api"
	"golang.org/x/sync/errgroup"
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
	storeProducts, err := s.storeService.GetProducts(ctx)
	if err != nil {
		slog.Error("failed to get products from product store", "error", err)
		return nil, err
	}

	products := []*api.Product{}
	for _, storeProduct := range storeProducts {
		products = append(products, &api.Product{
			ID:            storeProduct.ID,
			Name:          storeProduct.Name,
			Sku:           storeProduct.Sku,
			Barcode:       storeProduct.Barcode,
			Price:         storeProduct.Price,
			CostPrice:     storeProduct.CostPrice,
			ImageUrl:      storeProduct.ImageUrl,
			VatPercentage: storeProduct.VatPercentage,
			Status:        storeProduct.Status,
			UnitType:      storeProduct.UnitType,
			CreatedAt:     storeProduct.CreatedAt,
			UpdatedAt:     storeProduct.UpdatedAt,
		})
	}

	return products, nil
}

func (s Service) GetProductByID(ctx context.Context, id string) (*api.Product, error) {
	storeProduct, err := s.storeService.GetProduct(ctx, id)
	if err != nil {
		slog.Error("failed to get products from product store", "error", err)
		return nil, err
	}

	product := mapStoreToAPIProduct(storeProduct)

	return product, nil
}

func (s Service) UpdateProduct(ctx context.Context, product *api.Product) (*api.Product, error) {
	g, fetchCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := s.storeService.CreateOrUpdateProduct(fetchCtx, mapAPIToStoreProduct(product))
		if err != nil {
			return err
		}
		return nil
	})

	// TODO: Check if product requires change to be executed in Zettle as well
	g.Go(func() error {
		// TODO: We need to map the variants back into Zettle structure
		err := s.zettleService.UpdateProduct(
			zettle.UpdateProductParamsExt{ProductUuid: product.ID},
			zettle.FullProductUpdateRequest{
				Name: product.Name,
			},
		)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		slog.Error("failed to update product", "error", err)
		return nil, err
	}

	return product, nil
}

func mapStoreToAPIProduct(storeProduct *product_store.StoreProduct) *api.Product {
	if storeProduct == nil {
		return nil
	}
	return &api.Product{
		ID:            storeProduct.ID,
		Name:          storeProduct.Name,
		Sku:           storeProduct.Sku,
		Barcode:       storeProduct.Barcode,
		Price:         storeProduct.Price,
		CostPrice:     storeProduct.CostPrice,
		ImageUrl:      storeProduct.ImageUrl,
		VatPercentage: storeProduct.VatPercentage,
		Status:        storeProduct.Status,
		UnitType:      storeProduct.UnitType,
		CreatedAt:     storeProduct.CreatedAt,
		UpdatedAt:     storeProduct.UpdatedAt,
	}
}

func mapAPIToStoreProduct(product *api.Product) *product_store.StoreProduct {
	if product == nil {
		return nil
	}
	return &product_store.StoreProduct{
		Product: api.Product{
			ID:            product.ID,
			Name:          product.Name,
			Sku:           product.Sku,
			Barcode:       product.Barcode,
			Price:         product.Price,
			CostPrice:     product.CostPrice,
			ImageUrl:      product.ImageUrl,
			VatPercentage: product.VatPercentage,
			Status:        product.Status,
			UnitType:      product.UnitType,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
		},
		// TODO: Should they not just be part of the api.Product proto specs?
		Supplier: product_store.GetSupplierForProduct(product.Name),
		Source:   "zettle", // TODO: hardcoded Zettle right now
	}
}
