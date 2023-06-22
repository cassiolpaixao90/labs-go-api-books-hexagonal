package domain

type Book struct {
	ID   string
	Name string
}

type BookRepository interface {
	GetByID(ID string) (*Book, error)
	Save(book *Book) error
}

func NewBookDomain(ID string, Name string) *Book {
	return &Book{ID, Name}
}
