package repository

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/model"
	"BrokenAuthentication/util"
)

//author: gurbanli
func (r *Repository) CreateUser(rRequest dto.RegisterRequest) *model.User{
	user := model.User{
		Username: rRequest.Username,
		Password: util.Hash(rRequest.Password),
		Email: rRequest.Email,
		IsAdmin:  *rRequest.IsAdmin,
	}
	r.DB.Create(&user)
	return &user
}


func (r *Repository) FindByUsernameAndPassword(username string, password string) *model.User{
	var user model.User
	r.DB.Where("username = ? and password = ?", username, util.Hash(password)).First(&user)
	if user.Username != "" {
		return &user
	}
	return nil
}

