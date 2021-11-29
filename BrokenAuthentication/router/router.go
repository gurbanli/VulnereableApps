package router

import (
	"BrokenAuthentication/controller"
	"BrokenAuthentication/model"
	"BrokenAuthentication/repository"
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine{
	router := gin.Default()

	repository.Repo = &repository.Repository{DB: nil}
	repository.Repo.InitializeRepository()
	repository.Repo.DB.AutoMigrate(&model.User{})

	service.Auth = &service.AuthServiceImpl{}


	controller.Login = &controller.LoginController{}
	controller.Register = &controller.RegisterController{}

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login", controller.Login.LoginByPasswordTiming)
		apiV1.POST("/register", controller.Register.RegisterByUsername)
	}
	return router
}
