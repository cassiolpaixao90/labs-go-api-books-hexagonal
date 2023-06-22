package services

import (
	"labs-go-api-books-hexagonal/internal/core/domain"
)

type IBookService interface {
	GetByID(ID string) (*domain.Book, error)
	Save(user *domain.Book) error
}

type BookService struct {
	bookRepository domain.BookRepository
}

func NewBookService(bookRepository domain.BookRepository) IBookService {
	return &BookService{bookRepository: bookRepository}
}

func (b BookService) GetByID(ID string) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookService) Save(user *domain.Book) error {
	//TODO implement me
	panic("implement me")
}
