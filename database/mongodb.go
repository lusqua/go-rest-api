package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db_connect() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@cluster0.3agtn.mongodb.net/$"+os.Getenv("DB_NAME")+"?retryWrites=true&w=majority")

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

	log.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client

var Contacts *mongo.Collection