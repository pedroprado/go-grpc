package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	protoIndividual "pedro.prado.grpc.client.example/adapter/individual/proto"
	pb "pedro.prado.grpc.client.example/infra/grpc/protoFile"
)

func main() {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 8950), opts...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	ctx := context.Background()

	individualClient := protoIndividual.NewIndividualServiceClient(conn)

	fmt.Println("Rodando GetFeature")
	feature, err := client.GetFeature(ctx, &pb.Point{Latitue: 2, Longitue: 3})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%v\n", feature)

	fmt.Println("Rodando o Get Individuals: Id=1")
	individual, err := individualClient.GetIndividual(ctx, &protoIndividual.GetQuery{Id: "1"})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%v\n", individual)

	fmt.Println("Rodando o List Individuals")
	individuals, err := individualClient.ListIndividuals(ctx, &protoIndividual.ListQuery{Id: "1"})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%v\n", individuals)
}
