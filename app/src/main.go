package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4"
)

func main() {
	e := echo.New()

	e.Use(apmechov4.Middleware())

	e.GET("/*", func(c echo.Context) error {
		route := c.Param("*")
		msg := fmt.Sprintf("Hello Kubernetes and Elasticsearch!\nYou've hitted /%s", route)
		return c.String(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
