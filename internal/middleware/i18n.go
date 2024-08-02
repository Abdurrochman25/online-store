package middleware

import (
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/BurntSushi/toml"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func LocalizeMiddleware() fiber.Handler {
	return fiberi18n.New(&fiberi18n.Config{
		RootPath:         config.Path("./message"),
		AcceptLanguages:  []language.Tag{language.English, language.Indonesian},
		DefaultLanguage:  language.Indonesian,
		FormatBundleFile: "toml",
		UnmarshalFunc:    toml.Unmarshal,
	})
}
