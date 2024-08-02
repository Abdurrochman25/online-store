package repository

import (
	"context"
	"database/sql"

	"github.com/Abdurrochman25/online-store/internal/models"
	userModels "github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IUserRepository interface {
	FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, user *models.User) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context, userRequest *userModels.GetUserRequest) (models.UserSlice, error) {
	return models.Users(Select(models.UserColumns.ID, models.UserColumns.Name, models.UserColumns.Email),
		Where("users.deleted_at IS NULL"),
		Limit(userRequest.Limit),
		Offset((userRequest.Page-1)*userRequest.Limit),
	).All(ctx, boil.GetContextDB())
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	isEmailUsed, err := models.Users(models.UserWhere.Email.EQ(user.Email)).Exists(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	if isEmailUsed {
		return nil, common.ErrEmailAlreadyUsed
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err := user.Insert(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, err
	}

	if user.RoleID.Valid {
		if err := user.SetRole(ctx, boil.GetContextDB(), false, &models.Role{ID: int(user.RoleID.Int16)}); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	isEmailUsed, err := models.Users(models.UserWhere.Email.EQ(user.Email), models.UserWhere.ID.NEQ(user.ID)).
		Exists(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	if isEmailUsed {
		return nil, common.ErrEmailAlreadyUsed
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := user.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}

	if err := user.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	if user.RoleID.Valid {
		if err := user.SetRole(ctx, boil.GetContextDB(), false, &models.Role{ID: int(user.RoleID.Int16)}); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := user.Delete(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}

	if err := user.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
