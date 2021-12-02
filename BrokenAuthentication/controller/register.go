package controller

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct {

}

var Register *RegisterController

func (rController *RegisterController) RegisterByUsername(c *gin.Context){
	var rRequest dto.RegisterRequest
	if err := c.ShouldBindJSON(&rRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	user, err := service.Auth.RegisterUser(rRequest)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not register a user !",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful!",
		"user": user,
	})
}