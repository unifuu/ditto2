package mdb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB    = "ditto2"
	Games *mongo.Collection
	Users *mongo.Collection
)

func Init() {
	// Use environment variable or default to Docker container name
	mongoURI := "mongodb://mongo:27017"
	if uri := os.Getenv("MONGO_URI"); uri != "" {
		mongoURI = uri
	}

	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Games = client.Database(DB).Collection("games")
	Users = client.Database(DB).Collection("users")
}
