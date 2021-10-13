package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type BaseErrorValidation struct {
	Field      string
	Validation string
	Value      interface{}
}

func Validation(data interface{}) []*error {
	var validationErrors []*error
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			msg := fmt.Errorf("field %s validation %s %s", strings.ToLower(strings.Split(err.StructNamespace(), ".")[1]), err.Tag(), err.Param())

			validationErrors = append(validationErrors, &msg)
		}
	}
	return validationErrors
}
