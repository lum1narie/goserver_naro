package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	jsonData struct {
		Number int    `json:"number,omitempty`
		String string `json:"string,omitempty`
		Bool   bool   `json:"bool,omitempty`
	}

	AddRequest struct {
		Right int `json:"right"`
		Left  int `json:"left"`
	}

	AddResponse struct {
		Answer int
	}

	Class struct {
		ClassNumber int        `json:"class_number"`
		Students    []*Student `json:"students"`
	}

	Student struct {
		StudentNumber int    `json:"student_number"`
		Name          string `json:"name"`
	}
)

var (
	ClassesJson = []byte(`[
  {"class_number": 1, "students": [
    {"student_number": 1, "name": "Humming"},
    {"student_number": 2, "name": "masutech16"},
    {"student_number": 3, "name": "ninja"}
  ]},
  {"class_number": 2, "students": [
    {"student_number": 1, "name": "hukuda222"},
    {"student_number": 2, "name": "takashi_trap"},
    {"student_number": 3, "name": "nagatech"},
    {"student_number": 4, "name": "whiteonion"}
  ]},
  {"class_number": 3, "students": [
    {"student_number": 1, "name": "yamada"},
    {"student_number": 2, "name": "tubotu"},
    {"student_number": 3, "name": "tsukatomo"}
  ]},
  {"class_number": 4, "students": [
    {"student_number": 1, "name": "g2"},
    {"student_number": 2, "name": "hatasa-y"}
  ]}
]`)
	classes *[]Class = &[]Class{}
)

func main() {

	json.Unmarshal(ClassesJson, classes)

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
	e.POST("/add", addHandler)
	e.GET("/students/:class/:studentNumber", getStudentHandler)

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
	data := &jsonData{}

	if err := c.Bind(data); err != nil {
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
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
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

func addHandler(c echo.Context) error {
	addReq := &AddRequest{}

	if err := c.Bind(addReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	addRes := &AddResponse{
		Answer: addReq.Left + addReq.Right,
	}

	return c.JSON(http.StatusOK, addRes)
}

func getStudentHandler(c echo.Context) error {
	httpErrStudent := echo.NewHTTPError(http.StatusBadRequest, "Student Not Found")
	classNumber, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		return httpErrStudent
	}
	studentNumber, err := strconv.Atoi(c.Param("studentNumber"))
	if err != nil {
		return httpErrStudent
	}

	var class *Class = nil
	for _, c := range *classes {
		if c.ClassNumber == classNumber {
			class = &c
			break
		}
	}
	if class == nil {
		return httpErrStudent
	}

	var student *Student = nil
	for _, s := range class.Students {
		if s.StudentNumber == studentNumber {
			student = s
			break
		}
	}
	if student == nil {
		return httpErrStudent
	}

	return c.JSON(http.StatusOK, student)
}
