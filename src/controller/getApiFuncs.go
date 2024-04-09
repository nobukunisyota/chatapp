package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "API is running!")
}

func GetChatMessage(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id+" is requested")
}
