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

func (aS *AuthServiceImpl) RegisterUser(rRequest dto.RegisterRequest) (model.User, error){
	user := repository.Repo.CreateUser(rRequest)
	if user.Username != ""{
		return user, nil
	}else{
		return user, errors.New("can not register user")
	}
}

func (aS *AuthServiceImpl) CheckUsername(username string) error{
	user := repository.Repo.FindByUsername(username)
	if user.Username != "" {
		return nil
	}else{
		return errors.New("user not found")
	}
}

func (aS *AuthServiceImpl) CheckUsernameAndPassword(username string, password string) error{
	user := repository.Repo.FindByUsernameAndPassword(username, password)
	if user.Username != "" {
		return nil
	}else{
		return errors.New("user not found")
	}
}

func (aS *AuthServiceImpl) CheckOTP() error{
	err := new(error)
	return *err
}