package main

import (
	"awesomeProject/database"
	_ "awesomeProject/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	db := database.Connect()
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		e.
		return c.String(http.StatusOK, "Hello World")
	})

	e.Start(":8080")
}
