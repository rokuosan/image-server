package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/database"
	"github.com/rokuosan/image-server/internal/model"
)

type ImageHandler struct{}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

func (h *ImageHandler) Route() string {
	return "/image/:id"
}

func (h *ImageHandler) GET(c echo.Context) error {
	db := database.New()
	id, err := ParamsInt64(c, "id")
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID"})
	}

	image := &model.Image{Id: id}
	res := db.Find(image)
	if res.Error != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(200, image)
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
