package einvoice

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

func Validator() *validator.Validate {
	var validate *validator.Validate
	validate = validator.New()
	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("isDate", isDate)
	validate.RegisterValidation("isPrice", isPrice)
	validate.RegisterValidation("isTypeDocument", isTypeDocument)
	validate.RegisterValidation("isntSDIPec", isntSDIPec)
	validate.RegisterValidation("isNatura", isNatura)
	validate.RegisterValidation("isDateTime", isDateTime)
	validate.RegisterValidation("isIva", isIva)
	validate.RegisterValidation("isMP", isMP)
	validate.RegisterValidation("isRF", isRF)
	validate.RegisterStructValidation(datiTrasmissioneValidate, datiTrasmissione{})
	validate.RegisterStructValidation(cessionarioCommittenteValidate, CessionarioCommittente{})
	validate.RegisterStructValidation(validateDatiBeniServizi, datiBeniServizi{})
	return validate
}

func isRF(rf validator.FieldLevel) bool {
	RF := rf.Field().String()

	for key, _ := range RegimeFiscale {
		if RF == key {
			return false
		}
	}
	return true
}

// control id Field is Integer
func isInteger(t validator.FieldLevel) bool {
	_, err := strconv.Atoi(t.Field().String())
	if err != nil {
		return false
	}
	return true
}

//Check the data hold block, which must be filled in case of DatiCassaPrevidenziale.Ritenuta == "SI"
func checkDatiRitenuta(d validator.StructLevel) {
	data := d.Current().Interface().(datiGeneraliDocumento)
	if data.DatiCassaPrevidenziale.Ritenuta == "SI" {
		if *data.DatiRitenuta == (datiRitenuta{}) {
			d.ReportError(data.DatiRitenuta, "struct", "all", "required", "")
		} else {
			validate := Validator()
			if err := validate.Struct(data.DatiRitenuta); err != nil {
				d.ReportError("DatiRitenuta", "", fmt.Sprintf("%s", err), "", "")
			}

		}
	}

}

// control is is a price number 0.00
func isPrice(d validator.FieldLevel) bool {
	s := d.Field().String()
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return false
	}
	stringVersion := fmt.Sprintf("%.2f", val)
	if s != stringVersion {
		return false
	}
	return true
}

//is Data format: YYYY-MM-DD.
func isDate(field validator.FieldLevel) bool {
	data := field.Field().String()
	_, err := time.Parse(`2006-01-02`, data)
	if err != nil {
		return false
	}
	return true
}

// validate format: YYYY-MM-DDTHH:MM:SS.
func isDateTime(field validator.FieldLevel) bool {

	return true
}

func isNatura(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range NatureWithDescription {
		if key == c {
			return true
		}
	}
	return false
}

func isMP(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range MethodsPayments {
		if key == c {
			return true
		}
	}
	return false
}

func isTCP(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range EnumTipoCessionePrestazione {
		if key == c {
			return true
		}
	}
	return false
}

func isTypeDocument(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range TypeDocument {
		if key == c {
			return true
		}
	}
	return false
}

func isntSDIPec(field validator.FieldLevel) bool {
	c := field.Field().String()
	matched, _ := regexp.MatchString(`sdi\d\d@pec\.fatturapa\.it`, c)
	return matched
	//sdixx@pec.fatturapa.it
}

func isIva(field validator.FieldLevel) bool {
	c := field.Field().String()
	_, err := strconv.ParseFloat(c, 32)
	if err != nil {
		log.Println("ERROR: ", err)
		return true
	}

	return false
}
