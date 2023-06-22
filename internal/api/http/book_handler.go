package http

import (
	"github.com/labstack/echo/v4"
	"labs-go-api-books-hexagonal/internal/core/services"
)

type BookHandler struct {
	bookService services.IBookService
}

type IBookHandler interface {
	Save(echo.Context) error
	GetById(echo.Context) error
}

func NewBookHandler(bookService services.IBookService) IBookHandler {
	return &BookHandler{bookService}
}

func (b *BookHandler) Save(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (b *BookHandler) GetById(ctx echo.Context) error {
	//TODO implement me
	_, err := b.bookService.GetByID("")
	if err != nil {
		return err
	}

	panic("implement me")
	return nil
}
