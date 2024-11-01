package model

import (
	"time"
)

type Content struct {
	Id int64 `json:"id"`

	Name      string `json:"name"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Extension string `json:"extension"`

	ImageId int64 `json:"image_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*Content) TableName() string {
	return "content"
}
