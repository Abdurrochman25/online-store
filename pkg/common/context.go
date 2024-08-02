package common

import (
	"github.com/Abdurrochman25/online-store/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

func GetRequestID(c *fiber.Ctx) string {
	if requestID, ok := c.Locals(constants.CtxKeyRequestID).(string); ok {
		return requestID
	}

	return ""
}
