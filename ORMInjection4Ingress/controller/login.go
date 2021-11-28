package controller

import (
	"ORMInjection4Ingress/dto"
	"ORMInjection4Ingress/model"
	"ORMInjection4Ingress/repository"
	"ORMInjection4Ingress/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//author: gurbanli


type LoginController struct {

}

func (lc *LoginController) LoginByUsernameAndPassword(c *gin.Context){
	var lRequest dto.LoginRequest
	var user model.User
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	db := repository.Database
	query := fmt.Sprintf("username = '%s' and password = '%s'", lRequest.Username, utils.GeneratePasswordHash(lRequest.Password))
	db.Where(query).First(&user)
	if user.Username != ""{
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Successful !",
			"username": user.Username,
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or password is incorrect !",
		})
	}
}