package main

import (
	"fmt"
	"log"
	"net"

	"github.com/manaaan/ekolivs-oms/pkg/env"
	"github.com/manaaan/ekolivs-oms/product/api"
	"github.com/manaaan/ekolivs-oms/product/internal/product"
	"github.com/manaaan/ekolivs-oms/product/internal/server"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()
	productService, err := product.New()
	if err != nil {
		log.Fatalf("Unable to initialize product service")
		return
	}

	port := 8080
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
