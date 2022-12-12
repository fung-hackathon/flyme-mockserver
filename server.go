package main

import (
	"flyme-mockserver/handler"
	"flyme-mockserver/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(logger.EchoLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/user", handler.MockUser)
	e.POST("/login", handler.MockToken)
	e.GET("/user/:userID", handler.MockUserWithParam)
	e.PUT("/user/:userID", handler.MockUserWithParam)
	e.DELETE("/user/:userID", handler.MockUserWithParam)

	e.GET("/follow/:userID", handler.MockFriends)
	e.POST("/follow/:userID", handler.MockUser)
	e.GET("/request/:userID", handler.MockRequests)
	e.POST("/follow/:userID", handler.MockUser)

	e.POST("/history/:userID", handler.MockHistory)
	e.GET("/history/:userID", handler.MockHistories)

	e.Logger.Fatal(e.Start(":8080"))
}
