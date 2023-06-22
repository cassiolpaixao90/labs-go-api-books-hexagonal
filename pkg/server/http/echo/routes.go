package echo

import (
	"github.com/labstack/echo/v4"
	"labs-go-api-books-hexagonal/internal/api/http"
)

type Routes struct {
	http.IBookHandler
}

type IRoutes interface {
	routes(g *echo.Echo)
}

func NewRoutes(bh http.IBookHandler) *Routes {
	return &Routes{bh}
}

func (r *Routes) routes(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/books", r.Save)
	g.GET("/books/:id", r.GetById)
}
