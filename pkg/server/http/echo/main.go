package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"labs-go-api-books-hexagonal/internal/api/http"
	"labs-go-api-books-hexagonal/internal/core/services"
	"labs-go-api-books-hexagonal/pkg/database/mongodb"
)

type ServerEcho struct {
	mongoClient *mongo.Client
}

type IServerEcho interface {
	Start()
}

func NewServerEcho(mongoClient *mongo.Client) IServerEcho {
	return &ServerEcho{mongoClient}
}

func (s *ServerEcho) Start() {
	e := echo.New()
	e.Use(middleware.Logger())

	br := mongodb.NewBookRepository(s.mongoClient)
	bs := services.NewBookService(br)
	bh := http.NewBookHandler(bs)

	r := NewRoutes(bh)
	r.routes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
