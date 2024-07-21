package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Routes []fiber.Router
}
