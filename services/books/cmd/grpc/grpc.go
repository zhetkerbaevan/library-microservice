package grpc

import (
	"log"
	"net"

	"github.com/zhetkerbaevan/library-microservice/services/books/handler"
	"github.com/zhetkerbaevan/library-microservice/services/books/service"
	"github.com/zhetkerbaevan/library-microservice/services/books/store"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (s *GRPCServer) Run() error {
	//Firstly create tcp connection
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("FAILED TO LISTEN: %v", err)
	}

	grpcServer := grpc.NewServer()

	bookStore := store.NewBookStore()
	//Register gRPC services
	bookService := service.NewBookService(bookStore)
	handler.NewGRPCHandler(grpcServer, bookService)

	log.Println("Starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
