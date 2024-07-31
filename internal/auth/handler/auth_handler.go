package handler

import (
	"github.com/Abdurrochman25/online-store/internal/auth/models"
	"github.com/Abdurrochman25/online-store/internal/auth/repository"
	"github.com/Abdurrochman25/online-store/internal/auth/usecase"
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	common.HTTPHandler
	authUseCase usecase.IAuthUseCase
}

func NewAuthHandler(s *config.Server) fiber.Router {
	authRepository := repository.NewAuthRepository()
	authUseCase := usecase.NewAuthUseCase(authRepository, &s.Config)

	handler := authHandler{
		authUseCase: authUseCase,
	}

	return s.Fiber.Post("/v1/auth/login", handler.Login)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return h.BadRequest(c)
	}

	token, err := h.authUseCase.Login(c.Context(), loginRequest)
	if err != nil {
		return h.InternalServerError(c)
	}

	return h.OK(c, common.ActionRead, token)
}
