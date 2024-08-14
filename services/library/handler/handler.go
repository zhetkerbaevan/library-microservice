package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
	"github.com/zhetkerbaevan/library-microservice/services/library/tmp"
	"github.com/zhetkerbaevan/library-microservice/services/library/utils"
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
	router.HandleFunc("/delete/book/{id}", h.handleDeleteBook)
	router.HandleFunc("/update/book/{id}", h.handleUpdateBook)
}

func (h *LibraryHandler) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}
	client := books.NewBookServiceClient(h.conn)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var payload books.Book
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Fatalf("INVALID PAYLOAD: %v", err)
	}
	_, err := client.CreateBook(ctx, &books.CreateBookRequest{
		Name:   payload.Name,
		Author: payload.Author,
		Genre:  payload.Genre,
	})
	if err != nil {
		log.Fatalf("CLIENT ERROR: %v", err)
	}
}

func (h *LibraryHandler) handleGetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}
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

func (h *LibraryHandler) handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}
	client := books.NewBookServiceClient(h.conn)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	id := r.PathValue("id")
	_, err := client.DeleteBook(ctx, &books.DeleteBookRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("CLIENT ERROR: %v", err)
	}

}

func (h *LibraryHandler) handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}
	client := books.NewBookServiceClient(h.conn)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	id := r.PathValue("id")
	var payload books.Book
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Fatalf("INVALID PAYLOAD: %v", err)
	}
	_, err := client.UpdateBook(ctx, &books.UpdateBookRequest{
		Id:     id,
		Name:   payload.Name,
		Author: payload.Author,
		Genre:  payload.Genre,
	})
	if err != nil {
		log.Fatalf("CLIENT ERROR: %v", err)
	}
}
