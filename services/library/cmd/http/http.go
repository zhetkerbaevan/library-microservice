package http

import (
	"log"
	"net/http"

	"github.com/zhetkerbaevan/library-microservice/services/library/cmd/grpc"
	"github.com/zhetkerbaevan/library-microservice/services/library/handler"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()

	conn := grpc.NewGRPCClient(":9000")
	defer conn.Close()

	libraryHandler := handler.NewLibraryHandler(conn)
	libraryHandler.RegisterRoutes(router)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
