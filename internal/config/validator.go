package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validate     *validator.Validate
	UniTrans     *ut.UniversalTranslator
	DefaultTrans ut.Translator
}
