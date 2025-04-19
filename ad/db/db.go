package db

import (
	"ad/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func InitMongo(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
		return err
	}

	MongoClient = client
	fmt.Println("MongoDB connected!")

	return nil
}

func GetCollection(collectionName string, cfg *config.Config) *mongo.Collection {
	return MongoClient.Database(cfg.MongoDatabase).Collection(collectionName)
}
