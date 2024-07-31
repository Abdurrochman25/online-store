package common

import "github.com/gofiber/fiber/v2"

func GetRequestID(c *fiber.Ctx) string {
	if requestID, ok := c.Locals(CtxKeyRequestID).(string); ok {
		return requestID
	}

	return ""
}
