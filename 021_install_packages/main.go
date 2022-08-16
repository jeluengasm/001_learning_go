package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Path
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success first endpoint")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
