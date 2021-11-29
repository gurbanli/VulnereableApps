package dto

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin *bool `json:"is_admin" binding:"required"`
}