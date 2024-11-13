package main

import (
	"GSJA/db"
	"GSJA/routes"
	// "net/http"

	_ "github.com/labstack/echo/v4"
)

func main() {

	db.Init()
	e := routes.Init()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Logger.Fatal(e.Start(":8080"))
}
