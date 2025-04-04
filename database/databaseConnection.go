package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoDbUri := os.Getenv("MONGODB_URI")
	log.Print(os.Getenv("MONGODB_URI"), "MONGODB_URI")
	if mongoDbUri == "" {
		log.Fatal("DB connection url not present")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri))

	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	fmt.Println("connected to mongodb")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
