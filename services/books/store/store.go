package store

import (
	"context"
	"fmt"
	"log"

	"github.com/zhetkerbaevan/library-microservice/services/books/db"
	"github.com/zhetkerbaevan/library-microservice/services/books/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookStore struct {
}

func NewBookStore() *BookStore {
	return &BookStore{}
}

func (s *BookStore) InsertBook(book *models.MongoBook) error {
	collection := db.ReturnCollectionPointer()

	_, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		return err
	}

	return nil
}

func (s *BookStore) GetBooks() ([]primitive.M, error) {
	collection := db.ReturnCollectionPointer()
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var books []primitive.M
	for cursor.Next(context.Background()) {
		var book bson.M
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

func (s *BookStore) DeleteBook(id string) error {
	collection := db.ReturnCollectionPointer()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	log.Println("Result from Deleting", res.DeletedCount)

	return nil
}

func (s *BookStore) UpdateBook(book models.MongoBook) error {
	collection := db.ReturnCollectionPointer()
	filter := bson.D{{"_id", book.ID}}
	update := bson.D{}

	if book.Name != "" {
		update = append(update, bson.E{"name", book.Name})
	}
	if book.Author != "" {
		update = append(update, bson.E{"author", book.Author})
	}
	if book.Genre != "" {
		update = append(update, bson.E{"genre", book.Genre})
	}

	if len(update) == 0 {
		return nil
	}

	update = bson.D{{"$set", update}}
	updRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error from store", err)
		return err
	}

	fmt.Println("updRes:", updRes)

	return nil
}
