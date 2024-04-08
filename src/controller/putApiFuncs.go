package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func PutChatMessage(c echo.Context) error {
	return c.String(http.StatusOK, "put chat message")
}
