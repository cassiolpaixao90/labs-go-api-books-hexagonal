package mongo

import "labs-go-api-books-hexagonal/internal/core/domain"

type BookRepository struct {
}

func NewBookRepository() domain.BookRepository {
	return &BookRepository{}
}

func (b BookRepository) GetByID(ID string) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookRepository) Save(book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}
