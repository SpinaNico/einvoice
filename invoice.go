package invoice

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

// FatturaElettronica Fattura elettronica
type FatturaElettronica struct {
	FatturaElettronicaHeader FatturaElettronicaHeader `xml:"FatturaElettronicaHeader" json:"FatturaElettronicaHeader"`
	FatturaElettronicaBody   []FatturaElettronicaBody `xml:"FatturaElettronicaBody" json:"FatturaElettronicaBody"`
	// ds:Signature todo: firma digitale da implementare
}

// Version : print in Std out reference versione invoice.
func (f FatturaElettronica) Version() { println("invoice version 1.2") }

// Validate :
func (f FatturaElettronica) Validate() error {
	var validate *validator.Validate
	validate = validator.New()
	validate.RegisterValidation("regimeValidate", regimeFiscaleValidator)
	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("idData", isDate)
	validate.RegisterStructValidation(datiTrasmissioneValidate, datiTrasmissione{})
	validate.RegisterStructValidation(cessionarioCommittenteValidate, cessionarioCommittente{})
	//validate.RegisterStructValidation(anagraficaValidate, anagrafica{})
	if err := validate.Struct(f); err != nil {
		return fmt.Errorf("FatturaElettronica %s", err)
	}

	// var err error
	// for index, i := range f.FatturaElettronicaBody {

	// 	if err = i.Validate(); err != nil {
	// 		return fmt.Errorf("FatturaElettronica body:(%d) %s", index+1, err)
	// 	}
	// }

	// if err = f.FatturaElettronicaHeader.Validate(); err != nil {
	// 	return fmt.Errorf("FatturaElettronica %s", err)
	// }
	return nil
}
