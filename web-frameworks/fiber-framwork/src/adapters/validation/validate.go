package validation

import (
	"github.com/go-playground/validator/v10"
)

type (
	ValidationError struct {
		HasError bool
		Field    string
		Tag      string
		Value    interface{}
	}

	Validation struct {
		Validator *validator.Validate
	}
)

func (v Validation) Validate(validate *validator.Validate, data interface{}) []ValidationError {

	var validationErrors []ValidationError

	errs := validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				HasError: true,
				Field:    err.Field(),
				Tag:      err.Tag(),
				Value:    err.Value(),
			})
		}

		return validationErrors
	}

	return nil
}

func SetRoutesValidation(customValidator *Validation) error {

	err := OrderAgeValidation(customValidator)

	if err != nil {
		return err
	}

	return nil

}
