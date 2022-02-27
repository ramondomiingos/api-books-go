package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:“ID“ gorm:“primayKey“`
	Name        string         `json:“name“`
	Description string         `json:“description“`
	MediumPrice string         `json:“medium_price“`
	Author      string         `json:“author“`
	ImgUrl      string         `json:“img_url“`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
