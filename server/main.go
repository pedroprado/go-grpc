package main

import (
	"fmt"
	"log"
	"net"

	"pedro.prado.grpc.server.example/presentation/individual/proto"

	"google.golang.org/grpc"
	individualService "pedro.prado.grpc.server.example/core/useCases/individual"
	individualGrpcService "pedro.prado.grpc.server.example/presentation/individual"
)

func main() {

	individualServiceInstance := individualService.New(nil)
	individualGrpcServiceInstance := individualGrpcService.New(individualServiceInstance)

	var opts []grpc.ServerOption
	newGrpcServer := grpc.NewServer(opts...)
	proto.RegisterIndividualServiceServer(newGrpcServer, individualGrpcServiceInstance)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8950))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newGrpcServer.Serve(lis)
}
