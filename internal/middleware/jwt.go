package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/pkg/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type JWTSkipper struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type JWTConfig struct {
	JWTSkippers []JWTSkipper `yaml:"jwt_skippers"`
}

type JWTClaims struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type AuthConfig struct {
	common.HTTPHandler
	config   *config.Config
	skippers []JWTSkipper
}

func NewAuthConfig(conf *config.Config) *AuthConfig {
	var jwtCfg JWTConfig
	if err := common.YamlDecodeFile(config.Path("middleware.yaml"), &jwtCfg); err != nil {
		log.Panic("Error read file middleware.yaml", err)
	}

	return &AuthConfig{
		config:   conf,
		skippers: jwtCfg.JWTSkippers,
	}
}

func (ac *AuthConfig) CustomAuthentication() fiber.Handler {
	return jwtware.New(jwtware.Config{
		Filter: ac.skipperHandler,
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS256.Name,
			Key:    []byte(ac.config.AppSecret),
		},
		// SigningKeys: map[string]jwtware.SigningKey{
		// 	"v1": {
		// 		JWTAlg: jwt.SigningMethodHS256.Name,
		// 		Key:    []byte(ac.config.AppSecret),
		// 	},
		// },
		// TokenLookup:    "header:Authorization",
		// AuthScheme:     "Bearer ",
		KeyFunc:        ac.parseToken,
		SuccessHandler: ac.successHandler,
		ErrorHandler:   ac.errorHandler,
	})
}

func (ac *AuthConfig) parseToken(t *jwt.Token) (interface{}, error) {
	var claims JWTClaims
	t, err := jwt.ParseWithClaims(t.Raw, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(ac.config.AppSecret), nil
			}
		}
		return nil, fmt.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("invalid token")
	}
	return []byte(ac.config.AppSecret), nil
}

func (ac *AuthConfig) skipperHandler(c *fiber.Ctx) bool {
	for _, skipper := range ac.skippers {
		isMatched := strings.EqualFold(string(c.Request().Header.Method()), skipper.Method) && strings.EqualFold(c.Path(), skipper.Path)
		if isMatched {
			return true
		}
	}

	return false
}

func (ac *AuthConfig) successHandler(c *fiber.Ctx) error {
	switch c.Context().UserValue("user").(type) {
	case *jwt.Token:
		if token, ok := c.Context().UserValue("user").(*jwt.Token); ok {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Context().SetUserValue("name", claims["user_name"])
				c.Context().SetUserValue("email", claims["sub"])
				c.Context().SetUserValue("role", claims["role"])
			}
		}
	}
	return c.Next()
}

func (ac *AuthConfig) errorHandler(c *fiber.Ctx, err error) error {
	log.Errorf("Error JWT Middleware, %v", err)
	switch {
	case errors.Is(err, jwtware.ErrJWTMissingOrMalformed):
		return ac.Unauthorized(c)
	case strings.Contains(err.Error(), "token is expired by"):
		return ac.Unauthorized(c)
	default:
		return ac.Response(c, http.StatusUnauthorized, err.Error())
	}
}
