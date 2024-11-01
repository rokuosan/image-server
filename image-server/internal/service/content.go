package service

import (
	"github.com/rokuosan/image-server/internal/database"
	"github.com/rokuosan/image-server/internal/model"
)

func FindContentById(db *database.Database, id int64) (*model.Content, error) {
	content := &model.Content{Id: id}
	res := db.Find(content)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return content, nil
}
