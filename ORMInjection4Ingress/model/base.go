package model

import "time"

//author: gurbanli

type BaseModel struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}