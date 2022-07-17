package model

import (
    "time"
)


type User struct {
	Sequence uint `gorm:"autoIncrement:true"`
	Id string `gorm:"primary_key"`
	Name string
	Level int
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
