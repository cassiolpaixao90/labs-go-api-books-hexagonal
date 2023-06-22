package main

import "labs-go-api-books-hexagonal/pkg/datasource/server/http/echo"

func main() {
	s := echo.NewServerEcho()
	s.Start()
}
