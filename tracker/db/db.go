package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"tracker/config"
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
	// Create the database if it doesn't exist
	_ = client.Database(cfg.MongoDatabase).CreateCollection(ctx, "ads")
	return nil
}

func GetCollection(collectionName string, cfg *config.Config) *mongo.Collection {
	return MongoClient.Database(cfg.MongoDatabase).Collection(collectionName)
}
