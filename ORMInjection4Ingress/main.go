package main

import (
	"ORMInjection4Ingress/controller"
	"ORMInjection4Ingress/model"
	"ORMInjection4Ingress/repository"
	"github.com/gin-gonic/gin"
)

//author: gurbanli

func main(){
	router := gin.Default()

	repo := repository.Repository{}
	repo.InitializeDatabase()


	repository.Database.AutoMigrate(&model.User{})
	repository.Database.AutoMigrate(&model.Product{})

	rController := controller.RegistrationController{}
	lController := controller.LoginController{}
	uController := controller.UserController{}
	router.POST("/register", rController.RegisterByUsername)
	router.POST("/login", lController.LoginByUsernameAndPassword)
	router.GET("/user/:id", uController.GetUser)
	router.Run(":7777")
}