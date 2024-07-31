package models

type GetUserRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}
