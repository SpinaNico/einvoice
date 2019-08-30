package invoice

import (
	"fmt"
	"regexp"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

func getValidator() *validator.Validate {
	var validate *validator.Validate
	validate = validator.New()
	validate.RegisterValidation("regimeValidate", regimeFiscaleValidator)
	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("isDate", isDate)
	validate.RegisterValidation("isPrice", isPrice)
	validate.RegisterValidation("isNatura", isNatura)
	validate.RegisterValidation("isDateTime", isDateTime)
	validate.RegisterStructValidation(datiTrasmissioneValidate, datiTrasmissione{})
	validate.RegisterStructValidation(cessionarioCommittenteValidate, cessionarioCommittente{})
	return validate
}

func regimeFiscaleValidator(rf validator.FieldLevel) bool {
	RF := rf.Field().String()

	regimeFiscale := make(map[string]string)
	for i := 1; i < 20; i++ {
		regimeFiscale[fmt.Sprintf("RF%d", i)] = "true"
	}
	_, exists := regimeFiscale[string(RF)]
	//println(exists, RF)
	return exists == true
}

// control id Field is Integer
func isInteger(t validator.FieldLevel) bool {
	_, err := strconv.Atoi(t.Field().String())
	if err != nil {
		return false
	}
	return true
}

// Struct Level control
// This control at the structure level is essential, check that if the
// target code is "0000000" the pec is not empty
func datiTrasmissioneValidate(d validator.StructLevel) {
	data := d.Current().Interface().(datiTrasmissione)
	if data.CodiceDestinatario == "0000000" {
		if data.PECDestinatario == "" {
			d.ReportError(data.PECDestinatario, "PECDestinatario", "", "required", "")
		}
	}

}

// if the person making the invoice is a foreigner, the Italian office must be indicated.
// This validator also checks whether the Stabile Organization is in the
// Italian territory so the Nation value is "IT"
func cessionarioCommittenteValidate(d validator.StructLevel) {
	data := d.Current().Interface().(cessionarioCommittente)
	if data.Sede.Nazione != "IT" {
		if data.StabileOrganizzazione == (indirizzoType{}) {
			d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "required", "")
		} else {
			if data.StabileOrganizzazione.Nazione != "IT" {
				d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "eq", "")
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
	matched, _ := regexp.Match(`\d\d\d-\d\d-\d\d`, []byte(data))
	if !matched {
		return false
	}
	if len(data) != 10 {
		return false
	}

	return true
}

// validate format: YYYY-MM- DDTHH:MM:SS.
func isDateTime(field validator.FieldLevel) bool {
	return true
}

func isNatura(field validator.FieldLevel) bool {
	c := field.Field().String()
	if len(string(c)) != 2 {
		return false // return fmt. f("Natura %s", share. Egual(2))
	}
	values := make(map[string]string)
	for i := 1; i < 8; i++ {
		values[fmt.Sprintf("N%d", i)] = "true"
	}
	if _, ok := values[string(c)]; ok == false {
		return false
	}
	return true
}
