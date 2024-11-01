package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rokuosan/image-server/internal/database"
	imageprocessor "github.com/rokuosan/image-server/internal/image-processor"
	"github.com/rokuosan/image-server/internal/model"
	"github.com/rokuosan/image-server/internal/service"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) Route() string {
	return "/upload"
}

func (h *UploadHandler) GET(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *UploadHandler) POST(c echo.Context) error {
	db := database.New()

	// Content-Type チェック
	contentType := c.Request().Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "application/x-www-form-urlencoded" {
		return c.JSON(400, map[string]string{"error": "Invalid Content-Type"})
	}

	// リクエストをContent-Typeに応じて解析
	content := model.Content{}
	if contentType == "application/json" {
		body := &struct {
			Name      string `json:"name"`
			Content   string `json:"content"`
			Extension string `json:"extension"`
		}{}
		if err := c.Bind(body); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid Request Body"})
		}

		// 既存の画像かどうか確認する
		image, err := service.FindImageByContentBase64(db, body.Content)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal Server Error"})
		}
		if image == nil {
			// 既存画像がない場合は作成する
			image, err = service.CreateImageByBase64Content(db, body.Content)
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Internal Server Error"})
			}
		}

		// Base64をimage.Imageに変換
		img, err := imageprocessor.Base64ToImage(body.Content)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid Base64"})
		}

		// 画像を保存
		content.Name = body.Name
		content.Extension = body.Extension
		content.Width = img.Bounds().Dx()
		content.Height = img.Bounds().Dy()
		content.ImageId = image.Id
	} else {
		return c.JSON(400, map[string]string{"error": "Invalid Content-Type"})
	}

	result := db.Create(&content)
	if result.Error != nil {
		return c.JSON(500, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(201, content)
}

func (h *UploadHandler) PUT(c echo.Context) error {
	return MethodNotAllowed(c)
}

func (h *UploadHandler) DELETE(c echo.Context) error {
	return MethodNotAllowed(c)
}
