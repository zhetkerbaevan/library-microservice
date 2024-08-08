package models

import (
	"context"

	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
)

type BookService interface {
	CreateBook(context.Context, *books.Book) error
}
