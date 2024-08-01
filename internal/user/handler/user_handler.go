package handler

import (
	"errors"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/middleware"
	"github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/internal/user/repository"
	"github.com/Abdurrochman25/online-store/internal/user/usecase"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/Abdurrochman25/online-store/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

type userHTTPHandler struct {
	common.HTTPHandler
	userUseCase usecase.IUserUseCase
}

func NewUserHandler(s *config.Server) []fiber.Router {
	userRepository := repository.NewUserRepository(s.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	handler := userHTTPHandler{
		userUseCase: userUseCase,
	}

	return []fiber.Router{
		s.Fiber.Get("/v1/users", s.Authorization.RequiresRoles([]string{"admin"}), handler.findAllUser),
		s.Fiber.Post("/v1/users", s.Authorization.RequiresRoles([]string{"admin"}), handler.createUser),
	}
}

func (h *userHTTPHandler) findAllUser(c *fiber.Ctx) error {
	userRequest := new(models.GetUserRequest)
	if err := c.QueryParser(userRequest); err != nil {
		return h.BadRequest(c, err)
	}

	users, err := h.userUseCase.FindAll(c.Context(), userRequest)
	if err != nil {
		return h.InternalServerError(c)
	}

	return h.OK(c, constants.ActionRead, users)
}

func (h *userHTTPHandler) createUser(c *fiber.Ctx) error {
	userRequest := new(models.CreateUserRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return h.BadRequest(c, err)
	}
	if err := middleware.ValidateStruct(c, userRequest); err != nil {
		return h.BadRequest(c, err)
	}

	users, err := h.userUseCase.Create(c.Context(), userRequest)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrEmailAlreadyUsed):
			return h.BadRequest(c, err)
		default:
			return h.InternalServerError(c)
		}
	}

	return h.OK(c, constants.ActionRead, users)
}
