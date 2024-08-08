package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
	"github.com/zhetkerbaevan/library-microservice/services/library/tmp"
	"google.golang.org/grpc"
)

type LibraryHandler struct {
	conn *grpc.ClientConn
}

func NewLibraryHandler(conn *grpc.ClientConn) *LibraryHandler {
	return &LibraryHandler{conn: conn}
}

func (h *LibraryHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/book", h.handleCreateBook)
	router.HandleFunc("/", h.handleGetBooks)
}

func (h *LibraryHandler) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	client := books.NewBookServiceClient(h.conn)

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
}

func (h *LibraryHandler) handleGetBooks(w http.ResponseWriter, r *http.Request) {
	client := books.NewBookServiceClient(h.conn)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	bs, err := client.GetBooks(ctx, &books.GetBooksRequest{})
	if err != nil {
		log.Fatalf("CLIENT ERROR: %v", err)
	}

	t := template.Must(template.New("books").Parse(tmp.BooksTemplate)) //Create a new template

	if err := t.Execute(w, bs.GetBooks()); err != nil {
		log.Fatalf("TEMPLATE ERROR: %v", err)
	}
}
