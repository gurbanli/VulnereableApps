package repository

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/model"
	"BrokenAuthentication/util"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func (r *Repository) InitializeDatabaseRepository(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_name") )
	r.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
}

func (r *Repository) CreateUser(rRequest dto.RegisterRequest) model.User{
	user := model.User{
		Username: rRequest.Username,
		Password: util.Hash(rRequest.Password),
		Email: rRequest.Email,
		IsAdmin:  *rRequest.IsAdmin,
	}
	r.DB.Create(&user)
	return user
}

func (r *Repository) FindByUsername(username string) model.User{
	var user model.User
	r.DB.Where("username = ? ", username).First(&user)
	return user
}

func (r *Repository) FindByUsernameAndPassword(username string, password string) model.User{
	var user model.User
	r.DB.Where("username = ? and password = ?", username, util.Hash(password)).First(&user)
	return user
}

func (r *Repository) ChangeUserPassword(username string, password string) *model.User{
	var user model.User
	r.DB.Where("username = ? ", username).First(&user)
	user.Password = util.Hash(password)
	return &user
}
