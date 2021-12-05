package service

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/model"
	"BrokenAuthentication/repository"
	"errors"
)

type AuthService interface {
	CheckUsernameAndPassword()
	CheckOTP()
	CheckUsername()
}


type AuthServiceImpl struct {

}

var Auth *AuthServiceImpl

func (aS *AuthServiceImpl) RegisterUser(rRequest dto.RegisterRequest) (*model.User, error){
	user := repository.Repo.CreateUser(rRequest)
	return user, nil
}

func (aS *AuthServiceImpl) CheckUsernameAndPassword(username string, password string) (*model.User, error){
	user := repository.Repo.FindByUsernameAndPassword(username, password)
	if user != nil {
		return user, nil
	}else{
		return nil, errors.New("user not found")
	}
}





