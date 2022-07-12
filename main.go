package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ha-dns/controller"
	"net/http"
)

func main() {

	e := echo.New()

	/* middleware to log requests */
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method}: ${uri} status=${status}\n",
	}))

	/* App Routes */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello DNS!")
	})
	e.GET("/api/v1/dns", controller.GetLocation)
	e.Logger.Fatal(e.Start(":7001"))
}
