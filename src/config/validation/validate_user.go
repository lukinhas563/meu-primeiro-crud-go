package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_error error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return rest_error.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorCauses := []rest_error.Causes{}

		for _, e := range validation_error.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_error.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_error.NewBadRequestError("Error trying to convert fields")
	}
}
