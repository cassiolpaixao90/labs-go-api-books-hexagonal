package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"labs-go-api-books-hexagonal/internal/core/domain"
)

type BookRepository struct {
	*mongo.Client
}

func NewBookRepository(mc *mongo.Client) domain.BookRepository {
	return &BookRepository{mc}
}

func (b *BookRepository) getCollection() *mongo.Collection {
	return b.Database("BooksDB").Collection("books")
}

func (b *BookRepository) GetByID(ID string) (*domain.Book, error) {
	panic("implement me")
}

func (b *BookRepository) Save(book *domain.Book) (*domain.Book, error) {
	bs := BookSchema{}
	bs.ID = primitive.NewObjectID()
	bs.Name = book.Name

	res, err := b.getCollection().InsertOne(context.TODO(), bs)
	if err != nil {
		return nil, err
	}

	bm := &domain.Book{}
	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return bm, err
}
