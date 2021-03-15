package models

import "time"

type Tag struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
