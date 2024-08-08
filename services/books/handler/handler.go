package handler

import (
	"context"

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
		Id:     1,
		Name:   "Sapiens",
		Author: "Yuval Noah Harari",
		Genre:  "Non-fiction",
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
