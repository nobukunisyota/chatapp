package main

import (
	"chatapp/controller"

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
	e.GET("/message/:id", controller.GetChatMessage)
	e.POST("/message/:id", controller.PostChatMessage)
	e.PUT("/message/:id", controller.PutChatMessage)
	e.DELETE("/message/:id", controller.DeleteChatMessage)

	return e
}
