package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/backend/jobs/sync_products/internal/product"
	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
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
