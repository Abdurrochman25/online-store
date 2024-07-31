package repository

import (
	"context"

	"github.com/Abdurrochman25/online-store/internal/models"
	userModels "github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IUserRepository interface {
	FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error)
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error) {
	return models.Users(Select(models.UserColumns.ID, models.UserColumns.Name, models.UserColumns.Email),
		Where("users.deleted_at IS NULL"),
		Limit(userRequest.Limit),
		Offset((userRequest.Page-1)*userRequest.Limit),
	).All(ctx, boil.GetContextDB())
}
