package controller

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/repository"
	"BrokenAuthentication/service"
	"BrokenAuthentication/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {

}

var Login *LoginController

func(lc *LoginController) LoginByPasswordEnum(c *gin.Context){
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
	user, err := service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
		"user": user,
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
	user, err := service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Login Failed !",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
		"user": user,
	})

}

func(lc *LoginController) LoginByPasswordBlockIpVulnerable(c *gin.Context){
	var lRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	user, err := service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClient(c.Request), true){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Your ip address is blocked !",
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
		"user": user,
	})

}

func(lc *LoginController) LoginByPasswordBlockIpVulnerable2(c *gin.Context){
	var lRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClient(c.Request), true){
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Your ip address is blocked !",
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
	repository.Repo.SetIPBlockStatus(util.GetIpAddressOfClient(c.Request), 0)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful !",
		"user": user,
	})

}

func(lc *LoginController) LoginByPasswordBlockIpVulnerable3(c *gin.Context){
	var lRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&lRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	user, err := service.Auth.CheckUsernameAndPassword(lRequest.Username, lRequest.Password)
	if err != nil{
		if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClientVulnerable(c.Request), true){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Your ip address is blocked !",
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}else{
		if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClientVulnerable(c.Request), false){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Your ip address is blocked !",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Successful !",
			"user": user,
		})
	}
}

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
		if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClient(c.Request), true){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Your ip address is blocked !",
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Username or password is incorrect !",
		})
		return
	}else{
		if repository.Repo.GetIPBlockStatus(util.GetIpAddressOfClientVulnerable(c.Request), false){
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Your ip address is blocked !",
			})
			return
		}

		if service.Auth.SendOTP(user){
			c.JSON(http.StatusOK, gin.H{
				"message": "OTP Code is sent to your email !",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "OTP Code can not be sent !",
		})
	}
}