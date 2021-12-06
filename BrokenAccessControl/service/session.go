package service

import (
	"BrokenAuthentication/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//author: gurbanli
type SessionService interface {
	SetAuthSession()
	IsAdmin()
}


type SessionServiceImpl struct {

}

var Session *SessionServiceImpl

func (sS *SessionServiceImpl) SetAuthSession(c *gin.Context, user *model.User){
	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Set("is_admin", user.IsAdmin)
	session.Save()
}

func (sS *SessionServiceImpl) IsAdmin(c *gin.Context)bool{
	session := sessions.Default(c)
	if session.Get("id") != nil{
		return session.Get("is_admin").(bool)
	}
	return false
}