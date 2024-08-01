package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/pkg/constants"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HTTPHandler struct{}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

type HTTPResponse struct {
	Code      string              `json:"code"`
	Message   string              `json:"message"`
	Data      interface{}         `json:"data"`
	Errors    []HTTPErrorResponse `json:"errors"`
	Metadata  *Metadata           `json:"metadata,omitempty"`
	RequestID string              `json:"request_id"`
}

type HTTPErrorResponse struct {
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}

type Metadata struct {
	Pagination struct {
		Limit    int  `json:"limit"`
		NextPage bool `json:"next_page"`
		PrevPage bool `json:"prev_page"`
		Page     int  `json:"page"`
		Total    int  `json:"total"`
	} `json:"pagination"`
}

func (h *HTTPHandler) Response(c *fiber.Ctx, statusCode int, message string) error {
	log.Info(message)
	return c.Status(http.StatusOK).JSON(HTTPResponse{
		Code:      "OLS-" + GetErrorCode(statusCode),
		Message:   fiberi18n.MustLocalize(c, fmt.Sprintf("%d", statusCode)),
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) OK(c *fiber.Ctx, action string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(HTTPResponse{
		Code:      "OLS-" + OK,
		Message:   fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d_%s", http.StatusOK, action)),
		Data:      data,
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) Unauthorized(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(HTTPResponse{
		Code:      "OLS-" + Unauthorized,
		Message:   fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d", http.StatusUnauthorized)),
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) Forbidden(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(HTTPResponse{
		Code:      "OLS-" + Forbidden,
		Message:   fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d", http.StatusForbidden)),
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) BadRequest(c *fiber.Ctx, err error) error {
	var errValidations []HTTPErrorResponse
	var errorValidator validator.ValidationErrors
	message := fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d", http.StatusBadRequest))

	validate := c.Locals(constants.CtxKeyValidator).(*config.CustomValidator)
	switch {
	case errors.As(err, &errorValidator):
		for _, err := range err.(validator.ValidationErrors) {
			errValidations = append(errValidations, HTTPErrorResponse{
				Parameter: err.Field(),
				Message:   err.Translate(validate.DefaultTrans),
			})
		}
	default:
		localizeMessage, errLocalize := fiberi18n.Localize(c, GetLocalizeTag(err))
		if errLocalize != nil {
			message = err.Error()
		} else {
			message = localizeMessage
		}
	}

	return c.Status(http.StatusBadRequest).JSON(HTTPResponse{
		Code:      "OLS-" + BadRequest,
		Message:   message,
		Errors:    errValidations,
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) NotFound(c *fiber.Ctx, localizeMessage ...string) error {
	message := fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d", http.StatusNotFound))
	if len(localizeMessage) > 0 {
		message = fiberi18n.MustLocalize(c, localizeMessage[0])
	}
	return c.Status(http.StatusNotFound).JSON(HTTPResponse{
		Code:      "OLS-" + NotFound,
		Message:   message,
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) InternalServerError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(HTTPResponse{
		Code:      "OLS-" + InternalServerError,
		Message:   fiberi18n.MustLocalize(c, fmt.Sprintf("http.%d", http.StatusInternalServerError)),
		RequestID: GetRequestID(c),
	})
}
