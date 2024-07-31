package handler

import (
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/internal/user/repository"
	"github.com/Abdurrochman25/online-store/internal/user/usecase"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/gofiber/fiber/v2"
)

type userHTTPHandler struct {
	common.HTTPHandler
	userUseCase usecase.IUserUseCase
}

func NewUserHandler(s *config.Server) fiber.Router {
	userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	handler := userHTTPHandler{
		userUseCase: userUseCase,
	}

	return s.Fiber.Get("/v1/users", s.Authorization.RequiresRoles([]string{"admin"}), handler.findAllUser)
}

func (h *userHTTPHandler) findAllUser(c *fiber.Ctx) error {
	userRequest := new(models.GetUserRequest)
	if err := c.QueryParser(userRequest); err != nil {
		return h.BadRequest(c)
	}

	users, err := h.userUseCase.FindAll(c.Context(), userRequest)
	if err != nil {
		return h.InternalServerError(c)
	}

	return h.OK(c, common.ActionRead, users)
}
