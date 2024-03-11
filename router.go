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
	// e.POST("signup", controller.Signup())
	// e.POST("login", controller.Login())

	// r := e.Group("/auth")
	// r.GET("", handler.Auth)
	// r.GET("user", controller.UserShow())
	// r.PUT("user", controller.UserUpdate())
	// r.PUT("user/pwd", controller.UserPassUpdate())
	// r.GET("stampcard", controller.CardShow())
	// r.POST("stampcard", controller.CardCreate())
	// r.PUT("stampcard/:id", controller.CardUpdate())
	// r.POST("stamp", controller.StampCreate())
	// r.GET("notice", controller.NoticeShow())
	// r.POST("notice", controller.NoticeCreate())

	return e
}
