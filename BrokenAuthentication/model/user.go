package model


type User struct {
	Model
	Username string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	IsAdmin bool `gorm:"not null"`
}