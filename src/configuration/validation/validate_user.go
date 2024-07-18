package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/kevinfjq/crud-golang/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		eng := en.New()
		un := ut.New(eng, eng)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		var erroCauses []rest_err.Causes
		for _, e := range validationError.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			erroCauses = append(erroCauses, cause)
		}

		return rest_err.NewBadRequestErrorValidationError("Some fields are invalid", erroCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
