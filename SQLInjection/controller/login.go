package controller

import (
	"SQLInjection/dto"
	"SQLInjection/model"
	"SQLInjection/repository"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


type LoginController struct {

}


func (lc *LoginController) LoginByUsername(c *gin.Context){
	sess := sessions.Default(c)
	var loginRequest dto.LoginRequest
	var user model.User
	if err := c.ShouldBind(&loginRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	db := repository.Database
	rows, _ := db.Query(fmt.Sprintf("SELECT id, username, is_admin FROM users WHERE username='%s' and password='%s'", loginRequest.Username, loginRequest.Password))
	if rows != nil  && rows.Next(){
		rows.Scan(&user.Id, &user.Username, &user.IsAdmin)
		sess.Set("id", user.Id)
		sess.Set("is_admin", user.IsAdmin)
		sess.Save()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome %s !", user.Username),
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or Password is incorrect !",
		})
	}
}