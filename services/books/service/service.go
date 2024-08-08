package service

import (
	"context"

	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
)

// There is our business logic
var booksDB = make([]*books.Book, 0)

type BookService struct {
	//Store dependency injection
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) CreateBook(ctx context.Context, book *books.Book) error {
	booksDB = append(booksDB, book)
	return nil
}
