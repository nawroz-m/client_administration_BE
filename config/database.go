package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect DB
func Db() *mongo.Client{
	URI := os.Getenv("MONGODBURL")

	// Establish the connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}
	return client

}

