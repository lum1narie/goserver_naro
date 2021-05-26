package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world.\n")
	})
	e.GET("/lum1narie", func(c echo.Context) error {
		return c.String(http.StatusOK, "@lum1narieです、こんにちは\n汝、キーボードを愛せよ\n")
	})

	e.Logger.Fatal(e.Start(":10100"))
}
