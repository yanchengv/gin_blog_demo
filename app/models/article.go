package models

import "time"

type Article struct {
	ID        uint   `form:"id"`
	Title     string `form: "title"`
	Subtitle  string `form: "subtitle"`
	Content   string `form: "content";json: content;gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
