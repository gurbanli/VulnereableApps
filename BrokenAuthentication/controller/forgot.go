package controller

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForgotPasswordController struct {

}

var ForgotPassword *ForgotPasswordController

func(fpc *ForgotPasswordController) ForgotPasswordSendVulnerable(c *gin.Context){
	var fpRequest *dto.ForgotPasswordSendRequest
	if err := c.ShouldBindJSON(&fpRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	service.Auth.SendResetToken(fpRequest.Username)
	c.JSON(http.StatusOK, gin.H{
		"message": "If that username exists in our database, we will send you an email to reset your password",
	})
}


func(fpc *ForgotPasswordController) ForgotPasswordVerify(c *gin.Context){
	var fpRequest *dto.ForgotPasswordVerifyRequest
	if err := c.ShouldBindJSON(&fpRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if service.Auth.VerifyResetToken(fpRequest.Username, fpRequest.ResetToken){
		user := service.Auth.ChangePassword(fpRequest.Username, fpRequest.NewPassword)
		c.JSON(http.StatusOK, gin.H{
			"message": "Password reset is successful !",
			"user": user,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Token is wrong !",
	})
}

func(fpc *ForgotPasswordController) ForgotPasswordSend(c *gin.Context){

}