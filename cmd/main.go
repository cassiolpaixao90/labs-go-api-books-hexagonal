package main

import (
	"labs-go-api-books-hexagonal/pkg/database/mongodb"
	"labs-go-api-books-hexagonal/pkg/server/http/echo"
)

func main() {
	m := mongodb.NewMongoDB()
	mc := m.Start()
	s := echo.NewServerEcho(mc)
	s.Start()
}
