package main

import (
	"github.com/NicoPant/ad-tracking/proto"
	"github.com/NicoPant/ad-tracking/tracker/config"
	"github.com/NicoPant/ad-tracking/tracker/db"
	"github.com/NicoPant/ad-tracking/tracker/handler"
	"github.com/NicoPant/ad-tracking/tracker/model/tracker"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}*/

func main() {
	cfg := config.LoadConfig()
	trackerService := tracker.NewTrackerService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterTrackerServiceServer(grpcServer, &handler.TrackerServiceServer{
		TrackerRepository: trackerService,
	})

	log.Println("gRPC AdService running on :9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
