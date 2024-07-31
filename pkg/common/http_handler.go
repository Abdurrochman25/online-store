package common

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	return c.Status(http.StatusOK).JSON(HTTPResponse{
		Code:      "OLS-" + GetErrorCode(statusCode),
		Message:   message,
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) OK(c *fiber.Ctx, data interface{}) error {
	return c.Status(http.StatusOK).JSON(HTTPResponse{
		Code:      "OLS-" + OK,
		Message:   "Operation successfully executed",
		Data:      data,
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) Unauthorized(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(HTTPResponse{
		Code:      "OLS-" + Unauthorized,
		Message:   "Access denied due to invalid credential",
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) Forbidden(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(HTTPResponse{
		Code:      "OLS-" + Forbidden,
		Message:   "You don't have permission to access this resource",
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) BadRequest(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(HTTPResponse{
		Code:      "OLS-" + BadRequest,
		Message:   "Invalid request",
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) NotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(HTTPResponse{
		Code:      "OLS-" + NotFound,
		Message:   "Resource not found",
		RequestID: GetRequestID(c),
	})
}

func (h *HTTPHandler) InternalServerError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(HTTPResponse{
		Code:      "OLS-" + InternalServerError,
		Message:   "There's something wrong, please contact the administrator",
		RequestID: GetRequestID(c),
	})
}
