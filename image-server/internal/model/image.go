package model

import (
	"time"
)

type Image struct {
	Id int64 `json:"id"`

	Content string `json:"content"`
	Digest  string `json:"sha256_digest" gorm:"column:sha256_digest"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*Image) TableName() string {
	return "image"
}
