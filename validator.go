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
	//	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("isDate", isDate)
	//	validate.RegisterValidation("isPrice", isPrice)
	validate.RegisterValidation("isntSDIPec", isntSDIPec)
	validate.RegisterValidation("isDateTime", isDateTime)
	//validate.RegisterValidation("isIva", isIva)
	validate.RegisterValidation("isMP", isMP)
	validate.RegisterValidation("isRF", isRF)
	validate.RegisterValidation("isTD", isTD)
	validate.RegisterValidation("isNatura", isN)
	validate.RegisterStructValidation(datiTrasmissioneValidate, datiTrasmissione{})
	validate.RegisterStructValidation(cessionarioCommittenteValidate, CessionarioCommittente{})
	validate.RegisterStructValidation(validateDatiBeniServizi, DatiBeniServizi{})
	return validate
}

func makeValidatorWithMap(mapped map[string]string) func(validator.FieldLevel) bool {
	return func(in validator.FieldLevel) bool {
		data := in.Field().String()
		for key, _ := range mapped {
			if data == key {
				return false
			}
		}
		return true
	}

}

var isTC = makeValidatorWithMap(TipoCassa)
var isRF = makeValidatorWithMap(RegimeFiscale)
var isN = makeValidatorWithMap(Natura)
var isMP = makeValidatorWithMap(MetodiPagamento)
var isTCP = makeValidatorWithMap(TipoCessionePrestazione)
var isTD = makeValidatorWithMap(TipoDocumento)

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
	data := d.Current().Interface().(DatiGeneraliDocumento)
	if data.DatiCassaPrevidenziale.Ritenuta == "SI" {
		if *data.DatiRitenuta == (DatiRitenuta{}) {
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
	data := field.Field().String()
	_, err := time.Parse(`2006-01-02T15:04:05`, data)
	if err != nil {
		return false
	}
	return true
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
