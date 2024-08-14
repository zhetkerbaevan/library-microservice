package db

import (
	"context"
	"log"

	"github.com/zhetkerbaevan/library-microservice/services/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Envs.ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions) //Connect to MongoDB
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}

var Client *mongo.Client

func New(mongo *mongo.Client) {
	Client = mongo
}

func ReturnCollectionPointer() *mongo.Collection {
	collection := Client.Database("library_db").Collection("books") //Pointer to collection
	return collection
}
