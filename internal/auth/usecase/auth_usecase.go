package usecase

import (
	"context"
	"time"

	authModels "github.com/Abdurrochman25/online-store/internal/auth/models"
	"github.com/Abdurrochman25/online-store/internal/auth/repository"
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/middleware"
	"github.com/Abdurrochman25/online-store/internal/util"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type IAuthUseCase interface {
	Login(ctx context.Context, loginRequest *authModels.LoginRequest) (*authModels.TokenDetail, error)
}

type authUseCase struct {
	authRepository repository.IAuthRepository
	authConfig     *config.Config
}

func NewAuthUseCase(authRepository repository.IAuthRepository, authConfig *config.Config) IAuthUseCase {
	return &authUseCase{authRepository: authRepository, authConfig: authConfig}
}

func (uc *authUseCase) Login(ctx context.Context, loginRequest *authModels.LoginRequest) (*authModels.TokenDetail, error) {
	user, err := uc.authRepository.FindByEmail(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	if passwordMatched := util.CheckPasswordHash(user.Password.String, loginRequest.Password); !passwordMatched {
		log.Warn("Password doesn't matched")
		return nil, common.ErrPasswordDoesntMatched
	}

	currentTime := time.Now().Local()
	expiresAt := jwt.NewNumericDate(util.EndOfDay(currentTime))

	claims := middleware.JWTClaims{
		UserName: user.Name.String,
		Role:     user.R.Role.Name.String,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(currentTime),
			NotBefore: jwt.NewNumericDate(currentTime),
			Subject:   user.Email.String,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "v1"

	tokenString, err := token.SignedString([]byte(uc.authConfig.AppSecret))
	if err != nil {
		log.Warn("Failed get access token")
		return nil, err
	}

	tokenDetail := &authModels.TokenDetail{
		UserName:    user.Name.String,
		Email:       user.Email.String,
		Role:        user.R.Role.Name.String,
		ExpiresAt:   expiresAt.Unix(),
		AccessToken: tokenString,
	}

	return tokenDetail, nil
}
