package einvoice

import (
	"log"
	"regexp"
)

func SDIValidation(filename string, f *FatturaElettronica) error {
	var errors = SDIErrors{Errors: make([]*SDIErrorSingle, 0)}

	if filename == "" {
		errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["00001"], Code: "00001"})
	} else {
		// iso test
		if ok, err := regexp.Match("^[A-Z]{2}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["00001"], Code: "00001"})
			errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["A0001"], Code: "A0001"})
		}
		if ok, err := regexp.Match("^[A-Z]{2}[a-zA-Z0-9]{11,16}_[a-zA-Z0-9]{5}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["00001"], Code: "00001"})
			errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["A0002"], Code: "A0002"})
		}

	}

	if len(f.FatturaElettronicaBody) == 0 {
		errors.Errors = append(errors.Errors, &SDIErrorSingle{Error: ErrorsMap["A0003"], Code: "A0003"})
	}

	invocieValidator := Validator()

	if err := invocieValidator.Struct(f.FatturaElettronicaHeader); err != nil {
		errors.addErrors(err)
	}
	for _, value := range f.FatturaElettronicaBody {
		if err := invocieValidator.Struct(value); err != nil {
			errors.addErrors(err)
		}
	}

	if len(errors.Errors) == 0 {
		return nil
	}
	return &errors
}
