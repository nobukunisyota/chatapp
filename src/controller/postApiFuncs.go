package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func PostChatMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Post chat message")
}
