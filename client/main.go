package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	individualAdapter "pedro.prado.grpc.client.example/adapter/individual"
	protoIndividual "pedro.prado.grpc.client.example/adapter/individual/proto"
)

var (
	ctx  = context.Background()
	opts []grpc.DialOption
)

func main() {

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 8950), opts...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	individualClient := protoIndividual.NewIndividualServiceClient(conn)
	individualAdapterInstance := individualAdapter.New(ctx, individualClient)

	for {
		fmt.Println("Rodando o Get Individual")
		individual, err := individualAdapterInstance.GetIndividual("1")
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		fmt.Printf("%v\n", individual)

		time.Sleep(10 * time.Second)

		fmt.Println("Rodando o List Individuals")
		individuals, err := individualAdapterInstance.ListIndividuals()
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		fmt.Printf("%v\n", individuals)

		time.Sleep(10 * time.Second)
	}
}
