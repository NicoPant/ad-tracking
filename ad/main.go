package main

import (
	"github.com/NicoPant/ad-tracking/ad/config"
	"github.com/NicoPant/ad-tracking/ad/db"
	"github.com/NicoPant/ad-tracking/ad/handler"
	"github.com/NicoPant/ad-tracking/ad/model/ad"
	"github.com/NicoPant/ad-tracking/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	adService := ad.NewAdService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	conn, err := grpc.NewClient("trackerservice:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial AdService: %v", err)
	}
	trackerClient := proto.NewTrackerServiceClient(conn)

	grpcServer := grpc.NewServer()

	proto.RegisterAdServiceServer(grpcServer, &handler.AdServiceServer{
		AdRepository:  adService,
		TrackerClient: trackerClient,
	})

	log.Println("gRPC AdService running on :8000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
