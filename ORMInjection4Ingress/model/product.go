package model

//author: gurbanli


type Product struct {
	BaseModel
	ProductName string
	ProductType string
	Count uint
	UserId uint
}