package controller

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/repository"
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OTPController struct {

}

var OTP *OTPController

func (oC *OTPController) OTPVerifyVulnerable(c *gin.Context){
	var otpVerifyRequest dto.OTPVerifyRequest
	if err := c.ShouldBindJSON(&otpVerifyRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if service.Auth.VerifyOTP(otpVerifyRequest.Username, otpVerifyRequest.OTPCode){
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Successful !",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "OTP Code is wrong !",
	})

}

func (oC *OTPController) OTPVerify(c *gin.Context){
	var otpVerifyRequest dto.OTPVerifyRequest
	if err := c.ShouldBindJSON(&otpVerifyRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if service.Auth.VerifyOTP(otpVerifyRequest.Username, otpVerifyRequest.OTPCode){
		if repository.Repo.GetUserBlockStatus(otpVerifyRequest.Username, false){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "User is blocked !",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Successful !",
		})
		return
	}
	if repository.Repo.GetUserBlockStatus(otpVerifyRequest.Username, true){
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User is blocked !",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "OTP Code is wrong !",
	})

}