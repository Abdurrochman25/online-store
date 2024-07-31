package usecase

import (
	"context"

	"github.com/Abdurrochman25/online-store/internal/models"
	userModels "github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/internal/user/repository"
)

type IUserUseCase interface {
	FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error)
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
