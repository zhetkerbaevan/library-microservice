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

	/*
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			client := books.NewBookServiceClient(conn)

			ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
			defer cancel()

			_, err := client.CreateBook(ctx, &books.CreateBookRequest{
				Id:     2,
				Name:   "Atomic Habits",
				Author: "James Clear",
				Genre:  "Non-fiction",
			})
			if err != nil {
				log.Fatalf("CLIENT ERROR: %v", err)
			}

			bs, err := client.GetBooks(ctx, &books.GetBooksRequest{})
			if err != nil {
				log.Fatalf("CLIENT ERROR: %v", err)
			}

			t := template.Must(template.New("books").Parse(booksTemplate)) //Create a new template

			if err := t.Execute(w, bs.GetBooks()); err != nil {
				log.Fatalf("TEMPLATE ERROR: %v", err)
			}
		})*/

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
