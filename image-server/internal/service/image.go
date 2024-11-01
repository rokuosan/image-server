package service

import (
	"github.com/rokuosan/image-server/internal/database"
	"github.com/rokuosan/image-server/internal/model"
)

func FindImageById(db *database.Database, id int64) (*model.Image, error) {
	image := &model.Image{Id: id}
	res := db.Find(image)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return image, nil
}
