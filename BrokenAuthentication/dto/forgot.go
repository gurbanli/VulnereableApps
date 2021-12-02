package dto


type ForgotPasswordSendRequest struct {
	Username string `json:"username" binding:"required"`
}

type ForgotPasswordVerifyRequest struct {
	Username string `json:"username" binding:"required"`
	ResetToken string `json:"reset_token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}