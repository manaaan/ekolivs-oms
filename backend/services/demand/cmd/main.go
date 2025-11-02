package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
	"github.com/manaaan/ekolivs-oms/backend/specs/demand_api"
	"google.golang.org/grpc"

	"github.com/manaaan/ekolivs-oms/backend/pkg/gcp"
	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/demand"
	"github.com/manaaan/ekolivs-oms/backend/services/demand/internal/server"
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
	demandService := demand.New(firestoreClient)

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
	demand_api.RegisterDemandServiceServer(grpcServer, server.Server{
		DemandService: demandService,
	})

	fmt.Printf("demand service listening on %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
