package services

import (
	"labs-go-api-books-hexagonal/internal/core/domain"
)

type IBookService interface {
	GetByID(ID string) (*domain.Book, error)
	Save(book *domain.Book) (*domain.Book, error)
}

type BookService struct {
	bookRepository domain.BookRepository
}

func NewBookService(bookRepository domain.BookRepository) IBookService {
	return &BookService{bookRepository: bookRepository}
}

func (b *BookService) GetByID(ID string) (*domain.Book, error) {
	//TODO implement me
	return nil, nil
}

func (b *BookService) Save(book *domain.Book) (*domain.Book, error) {
	res, err := b.bookRepository.Save(book)
	if err != nil {
		return nil, err
	}
	return res, nil
}
