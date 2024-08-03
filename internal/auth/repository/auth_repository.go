package repository

import (
	"context"

	authModels "github.com/Abdurrochman25/online-store/internal/auth/models"
	"github.com/Abdurrochman25/online-store/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IAuthRepository interface {
	FindByEmail(ctx context.Context, loginRequest *authModels.LoginRequest) (*models.User, error)
	FindPermissionByUserID(ctx context.Context, userID int) (*models.User, error)
}

type authRepository struct{}

func NewAuthRepository() IAuthRepository {
	return &authRepository{}
}

func (r *authRepository) FindByEmail(ctx context.Context, loginRequest *authModels.LoginRequest) (*models.User, error) {
	user, err := models.Users(
		Select(
			models.UserColumns.ID, models.UserColumns.Name, models.UserColumns.Email,
			models.UserColumns.Password, models.UserColumns.RoleID,
		), models.UserWhere.Email.EQ(null.NewString(loginRequest.Email, true)),
		Load(Rels(models.UserRels.Role), Select(models.RoleColumns.ID, models.RoleColumns.Name)),
	).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *authRepository) FindPermissionByUserID(ctx context.Context, userID int) (*models.User, error) {
	return models.Users(Where("user_id=?", userID)).
		One(ctx, boil.GetContextDB())
}
