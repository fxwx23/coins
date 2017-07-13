package main

import (
	"net/http"
	"os"

	"github.com/fxwx23/coin-checker/coin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Coin Checker\n")
	})

	e.POST("/coins", coin.All)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
