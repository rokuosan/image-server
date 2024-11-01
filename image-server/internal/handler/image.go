package handler

import "github.com/labstack/echo/v4"

type ImageHandler struct{}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

func (h *ImageHandler) Route() string {
	return "/image"
}

func (h *ImageHandler) GET(c echo.Context) error {
	return c.String(200, "GET")
}

func (h *ImageHandler) POST(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ImageHandler) PUT(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ImageHandler) DELETE(c echo.Context) error {
	return MethodNotAllowed(c)
}
