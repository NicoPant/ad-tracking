package main

import (
	"fmt"
	"tracker/config"
	"tracker/db"
	"tracker/model/tracker"
	//"tracker/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg := config.LoadConfig()
	trackerService := tracker.NewTrackerService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trackerService)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	/*proto.RegisterAdServiceServer(grpcServer, &handler.AdServiceServer{
	  AdRepository: adService,
	})*/

	log.Println("gRPC AdService running on :9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
