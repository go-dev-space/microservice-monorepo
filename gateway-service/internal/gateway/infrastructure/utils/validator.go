package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Instance *validator.Validate
}

func NewValidator() *Validation {
	return &Validation{
		Instance: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v Validation) Struct(data any) error {
	return v.Instance.Struct(data)
}

func (v Validation) Test(e error) (map[string]string, error) {
	out := make(map[string]string)
	if hits, ok := e.(validator.ValidationErrors); ok {
		for _, hit := range hits {
			switch hit.Tag() {
			case "required":
				out[hit.Field()] = fmt.Sprintf("%s is required", hit.Field())
			}
		}
		return out, errors.New("throw")
	}
	return out, nil
}
