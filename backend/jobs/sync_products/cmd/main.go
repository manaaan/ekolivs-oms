package main

import (
	"context"
	"flag"
	"log"
	"log/slog"

	"github.com/manaaan/ekolivs-oms/backend/jobs/sync_products/internal/product"
	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
)

var fileName = ".env"

func init() {
	flag.StringVar(&fileName, "f", ".env", "Requires the absolute path of the filename slash the fileName. Example: /absolute_path/fileName")
}

func main() {
	flag.Parse()
	env.LoadEnv(fileName)
	slog.Info("Received request to sync_products")
	firestoreClient := gcp.InitFirestore()

	productService := product.Init(firestoreClient)
	err := productService.SyncProducts(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Completed sync_products")
}
