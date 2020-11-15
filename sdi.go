package einvoice

import (
	"log"
	"regexp"
)

func SDIValidation(filename string, f *FatturaElettronica) *SDIErrors {
	var errors = SDIErrors{Errors: make([]*SDIError, 0)}

	if filename == "" {
		errors.Errors = append(errors.Errors, NewErrorSDI("00001"))
	} else {
		// iso test
		if ok, err := regexp.Match("^[A-Z]{2}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors = append(errors.Errors, NewErrorSDI("00001"))
		}
		if ok, err := regexp.Match("^[A-Z]{2}[a-zA-Z0-9]{11,16}_[a-zA-Z0-9]{5}", []byte(filename)); ok != true {
			if err != nil {
				log.Fatal(err)
			}
			errors.Errors = append(errors.Errors, NewErrorSDI("00001"))
		}

	}

	// if len(f.FatturaElettronicaBody) == 0 {
	// 	errors.Errors = append(errors.Errors, NewErrorSDI(" A0003"))
	// }

	val := invoiceValidator()

	if err := val.Struct(f.FatturaElettronicaHeader); err != nil {
		errors.errorsSwitch(err)
	}
	for _, value := range f.FatturaElettronicaBody {
		if err := val.Struct(value); err != nil {
			errors.errorsSwitch(err)
		}
	}

	if len(errors.Errors) == 0 {
		return nil
	}
	return &errors
}
