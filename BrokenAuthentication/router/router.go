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

	repository.Repo = &repository.Repository{DB: nil, IPCache: nil, UserCache: nil, OTPCache: nil, ResetTokenCache: nil}
	repository.Repo.InitializeDatabaseRepository()
	repository.Repo.InitializeIPCacheRepository()
	repository.Repo.InitializeUserCacheRepository()
	repository.Repo.InitializeOTPCacheRepository()
	repository.Repo.InitializeResetTokenCacheRepository()

	repository.Repo.DB.AutoMigrate(&model.User{})
	repository.Repo.ClearCache()

	service.Auth = &service.AuthServiceImpl{}




	controller.Login = &controller.LoginController{}
	controller.Register = &controller.RegisterController{}
	controller.OTP = &controller.OTPController{}
	controller.ForgotPassword = &controller.ForgotPasswordController{}

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/login/enum", controller.Login.LoginByPasswordEnum) // user enumeration
		apiV1.POST("/login/timing", controller.Login.LoginByPasswordTiming) // user enumeration via response time
		apiV1.POST("/login/ip", controller.Login.LoginByPasswordBlockIpVulnerable) // ip lockout bypass
		apiV1.POST("/login/ip2", controller.Login.LoginByPasswordBlockIpVulnerable2) // ip lockout bypass with valid account
		apiV1.POST("/login/ip3", controller.Login.LoginByPasswordBlockIpVulnerable3) // ip lockout bypass with X-Real-Ip or X-Forwarded-For
		apiV1.POST("/login", controller.Login.LoginByPassword) // not vulnerable

		apiV1.POST("/otp/verify", controller.OTP.OTPVerifyVulnerable) // vulnerable to bruteforce
		apiV1.POST("/otp/verify2", controller.OTP.OTPVerify) // not vulnerable to brute force, but still have issues


		apiV1.POST("/forgot/send/vulnerable", controller.ForgotPassword.ForgotPasswordSendVulnerable) // vulnerable to insecure random generation vulnerability
		apiV1.POST("/forgot/send", controller.ForgotPassword.ForgotPasswordSend)

		apiV1.POST("/forgot/verify", controller.ForgotPassword.ForgotPasswordVerify)


		apiV1.POST("/register", controller.Register.RegisterByUsername)
	}
	return router
}
