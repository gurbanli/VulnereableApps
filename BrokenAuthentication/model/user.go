package model


type User struct {
	Model
	Username string `gorm:"unique; not null"`
	Email string `gorm:"unique; not null"`
	Password string `json:"-" gorm:"not null"`
	IsAdmin bool `json:"-" gorm:"not null"`
}