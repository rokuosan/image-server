package service

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/rokuosan/image-server/internal/database"
	"github.com/rokuosan/image-server/internal/model"
)

func CreateImageByBase64Content(db *database.Database, content string) (*model.Image, error) {
	digest := sha256.Sum256([]byte(content))

	image := &model.Image{
		Content: content,
		Digest:  hex.EncodeToString(digest[:]),
	}

	res := db.Create(image)
	if res.Error != nil {
		return nil, res.Error
	}

	return image, nil
}

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

func FindImageByContentDigest(db *database.Database, digest string) (*model.Image, error) {
	image := &model.Image{Digest: digest}
	res := db.Find(image)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return image, nil
}

func FindImageByContentBase64(db *database.Database, content string) (*model.Image, error) {
	digest := sha256.Sum256([]byte(content))
	digestString := hex.EncodeToString(digest[:])

	return FindImageByContentDigest(db, digestString)
}
