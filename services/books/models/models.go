package models

import (
	"context"

	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoBook struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Author string             `bson:"author"`
	Genre  string             `bson:"genre"`
}

type BookService interface {
	CreateBook(context.Context, *books.Book) error
	GetBooks(context.Context) ([]*books.Book, error)
	DeleteBook(context.Context, string) error
	UpdateBook(context.Context, string, *books.Book) error
}

type BookStore interface {
	InsertBook(*MongoBook) error
	GetBooks() ([]primitive.M, error)
	DeleteBook(string) error
	UpdateBook(MongoBook) error
}
