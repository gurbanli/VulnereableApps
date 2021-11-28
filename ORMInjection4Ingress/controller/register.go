package controller

import (
	"ORMInjection4Ingress/dto"
	"ORMInjection4Ingress/model"
	"ORMInjection4Ingress/repository"
	"ORMInjection4Ingress/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//author: gurbanli

type RegistrationController struct{
}

func (r *RegistrationController) RegisterByUsername(c *gin.Context){
	var registerRequest dto.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	user := model.User{
		Username:  registerRequest.Username,
		Password:  utils.GeneratePasswordHash(registerRequest.Password),
		IsAdmin:   registerRequest.IsAdmin,
	}
	db := repository.Database
	result := db.Create(&user)
	if result.Error != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"Failed to register user !",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Created Successfully !",
		"user": user,
	})
}