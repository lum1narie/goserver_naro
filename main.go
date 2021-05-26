package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong\n")
	})

	e.GET("/fizzbuzz", fizzBuzzHandler)

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
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", data))
	}
	return c.JSON(http.StatusOK, data)
}

func helloHandler(c echo.Context) error {
	userID := c.Param("username")
	return c.String(http.StatusOK, "Hello, "+userID+".\n")
}

func fizzBuzzHandler(c echo.Context) error {
	countStr := c.QueryParam("count")

	if countStr == "" {
		countStr = "30"
	}
	count, err := strconv.Atoi(countStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", countStr))
	}

	lines := []string{}
	for i := 1; i <= count; i++ {
		if i%15 == 0 {
			lines = append(lines, "FizzBuzz")
		} else if i%3 == 0 {
			lines = append(lines, "Fizz")
		} else if i%5 == 0 {
			lines = append(lines, "Buzz")
		} else {
			strI := strconv.Itoa(i)
			lines = append(lines, strI)
		}
	}

	return c.String(http.StatusOK, strings.Join(lines, "\n"))
}
