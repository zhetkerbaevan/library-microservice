package main

import (
	"context"
	"log"
	"time"

	"github.com/zhetkerbaevan/library-microservice/services/books/cmd/grpc"
	"github.com/zhetkerbaevan/library-microservice/services/books/db"
)

func main() {

	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db.New(mongoClient)
	gRPCServer := grpc.NewGRPCServer(":9000")
	gRPCServer.Run()
}
