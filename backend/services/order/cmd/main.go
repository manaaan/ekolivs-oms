package main

import (
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

func main() {
	env.LoadEnv()
	firestoreClient := gcp.InitFirestore()
	orderService := order.New(firestoreClient)

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
	order_api.RegisterOrderServiceServer(grpcServer, server.Server{
		OrderService: orderService,
	})
	fmt.Printf("order service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
