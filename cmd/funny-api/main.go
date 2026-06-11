package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func HelloUser(c *echo.Context) error {
	name := c.QueryParam("name")
	response := StringResponse{}
	if name == "" {
		response.String = "Hello, unknow user"
		response.Tip = "You can also send you name as query 'name'"
		return c.JSON(
			http.StatusOK,
			response)
	}
	response.String = fmt.Sprintf("Hello, %s", name)
	return c.JSON(
		http.StatusOK,
		response)
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/fun/hello", HelloUser)

	if err := e.Start(":2287"); err != nil {
		e.Logger.Error(
			"failed to start server",
			"error",
			err)
	}
}
