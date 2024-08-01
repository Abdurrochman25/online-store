package usecase

import (
	"context"

	"github.com/Abdurrochman25/online-store/internal/models"
	userModels "github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/internal/user/repository"
	"github.com/Abdurrochman25/online-store/internal/util"
	"github.com/volatiletech/null/v8"
)

type IUserUseCase interface {
	FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error)
	Create(ctx context.Context, userRequest *userModels.CreateUserRequest) (*models.User, error)
}

type userUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) IUserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (uc *userUseCase) FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error) {
	return uc.userRepository.FindAll(ctx, userRequest)
}

func (uc *userUseCase) Create(ctx context.Context, userRequest *userModels.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := util.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:     null.StringFrom(userRequest.Name),
		Email:    null.StringFrom(userRequest.Email),
		Password: null.StringFrom(hashedPassword),
		RoleID:   null.NewInt16(int16(userRequest.RoleID), userRequest.RoleID != 0),
	}
	return uc.userRepository.Create(ctx, user)
}
