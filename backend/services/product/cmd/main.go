package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/product/api"
	"github.com/manaaan/ekolivs-oms/product/internal/product"
	"github.com/manaaan/ekolivs-oms/product/internal/server"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()
	firestoreClient := gcp.InitFirestore()
	productService, err := product.New(firestoreClient)
	if err != nil {
		log.Fatalf("Unable to initialize product service")
		return
	}

	port, err := strconv.Atoi(env.Required("PORT"))
	if err != nil {
		log.Fatalf("failed to convert port to number: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterProductServiceServer(grpcServer, server.Server{
		ProductService: productService,
	})
	fmt.Printf("product service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
