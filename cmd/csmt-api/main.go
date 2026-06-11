package main

import (
	"net/http"

	"github.com/cyrillemad/csmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func SearchHandler(
	client *csmt.NoAuthorizeClient) echo.HandlerFunc {

	return func(c *echo.Context) error {
		search := c.Param("search")
		hash, err := client.Community.SearchHash(search)
		if err != nil {
			return err
		}
		price, err := client.Community.PriceOverview(hash)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, price.MedianPrice)
	}
}

func main() {
	c := csmt.NewNoAuthClient()
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/:search", SearchHandler(c))

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
