package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nnnnn81/stampy-be/controller"
	"github.com/nnnnn81/stampy-be/db"
)

func newRouter() *echo.Echo {
	godotenv.Load(".env")
	e := echo.New()
	db.Connect()
	// db.Migrate()

	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.POST("/check-email", controller.UserEmailCheck)

	r := e.Group("/auth")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", controller.Auth)
	r.GET("/user", controller.UserShow)
	r.PUT("/user", controller.UserUpdate)
	r.PUT("/user/pwd", controller.UserPassUpdate)
	r.GET("/stampcard", controller.CardsShow)
	r.GET("/stampcard/:id", controller.CardShow)
	r.POST("/stampcard", controller.CardCreate)
	r.PUT("/stampcard/:id", controller.CardUpdate)
	r.PUT("/stamp", controller.StampCreate)
	r.GET("/notice", controller.NoticesShow)
	r.GET("/notice/:id", controller.NoticeShow)
	r.GET("/letter", controller.LettersShow)
	r.GET("/letter/:id", controller.LetterShow)
	r.POST("/notice", controller.NoticeCreate)
	r.POST("/letter", controller.LetterCreate)
	r.PUT("/notice/read", controller.ReadUpdate)

	return e
}
