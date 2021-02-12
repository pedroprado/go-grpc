package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
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

	fmt.Println("Rodando GetFeature")
	feature, err := client.GetFeature(ctx, &pb.Point{Latitue: 2, Longitue: 3})
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%v\n", feature)
}
