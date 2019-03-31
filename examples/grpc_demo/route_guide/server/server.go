package main

import (
	"context"
	"flag"
	pb "github.com/aaronize/grandonion/examples/grpc_demo/route_guide/routeguide"
	"github.com/golang/protobuf/proto"
	"math"
	"sync"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A Json file containing a list of features")
	port = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {
	savedFeatures	[]*pb.Feature

	mu 				sync.Mutex
	routeNotes 		map[string][]*pb.RouteNote
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}

	return &pb.Feature{Location: point}, nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeatureServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Low.Longitude), float64(rect.High.Longitude))
	right := math.Max(float64(rect.Low.Longitude), float64(rect.High.Longitude))
	top := math.Max(float64(rect.Low.Latitude), float64(rect.High.Latitude))
	bottom := math.Min(float64(rect.Low.Latitude), float64(rect.High.Latitude))

	if float64(point.Longitude) >= left && float64(point.Longitude) <= right && float64(point.Latitude) >= bottom && float64(point.Latitude) <= top {
		return true
	}
	return false
}
