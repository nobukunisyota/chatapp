package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeleteChatMessage(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id+" is deleted")
}
