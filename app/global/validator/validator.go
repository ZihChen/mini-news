package validator

import "github.com/go-playground/validator/v10"

func ValidateParams(params interface{}) error {
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return err
	}
	return nil
}
