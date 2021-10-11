package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

type BaseErrorValidation struct {
	Field      string
	Validation string
	Value      interface{}
}

func Validation(data interface{}) []error {
	var validationErrors []error
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// baseError := BaseErrorValidation{
			// 	Field:      err.StructNamespace(),
			// 	Validation: err.Tag(),
			// 	Value:      err.Value(),
			// }

			msg := fmt.Errorf("field %s validation %s %s", err.StructNamespace(), err.Tag(), err.Param())

			validationErrors = append(validationErrors, msg)
		}
	}
	return validationErrors
}
