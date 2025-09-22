package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"

	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/demand"
	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/server"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()

	firestoreClient := gcp.InitFirestore()
	demandService := demand.New(firestoreClient)

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
	demand_api.RegisterDemandServiceServer(grpcServer, server.Server{
		DemandService: demandService,
	})
	fmt.Printf("demand service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
