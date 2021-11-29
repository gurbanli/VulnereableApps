package model

//author: gurbanli

type User struct {
	BaseModel
	Username string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	IsAdmin bool
	Products []Product
}