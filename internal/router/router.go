package router

import (
	authHandler "github.com/Abdurrochman25/online-store/internal/auth/handler"
	"github.com/Abdurrochman25/online-store/internal/config"
	userHandler "github.com/Abdurrochman25/online-store/internal/user/handler"
	"github.com/gofiber/fiber/v2"
)

func AttachAllRoutes(s *config.Server) {
	s.Router.Routes = []fiber.Router{
		authHandler.NewAuthHandler(s),
		userHandler.NewUserHandler(s),
	}
}
