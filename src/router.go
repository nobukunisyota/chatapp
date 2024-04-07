package main

import (
	"src/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getRouter() *echo.Echo {
	e := echo.New()
	api := e.Group("/api/v1")

	// middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())

	//routing
	api.GET("/", controller.Index)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	return e
}
