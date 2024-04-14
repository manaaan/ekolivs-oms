package main

import (
	"fmt"
	"log"
	"net"

	"github.com/manaaan/ekolivs-oms/product/api"
	"google.golang.org/grpc"
)

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterProductServiceServer(grpcServer, api.UnimplementedProductServiceServer{})
	fmt.Printf("product service listening on %d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
