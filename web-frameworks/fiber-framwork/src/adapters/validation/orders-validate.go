package validation

import "github.com/go-playground/validator/v10"

func OrderAgeValidation(customValidator *Validation) error {
	return customValidator.Validator.RegisterValidation("oldAge", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() > 18 && fl.Field().Int() < 90
	})
}
