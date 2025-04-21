package main

import (
	"tracker/config"
	"tracker/db"
	//"tracker/handler"
	//"tracker/model/ad"
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
	//adService := ad.NewAdService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

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
