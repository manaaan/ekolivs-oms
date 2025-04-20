package main

import (
	"fmt"
	"log"
	"net"

	"github.com/manaaan/ekolivs-oms/demand/api"
	"github.com/manaaan/ekolivs-oms/demand/internal/demand"
	"github.com/manaaan/ekolivs-oms/demand/internal/server"
	"github.com/manaaan/ekolivs-oms/pkg/gcp"

	"google.golang.org/grpc"
)

func main() {
	firestoreClient := gcp.InitFirestore()
	demandService := demand.New(firestoreClient)

	port := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterDemandServiceServer(grpcServer, server.Server{
		DemandService: demandService,
	})
	fmt.Printf("demand service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
