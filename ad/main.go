package main

import (
	"ad/config"
	"ad/db"
	"ad/handler"
	"ad/model/ad"
	"ad/proto"
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
	adService := ad.NewAdService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAdServiceServer(grpcServer, &handler.AdServiceServer{
		AdRepository: adService,
	})

	log.Println("gRPC AdService running on :8000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
