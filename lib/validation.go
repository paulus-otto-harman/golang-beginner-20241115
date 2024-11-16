package lib

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Error struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
	Param string      `json:"param"`
}

func Validate(input interface{}) interface{} {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.Struct(input)

	if err != nil {
		var errorBags []Error
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, inputError := range validationErrors {
				fieldError := Error{
					Field: inputError.Field(),
					Tag:   inputError.Tag(),
					Value: inputError.Value(),
					Param: inputError.Param(),
				}

				if fieldError.Tag == "eqfield" {
					if param, validationHasParam := reflect.TypeOf(input).FieldByName(fieldError.Param); validationHasParam {
						fieldError.Param = param.Tag.Get("json")
					}
				}
				errorBags = append(errorBags, fieldError)
			}
			return errorBags
		}
		return errors.New(fmt.Sprintf("validation error %v", err))
	}
	return nil
}
