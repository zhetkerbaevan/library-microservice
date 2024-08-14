package handler

import (
	"context"
	"fmt"

	"github.com/zhetkerbaevan/library-microservice/services/books/models"
	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
	"google.golang.org/grpc"
)

type BooksGRPCHandler struct {
	bookService models.BookService //Service injection
	books.UnimplementedBookServiceServer
}

func NewGRPCHandler(grpc *grpc.Server, bookService models.BookService) {
	gRPCHandler := &BooksGRPCHandler{
		bookService: bookService,
	}

	//Register BookServiceServer
	books.RegisterBookServiceServer(grpc, gRPCHandler) //If your gRPC handler has missing implementation then we will have error here
}

// Implementation of RPC methods
func (h *BooksGRPCHandler) CreateBook(ctx context.Context, req *books.CreateBookRequest) (*books.CreateBookResponse, error) {
	book := &books.Book{ //Get Payload
		Name:   req.Name,
		Author: req.Author,
		Genre:  req.Genre,
	}

	err := h.bookService.CreateBook(ctx, book) //Send payload to service
	if err != nil {
		return nil, err
	}

	res := &books.CreateBookResponse{ //Write response
		Status: "Created",
	}

	return res, nil
}

func (h *BooksGRPCHandler) GetBooks(ctx context.Context, req *books.GetBooksRequest) (*books.GetBooksResponse, error) {
	bs, err := h.bookService.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	res := &books.GetBooksResponse{
		Books: bs,
	}
	return res, nil
}

func (h *BooksGRPCHandler) DeleteBook(ctx context.Context, req *books.DeleteBookRequest) (*books.DeleteBookResponse, error) {
	id := req.Id
	err := h.bookService.DeleteBook(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &books.DeleteBookResponse{
		Status: "Deleted",
	}

	return res, nil
}

func (h *BooksGRPCHandler) UpdateBook(ctx context.Context, req *books.UpdateBookRequest) (*books.UpdateBookResponse, error) {
	id := req.Id

	book := &books.Book{
		Id:     req.Id,
		Name:   req.Name,
		Author: req.Author,
		Genre:  req.Genre,
	}
	err := h.bookService.UpdateBook(ctx, id, book)
	if err != nil {
		fmt.Println("error from handler", err)
		return nil, err
	}
	res := &books.UpdateBookResponse{
		Status: "Updated",
	}

	return res, nil
}
