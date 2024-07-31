package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func DefaultHeaderAcceptLanguage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Get(fiber.HeaderAcceptLanguage, language.Indonesian.String())
		return c.Next()
	}
}
