package router

import (
	"BrokenAuthentication/controller"
	"BrokenAuthentication/model"
	"BrokenAuthentication/repository"
	"BrokenAuthentication/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeRouter() *gin.Engine{
	router := gin.Default()

	repository.Repo = &repository.Repository{DB: nil}
	repository.Repo.InitializeDatabaseRepository()

	store, _ := redis.NewStore(8, "tcp", "127.0.0.1:6379", "", []byte(os.Getenv("secret_key")))
	store.Options(sessions.Options{
		Path: "/",
		MaxAge:   900,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("session_id", store))



	repository.Repo.DB.AutoMigrate(&model.User{})
	repository.Repo.DB.AutoMigrate(&model.Product{})


	service.Auth = &service.AuthServiceImpl{}
	service.Session = &service.SessionServiceImpl{}
	controller.Login = &controller.LoginController{}
	controller.Register = &controller.RegisterController{}
	controller.Admin = &controller.AdminController{}

	apiV1 := router.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			auth.POST("/login", controller.Login.LoginByPassword)
			auth.POST("/register", controller.Register.RegisterByUsername)
		}

		admin := apiV1.Group("/admin")
		{
			admin.GET("/", controller.Admin.Index )
		}
	}

	return router
}
