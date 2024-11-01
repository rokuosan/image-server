package handler

import (
	"encoding/base64"

	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/database"
	imageprocessor "github.com/rokuosan/image-server/internal/image-processor"
	"github.com/rokuosan/image-server/internal/service"
)

type ContentViewHandler struct{}

func NewContentViewHandler() *ContentViewHandler {
	return &ContentViewHandler{}
}

func (h *ContentViewHandler) Route() string {
	return "/content/:id/view"
}

func (h *ContentViewHandler) GET(c echo.Context) error {
	db := database.New()
	id, err := ParamsInt64(c, "id")
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID"})
	}

	// IDからContentを取得
	content, err := service.FindContentById(db, id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}
	if content == nil {
		return c.JSON(404, map[string]string{"error": "Not Found"})
	}

	// Contentに紐づくImageを取得
	image, err := service.FindImageById(db, content.ImageId)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}
	if image == nil {
		return c.JSON(404, map[string]string{"error": "Not Found"})
	}

	// クエリパラメータsizeがあればリサイズする
	size := c.QueryParam("size")
	if size != "" {
		w, h, err := imageprocessor.ParseSizeNxM(size)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid Size"})
		}

		img, err := imageprocessor.Base64ToImage(image.Content)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to convert base64 to image"})
		}

		img, err = imageprocessor.ResizeImage(img, w, h)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to resize image"})
		}

		encoded, err := imageprocessor.EncodeImage(img, content.Extension)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to encode image"})
		}
		return c.Blob(200, "image/"+content.Extension, encoded)
	}

	// Base64から画像に変換
	data, erc := base64.StdEncoding.DecodeString(image.Content)
	if erc != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	return c.Blob(200, "image/"+content.Extension, data)
}

func (h *ContentViewHandler) POST(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ContentViewHandler) PUT(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *ContentViewHandler) DELETE(c echo.Context) error {
	return MethodNotAllowed(c)
}
