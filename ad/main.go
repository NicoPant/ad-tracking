package main

import (
	"ad/config"
	"ad/db"
	"ad/model/ad"
	"ad/proto"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var opts []grpc.DialOption

	cfg := config.LoadConfig()
	adService := ad.NewAdService(cfg)
	err := db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewAdServiceClient(conn)
	fmt.Println(client)

	//client := pb.NewAdServiceClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	/*newAd := &ad.Ad{
		Id:          uuid.NewString(),
		Title:       "test",
		Description: "description_test",
		Url:         "test.com",
	}

	err = adService.CreateAd(context.Background(), newAd)
	if err != nil {
		log.Fatalf("Error creating ad: %v", err)
	} else {
		log.Println("Ad created successfully")
	}*/
	ad, err := adService.GetAdById(context.Background(), "20ee03e5-75e7-4362-a482-a9b2fa368947")
	if err != nil {
		log.Fatalf("Error getting ad: %v", err)
	} else {
		log.Println("Ad retrieved successfully:", ad)
	}
}
