package einvoice

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type SDIError struct {
	Errors map[string]string
}

func (e *SDIError) Error() string {

	var strs []string
	for k, v := range e.Errors {
		strs = append(strs, fmt.Sprintf("sdiError-%s: %s", k, v))
	}

	return strings.Join(strs, "\n")
}

func SDIValidation(f *FatturaElettronica) error {

	var errors SDIError
	err := f.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "isntSDIPec":
				errors.Errors["00330"] = ErrorsMap["00330"]
				break
			}
		}
	}
	return &errors
}
