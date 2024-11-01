package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/database"
	"github.com/rokuosan/image-server/internal/model"
)

type ContentHandler struct{}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{}
}

func (h *ContentHandler) Route() string {
	return "/content/:id"
}

func (h *ContentHandler) GET(c echo.Context) error {
	db := database.New()
	id, err := ParamsInt64(c, "id")
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID"})
	}

	content := &model.Content{Id: id}
	res := db.Find(content)
	if res.Error != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(200, content)
}

func (h *ContentHandler) POST(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ContentHandler) PUT(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ContentHandler) DELETE(c echo.Context) error {
	return MethodNotAllowed(c)
}
