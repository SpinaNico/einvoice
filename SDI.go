package einvoice

import (
	"log"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

func SDIValidation(filename string, f *FatturaElettronica) error {
	var errors = SDIError{Errors: make(map[string]string, 0)}

	if filename == "" {
		errors.Errors["00001"] = ErrorsMap["00001"]
	} else {
		// iso test
		if ok, err := regexp.Match("^[A-Z]{2}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors["00001"] = ErrorsMap["00001"]
			errors.Errors["A0001"] = ErrorsMap["A0001"]
		}
		if ok, err := regexp.Match("^[A-Z]{2}[a-zA-Z0-9]{11,16}_DF_[a-zA-Z0-9]{5}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors["00001"] = ErrorsMap["00001"]
			errors.Errors["A0002"] = ErrorsMap["A0002"]
		}

	}

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
	if len(errors.Errors) == 0 {
		return nil
	}
	return &errors
}
