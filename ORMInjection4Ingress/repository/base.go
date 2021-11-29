package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
//author: gurbanli

type Repository struct {

}

var Database *gorm.DB

func (r *Repository) InitializeDatabase(){
	dsn := "root:@tcp(127.0.0.1:3306)/ingress_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Database = db
}