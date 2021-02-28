package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	_ = getMongodbVersion()
	// connect to mongodb
	// get version
	// get resources
	// start web page
}

func getMongodbVersion() string {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	result := client.Database("admin").RunCommand(nil,"build")

	var vers struct{
		Version string `bson:"version"`
	}
	fmt.Println(result.Decode(vers))
	fmt.Println(vers)

	str, err := result.DecodeBytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)

	return ""
}