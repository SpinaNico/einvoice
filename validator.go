package invoice

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

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

func isPrice(d validator.FieldLevel) bool {
	return false
}
