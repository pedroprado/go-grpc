package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "pedro.prado.grpc.server.example/infra/grpc/protoFile"
	"pedro.prado.grpc.server.example/infra/grpc/server"
)

func main() {

	savedFeatures := []*pb.Feature{
		&pb.Feature{Name: "Feature1", Location: &pb.Point{Latitue: 1, Longitue: 1}},
		&pb.Feature{Name: "Feature2", Location: &pb.Point{Latitue: 2, Longitue: 2}},
		&pb.Feature{Name: "Feature3", Location: &pb.Point{Latitue: 3, Longitue: 3}},
		&pb.Feature{Name: "Feature4", Location: &pb.Point{Latitue: 4, Longitue: 4}},
		&pb.Feature{Name: "Feature5", Location: &pb.Point{Latitue: 5, Longitue: 5}},
		&pb.Feature{Name: "Feature6", Location: &pb.Point{Latitue: 6, Longitue: 6}},
		&pb.Feature{Name: "Feature7", Location: &pb.Point{Latitue: 7, Longitue: 7}},
	}

	routeNotes := map[string][]*pb.RouteNote{
		"11": []*pb.RouteNote{{Location: &pb.Point{Latitue: 1, Longitue: 1}, Message: "message1"}},
		"22": []*pb.RouteNote{{Location: &pb.Point{Latitue: 2, Longitue: 2}, Message: "message2"}},
		"33": []*pb.RouteNote{{Location: &pb.Point{Latitue: 3, Longitue: 3}, Message: "message3"}},
		"44": []*pb.RouteNote{{Location: &pb.Point{Latitue: 4, Longitue: 4}, Message: "message4"}},
		"55": []*pb.RouteNote{{Location: &pb.Point{Latitue: 5, Longitue: 5}, Message: "message5"}},
		"66": []*pb.RouteNote{{Location: &pb.Point{Latitue: 6, Longitue: 6}, Message: "message6"}},
		"77": []*pb.RouteNote{{Location: &pb.Point{Latitue: 7, Longitue: 7}, Message: "message7"}},
	}

	grpcService := server.NewGrpcSevice(savedFeatures, routeNotes)

	var opts []grpc.ServerOption
	grpcNewServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcNewServer, grpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8950))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcNewServer.Serve(lis)
}
