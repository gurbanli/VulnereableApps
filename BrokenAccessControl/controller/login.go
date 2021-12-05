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

	user, err := service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}
	service.Session.SetAuthSession(c, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
		"user": *user,
	})

}