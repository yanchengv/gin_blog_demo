package models

import (
	myutils "go_mars/lib"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("password", myutils.Md5(u.Password))
	return
}
