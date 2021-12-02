package dto


type OTPVerifyRequest struct {
	Username string `json:"username" binding:"required"`
	OTPCode string `json:"otp_code" binding:"required,min=4,max=4"`
}
