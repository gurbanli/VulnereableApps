package service

import (
	"BrokenAuthentication/dto"
	"BrokenAuthentication/model"
	"BrokenAuthentication/repository"
	"BrokenAuthentication/util"
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

func (aS *AuthServiceImpl) CheckUsernameAndPassword(username string, password string) (*model.User, error){
	user := repository.Repo.FindByUsernameAndPassword(username, password)
	if user.Username != "" {
		return &user, nil
	}else{
		return nil, errors.New("user not found")
	}
}

func (aS *AuthServiceImpl) ChangePassword(username string, password string) *model.User{
	user := repository.Repo.ChangeUserPassword(username, password)
	return user
}


func (aS *AuthServiceImpl) SendOTP(user *model.User) bool{
	otpCode := util.GenerateOTP()
	if util.SendOTP(otpCode, user.Email) {
		repository.Repo.SetOTP(user.Username, otpCode)
		return true
	}
	return false
}
func (aS *AuthServiceImpl) VerifyOTP(username string, otpCode string) bool{
	return repository.Repo.GetOTP(username) == otpCode
}

func (aS *AuthServiceImpl) SendResetToken(username string){
	user := repository.Repo.FindByUsername(username)
	if user.Email != ""{
		resetToken := util.GenerateResetToken()
		if util.SendForgotPasswordToken(resetToken,user.Email){
			repository.Repo.SetResetToken(username, resetToken)
		}
	}
}

func (aS *AuthServiceImpl) VerifyResetToken(username string, resetToken string)bool{
	return repository.Repo.GetResetToken(username) == resetToken
}

