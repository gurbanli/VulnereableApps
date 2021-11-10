package main

import (
	"SQLInjection/controller"
	"SQLInjection/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	repo := repository.Repository{}
	repo.InitializeDatabase()


	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))


	loginController := controller.LoginController{}
	productController := controller.ProductController{}

	router.POST("/login", loginController.LoginByUsername)
	router.GET("/products/:id", productController.GetProduct)

	router.Run(":4444")
}
