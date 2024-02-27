package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nnnnn81/stampy-be/db"
)

func newRouter() *echo.Echo {
	godotenv.Load(".env")
	e := echo.New()
	db.Connect()

	e.Use(middleware.CORS())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.POST("user", controller.Create())
	// e.GET("user/:id", controller.UserShow())
	// e.PUT("user/:id", controller.Update())
	return e
}
