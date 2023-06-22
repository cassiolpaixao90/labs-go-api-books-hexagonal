package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"labs-go-api-books-hexagonal/internal/api/http"
	"labs-go-api-books-hexagonal/internal/core/services"
	"labs-go-api-books-hexagonal/pkg/datasource/database/mongo"
)

type ServerEcho struct {
}

type IServerEcho interface {
	Start()
}

func NewServerEcho() IServerEcho {
	return &ServerEcho{}
}

func (s *ServerEcho) Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	br := mongo.NewBookRepository()
	bs := services.NewBookService(br)
	bh := http.NewBookHandler(bs)

	r := NewRoutes(bh)
	r.routes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
