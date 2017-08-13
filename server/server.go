package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	pb "github.com/stephen-bartell/routeguide/api"
)

var (
	jsonDBFile = flag.String("json_db_file", "db/features.json", "A json file containing a list of Features")
	port       = flag.Int("port", 10000, "The server port")
)

type guideServer struct {
	savedFeatures []*pb.Feature
	routeNotes    map[string][]*pb.RouteNote
}

// GetFeatures returns the feature at a given point
func (s *guideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, f := range s.savedFeatures {
		if proto.Equal(f.Location, point) {
			return f, nil
		}
	}

	// No feature found, return an unnamed one
	return &pb.Feature{Location: point}, nil
}

// ListFeatures lists all features contained within the given Rectangle
func (s *guideServer) ListFeatures(rect *pb.Rectangle, stream pb.Guide_ListFeaturesServer) error {
	for _, f := range s.savedFeatures {
		if inRange(f.Location, rect) {
			if err := stream.Send(f); err != nil {
				return err
			}
		}
	}
	return nil
}

// RecordRoute records a route composited of a sequence of points
func (s *guideServer) RecordRoute(stream pb.Guide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *pb.Point
	startTime := time.Now()

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++

		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if lastPoint != nil {
			distance += calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

// RouteChat receives a stream of message/location pairs and responds with a stream of all
// previous messages at each of those locations.
func (s *guideServer) RouteChat(stream pb.Guide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := serialize(in.Location)
		if _, present := s.routeNotes[key]; !present {
			s.routeNotes[key] = []*pb.RouteNote{in}
		} else {
			s.routeNotes[key] = append(s.routeNotes[key], in)
		}

		for _, note := range s.routeNotes[key] {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func (s *guideServer) loadFeatures(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	if err := json.Unmarshal(file, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func inRange(p *pb.Point, r *pb.Rectangle) bool {
	left := math.Min(float64(r.Lo.Longitude), float64(r.Hi.Longitude))
	right := math.Max(float64(r.Lo.Longitude), float64(r.Hi.Longitude))
	top := math.Max(float64(r.Lo.Latitude), float64(r.Hi.Latitude))
	bottom := math.Max(float64(r.Lo.Latitude), float64(r.Hi.Latitude))

	if float64(p.Longitude) >= left &&
		float64(p.Longitude) <= right &&
		float64(p.Latitude) >= bottom &&
		float64(p.Latitude) <= top {
		return true
	}
	return false
}

// calcDistance calculates the distance between two points using the "haversine" formula.
// This code was taken from http://www.movable-type.co.uk/scripts/latlong.html.
func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R float64 = float64(6371000) // metres
	lat1 := float64(p1.Latitude) / CordFactor
	lat2 := float64(p2.Latitude) / CordFactor
	lng1 := float64(p1.Longitude) / CordFactor
	lng2 := float64(p2.Longitude) / CordFactor
	φ1 := toRadians(lat1)
	φ2 := toRadians(lat2)
	Δφ := toRadians(lat2 - lat1)
	Δλ := toRadians(lng2 - lng1)

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func newServer() *guideServer {
	s := new(guideServer)
	s.loadFeatures(*jsonDBFile)
	s.routeNotes = make(map[string][]*pb.RouteNote)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on :%v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
