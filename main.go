package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonData struct {
	Number int    `json:"number,omitempty`
	String string `json:"string,omitempty`
	Bool   bool   `json:"bool,omitempty`
}

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world.\n")
	})

	e.GET("/hello/:username", helloHandler)

	e.GET("/lum1narie", func(c echo.Context) error {
		return c.String(http.StatusOK, "@lum1narieです、こんにちは\n汝、キーボードを愛せよ\n")
	})

	e.GET("/json", jsonHandler)

	e.POST("/post", postHandler)

	e.Logger.Fatal(e.Start(":10100"))
}

func jsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}

	return c.JSON(http.StatusOK, &res)
}

func postHandler(c echo.Context) error {
	data := new(jsonData)
	err := c.Bind(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}

func helloHandler(c echo.Context) error {
	userID := c.Param("username")
	return c.String(http.StatusOK, "Hello, "+userID+".\n")
}
