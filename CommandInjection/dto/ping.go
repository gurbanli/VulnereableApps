package dto


type PingRequest struct {
	Host string `json:"host" binding:"required"`
}

