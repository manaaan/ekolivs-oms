package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/sync-products/internal/product"
)

func main() {
	slog.Info("Received request to sync_products")
	env.LoadEnv()
	firestoreClient := gcp.InitFirestore()

	productService := product.Init(firestoreClient)
	err := productService.SyncProducts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Completed sync_products")
}
