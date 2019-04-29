package main

import (
	"net/http"
	"github.com/labstack/echo"
)

type Comment struct {
	id    int64    `json:"id"`
	Name  string   `json:"name" form:"name"`
	Text  string   `json:"text" form:"text"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello " + name)
	})
	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hello " + name)
	})
	e.POST("/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello " + name)
	})
	e.POST("/api/comments", func(c echo.Context) error {
		var comment Comment
		if err := c.Bind(&comment); err != nil {
			return c.String(http.StatusBadRequest, "Bind: " + err.Error())
		}
		return c.String(http.StatusOK, "hello")
	})
	e.GET("/api/comments", func(c echo.Context) error {
		comments := []Comment {
			{id: 1, Name: "bob", Text: "Hello"},
			{id: 2, Name: "mike", Text: "Good Morning"},
		}
		return c.JSON(http.StatusOK, comments)
	})
	e.Logger.Fatal(e.Start(":8080"))
}