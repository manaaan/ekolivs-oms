package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/backend/services/product/api"
	"github.com/manaaan/ekolivs-oms/backend/services/product/internal/product"
	"github.com/manaaan/ekolivs-oms/backend/services/product/internal/server"

	"google.golang.org/grpc"
)

var filename = ".env"

func init() {
	flag.StringVar(&filename, "f", ".env", "Requires the absolute path of the filename slash the filename. Example: /absolute_path/filename")
}

func main() {
	flag.Parse()

	if filename != ".env" {
		env.LoadEnvFromAbsolutePath(filename)
	} else {
		env.LoadEnv()
	}
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

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
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
