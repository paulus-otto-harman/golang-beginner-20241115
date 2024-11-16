package lib

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param"`
}

func Validate(input interface{}) interface{} {
	err := validator.New().Struct(input)

	if err != nil {
		var errorBags []Error
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, inputError := range validationErrors {
				errorBags = append(errorBags, Error{
					Field: inputError.Field(),
					Tag:   inputError.Tag(),
					Value: inputError.Value(),
					Param: inputError.Param(),
				})
			}
			return errorBags
		}
		return errors.New(fmt.Sprintf("validation error %v", err))
	}
	return nil
}
