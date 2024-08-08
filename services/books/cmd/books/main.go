package main

import "github.com/zhetkerbaevan/library-microservice/services/books/cmd/grpc"

func main() {
	gRPCServer := grpc.NewGRPCServer(":9000")
	gRPCServer.Run()
}
