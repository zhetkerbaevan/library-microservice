package service

import (
	"context"
	"fmt"

	"github.com/zhetkerbaevan/library-microservice/services/books/models"
	"github.com/zhetkerbaevan/library-microservice/services/common/genproto/books"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// There is our business logic

type BookService struct {
	bookStore models.BookStore //Store dependency injection
}

func NewBookService(bookStore models.BookStore) *BookService {
	return &BookService{bookStore: bookStore}
}

func (s *BookService) CreateBook(ctx context.Context, book *books.Book) error {
	mongoBook := &models.MongoBook{
		Name:   book.Name,
		Author: book.Author,
		Genre:  book.Genre,
	}
	err := s.bookStore.InsertBook(mongoBook)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBooks(ctx context.Context) ([]*books.Book, error) {
	booksData, err := s.bookStore.GetBooks()
	if err != nil {
		return nil, err
	}

	//[]primitive.M into []Book
	var bs []*books.Book
	for _, bookData := range booksData {
		book := &books.Book{
			Id:     bookData["_id"].(primitive.ObjectID).Hex(),
			Name:   bookData["name"].(string),
			Author: bookData["author"].(string),
			Genre:  bookData["genre"].(string),
		}
		bs = append(bs, book)
	}

	return bs, nil
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	err := s.bookStore.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) UpdateBook(ctx context.Context, id string, book *books.Book) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateBook := models.MongoBook{
		ID:     objID,
		Name:   book.Name,
		Author: book.Author,
		Genre:  book.Genre,
	}
	err = s.bookStore.UpdateBook(updateBook)
	if err != nil {
		fmt.Println("error from service", err)
		return err
	}

	return nil
}
