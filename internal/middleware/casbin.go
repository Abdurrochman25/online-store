package middleware

import (
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/pkg/common"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	casbinFiber "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type AuthzConfig struct {
	common.HTTPHandler
}

func NewAuthzConfig() *AuthzConfig {
	return &AuthzConfig{}
}

func (ac *AuthzConfig) CustomAuthorization() *casbinFiber.Middleware {
	return casbinFiber.New(casbinFiber.Config{
		ModelFilePath: config.Path("model.conf"),
		PolicyAdapter: fileadapter.NewAdapter(config.Path("policy.csv")),
		Lookup:        func(c *fiber.Ctx) string { return c.Context().UserValue("name").(string) },
		Unauthorized:  ac.Unauthorized,
		Forbidden:     ac.Forbidden,
	})
}
