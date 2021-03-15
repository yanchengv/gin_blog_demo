package models

import "time"

type Article struct {
	ID        uint
	Title     string
	Subtitle  string
	Content   string `json: content;gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
