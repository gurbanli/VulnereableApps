package controller

import (
	"ORMInjection4Ingress/model"
	"ORMInjection4Ingress/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//author: gurbanli

type UserController struct {

}

func (uController *UserController)GetUser(c *gin.Context){
	id := c.Param("id")
	var user model.User
	db := repository.Database
	db.First(&user, id)
	if user.Username != ""{
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}else{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found!",
		})
	}


}