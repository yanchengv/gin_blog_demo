package models

import "time"

type ArticleTag struct {
	ID uint
	TagId uint
	ArticleId uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
