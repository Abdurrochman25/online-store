package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenDetail struct {
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
