package controller

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {

}

var Login *LoginController

func(lc *LoginController) LoginByPassword(c *gin.Context){
	var lRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	err := service.Auth.CheckUsername(lRequest.Username)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found !",
		})
		return
	}
	err = service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
	})

}

func(lc *LoginController) LoginByPasswordTiming(c *gin.Context){
	var lRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	err := service.Auth.CheckUsername(lRequest.Username)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Login Failed !",
		})
		return
	}
	err = service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Login Failed !",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
	})

}
