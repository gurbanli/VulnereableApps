package controller

import (
	"BrokenAuthentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//author: gurbanli

type AdminController struct {

}

var Admin *AdminController

func (aC *AdminController) Index(c *gin.Context){
	if service.Session.IsAdmin(c) || c.GetHeader("IsAdmin") == "true"{
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Admin !",
		})
		return
	}
	c.JSON(http.StatusForbidden, gin.H{
		"message": "You are not admin !",
	})

}