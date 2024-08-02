package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/user/models"
	"github.com/Abdurrochman25/online-store/internal/user/repository"
	"github.com/Abdurrochman25/online-store/internal/user/usecase"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/Abdurrochman25/online-store/pkg/constants"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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
		s.Fiber.Put("/v1/users/:id", s.Authorization.RequiresRoles([]string{"admin"}), handler.updateUser),
		s.Fiber.Delete("/v1/users/:id", s.Authorization.RequiresRoles([]string{"admin"}), handler.deleteUser),
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
	if err := h.BindAndValidate(c, userRequest); err != nil {
		return h.BadRequest(c, err)
	}

	user, err := h.userUseCase.Create(c.Context(), userRequest)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrEmailAlreadyUsed):
			return h.BadRequest(c, err)
		default:
			return h.InternalServerError(c)
		}
	}

	return h.OK(c, constants.ActionRead, user)
}

func (h *userHTTPHandler) updateUser(c *fiber.Ctx) error {
	userRequest := new(models.UpdateUserRequest)
	if err := h.BindAndValidate(c, userRequest); err != nil {
		return h.BadRequest(c, err)
	}

	user, err := h.userUseCase.Update(c.Context(), userRequest)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, common.ErrEmailAlreadyUsed):
			return h.BadRequest(c, err)
		case errors.Is(err, common.ErrRecordNotFound):
			message := fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
				MessageID: common.LocalizeRecordNotFound,
				TemplateData: map[string]string{
					"Field": fmt.Sprintf("ID = %d", userRequest.ID),
				},
			})
			return h.Response(c, http.StatusBadRequest, message)
		default:
			return h.InternalServerError(c)
		}
	}

	return h.OK(c, constants.ActionUpdate, user)
}

func (h *userHTTPHandler) deleteUser(c *fiber.Ctx) error {
	userRequest := new(models.DeleteUserRequest)
	if err := h.BindAndValidate(c, userRequest); err != nil {
		return h.BadRequest(c, err)
	}

	user, err := h.userUseCase.Delete(c.Context(), userRequest)
	if err != nil {
		log.Error(err)
		switch {
		case errors.Is(err, common.ErrRecordNotFound):
			message := fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
				MessageID: common.LocalizeRecordNotFound,
				TemplateData: map[string]string{
					"Field": fmt.Sprintf("ID = %d", userRequest.ID),
				},
			})
			return h.Response(c, http.StatusBadRequest, message)
		default:
			return h.InternalServerError(c)
		}
	}

	return h.OK(c, constants.ActionDelete, user)
}
