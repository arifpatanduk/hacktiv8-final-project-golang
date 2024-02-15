package utils

import (
	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)


var Translator ut.Translator

func InitTranslator(validate *validator.Validate) {
	enLocale := en_US.New()
	uni := ut.New(enLocale, enLocale)
	Translator, _ = uni.GetTranslator("en")
	en.RegisterDefaultTranslations(validate, Translator)
}

func TranslateValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		errors[field] = err.Translate(Translator)
	}
	return errors
}

func ValidateStruct(value interface{}) (err error) {
	validate := validator.New()
	InitTranslator(validate)

	result := validate.Struct(value)
	if result != nil {
		return result
	}

	return nil
}