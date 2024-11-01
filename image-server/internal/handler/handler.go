package handler

import (
	"net/http"
	"strconv"

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
		NewContentHandler(),
		NewUploadHandler(),
		NewContentViewHandler(),
	}
}

func ParamsInt64(c echo.Context, id string) (int64, error) {
	return strconv.ParseInt(c.Param(id), 10, 64)
}

func MethodNotAllowed(c echo.Context) error {
	return c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
}
