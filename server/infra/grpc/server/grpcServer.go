package server

import (
	"context"
	"fmt"
	"io"
	"math"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
	pb "pedro.prado.grpc.server.example/infra/grpc/protoFile"
)

type grpcServer struct {
	pb.UnimplementedRouteGuideServer
	mu            sync.Mutex
	savedFeatures []*pb.Feature
	routeNotes    map[string][]*pb.RouteNote
}

func NewGrpcSevice(savedFeatures []*pb.Feature, routeNotes map[string][]*pb.RouteNote) pb.RouteGuideServer {
	return &grpcServer{
		savedFeatures: savedFeatures,
		routeNotes:    routeNotes,
	}
}

func (ref *grpcServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {

	for _, feature := range ref.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

//server side streaming
func (ref *grpcServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {

	for _, feature := range ref.savedFeatures {

		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}

	}
	return nil
}

//client side streaming
func (ref *grpcServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *pb.Point

	startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endtime := time.Now()
			routeSummary := &pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endtime.Sub(startTime).Seconds()),
			}
			return stream.SendAndClose(routeSummary)

		}
		if err != nil {
			return err
		}

		pointCount++
		for _, feature := range ref.savedFeatures {
			if proto.Equal(feature.Location, point) {
				pointCount++
			}
		}

		if lastPoint != nil {
			distance += calculateDistance(lastPoint, point)
		}
		lastPoint = point
	}

}

//bidirectional streaming
func (ref *grpcServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(note.Location)

		ref.mu.Lock()
		ref.routeNotes[key] = append(ref.routeNotes[key], note)
		rn := make([]*pb.RouteNote, len(ref.routeNotes[key]))
		copy(rn, ref.routeNotes[key])
		ref.mu.Unlock()

		for _, routeNote := range rn {
			if err := stream.Send(routeNote); err != nil {
				return err
			}
		}

	}
}

func inRange(point *pb.Point, rect *pb.Rectangle) bool {

	left := math.Min(float64(rect.Lo.Longitue), float64(rect.Hi.Longitue))
	right := math.Max(float64(rect.Lo.Longitue), float64(rect.Hi.Longitue))
	top := math.Max(float64(rect.Lo.Latitue), float64(rect.Hi.Latitue))
	bottom := math.Min(float64(rect.Lo.Latitue), float64(rect.Hi.Latitue))

	if float64(point.Longitue) >= left &&
		float64(point.Longitue) <= right &&
		float64(point.Latitue) >= bottom &&
		float64(point.Latitue) <= top {
		return true
	}
	return false
}

func calculateDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitue) / CordFactor)
	lat2 := toRadians(float64(p2.Latitue) / CordFactor)
	lng1 := toRadians(float64(p1.Longitue) / CordFactor)
	lng2 := toRadians(float64(p2.Longitue) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func serialize(point *pb.Point) string {
	return fmt.Sprint("%d%d", point.Latitue, point.Longitue)
}
