package models

type GetUserRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type BodyUserRequest struct {
	Name     string `json:"name" validate:"required,max=50" example:"User Test"`
	Email    string `json:"email" validate:"required,email,max=50" example:"user.test@mail.com"`
	Password string `json:"password" validate:"required,max=255" example:"test_password"`
	RoleID   int    `json:"role_id" validate:"omitempty" example:"1"`
}
type CreateUserRequest struct {
	BodyUserRequest
}

type UpdateUserRequest struct {
	ID int `param:"id" validate:"required"`
	BodyUserRequest
}

type DeleteUserRequest struct {
	ID int `param:"id" validate:"required"`
}
