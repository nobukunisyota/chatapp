package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeleteChatMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Delete chat message")
}
