package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Route() string
	GET(echo.Context) error
	POST(echo.Context) error
	PUT(echo.Context) error
	DELETE(echo.Context) error
}

func Handlers() []Handler {
	return []Handler{
		NewImageHandler(),
	}
}

func MethodNotAllowed(c echo.Context) error {
	return c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
}
