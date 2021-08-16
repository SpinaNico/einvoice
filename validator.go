package einvoice

import (
	"log"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

func invoiceValidator() *validator.Validate {
	var validate *validator.Validate
	validate = validator.New()
	//	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("isDate", isDate)
	//	validate.RegisterValidation("isPrice", isPrice)
	validate.RegisterValidation("isntSDIPec", isntSDIPec)
	validate.RegisterValidation("isDateTime", isDateTime)
	validate.RegisterValidation("isIva", isIva)
	validate.RegisterValidation("isPIVA", isPIVA)
	validate.RegisterValidation("isCF", isCF)
	validate.RegisterValidation("isMP", isMP)
	validate.RegisterValidation("isRF", isRF)
	validate.RegisterValidation("isTC", isTC)
	validate.RegisterValidation("isTD", isTD)
	validate.RegisterValidation("isNatura", isN)
	validate.RegisterValidation("noFuture", noFuture)

	registerAllValidatorStructLevel(validate)
	return validate
}

func makeValidatorWithMap(mapped map[string]string) func(validator.FieldLevel) bool {
	return func(in validator.FieldLevel) bool {
		data := in.Field().String()
		for key := range mapped {
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
var isTR = makeValidatorWithMap(TipoRitenuta)
var isCP = makeValidatorWithMap(CondizioniPagamento)

// control id Field is Integer
func isInteger(t validator.FieldLevel) bool {
	_, err := strconv.Atoi(t.Field().String())
	if err != nil {
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

func noFuture(field validator.FieldLevel) bool {

	data := field.Field().String()
	t, err := time.Parse(`2006-01-02`, data)
	if err != nil {
		return false
	}
	if t.Unix() > time.Now().Unix() {
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
	value := field.Field().Float()
	// controlli
	if value > 100 {
		return false
	}
	if value < 0 {
		return false
	}

	return true
}

//controllo se la partita iva è corretta questo controllo se fallisce genera l'errore 00301
func isPIVA(t validator.FieldLevel) bool {
	data := t.Field().String()
	ok, err := CheckPIVA(data)
	if err != nil {
		log.Println(err)
	}
	return ok
}

//controlla se è un codice fiscale italiano valido oppure PIVa
func isCF(t validator.FieldLevel) bool {
	data := t.Field().String()
	if ok, err := CheckCodiceFiscale(data); ok == true {
		log.Println(err)
		return true
	}
	ok, err := CheckPIVA(data)
	if err != nil {
		log.Println(err)
	}
	return ok

}
