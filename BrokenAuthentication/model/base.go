package model

import "time"

type Model struct {
	ID uint `json:"-" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
