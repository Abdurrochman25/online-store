package middleware

import (
	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/pkg/constants"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func ValidatorMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		validate := validator.New()
		en := en.New()
		id := id.New()
		uni := ut.New(en, en, id)

		idTrans, _ := uni.GetTranslator(language.Indonesian.String())
		id_translations.RegisterDefaultTranslations(validate, idTrans)

		customValidator := &config.CustomValidator{
			Validate:     validate,
			UniTrans:     uni,
			DefaultTrans: idTrans,
		}

		c.Locals(constants.CtxKeyValidator, customValidator)
		return c.Next()
	}
}

func ValidateStruct(c *fiber.Ctx, s interface{}) validator.ValidationErrors {
	validate := c.Locals(constants.CtxKeyValidator).(*config.CustomValidator)
	if errs := validate.Validate.Struct(s); errs != nil {
		return errs.(validator.ValidationErrors)
	}

	return nil
}
