package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/backend/services/order/internal/order"
	"github.com/manaaan/ekolivs-oms/backend/services/order/internal/server"
	"github.com/manaaan/ekolivs-oms/backend/services/order/order_api"

	"google.golang.org/grpc"
)

var fileName = ".env"

func init() {
	flag.StringVar(&fileName, "f", ".env", "Requires the absolute path of the filename slash the filename. Example: /absolute_path/filename")
}

func main() {
	flag.Parse()
	env.LoadEnv(fileName)

	firestoreClient := gcp.InitFirestore()
	orderService := order.New(firestoreClient)

	port, err := strconv.Atoi(env.Required("PORT"))
	if err != nil {
		log.Fatalf("failed to convert port to number: %v", err)
	}

	//nolint
	// deactivate the linting of the below
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	order_api.RegisterOrderServiceServer(grpcServer, server.Server{
		OrderService: orderService,
	})
	fmt.Printf("order service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
