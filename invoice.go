package invoice

import (
	"fmt"
	body "github.com/SpinaNico/go-struct-invoice/body"
	header "github.com/SpinaNico/go-struct-invoice/header"
)

// FatturaElettronica Fattura elettronica
type FatturaElettronica struct {
	FatturaElettronicaHeader header.FatturaElettronicaHeader `xml:"FatturaElettronicaHeader" json:"header"`
	FatturaElettronicaBody   []body.FatturaElettronicaBody   `xml:"FatturaElettronicaBody" json:"body"`
	// ds:Signature todo: firma digitale da implementare
}

// Version : print in Std out reference versione invoice.
func (f FatturaElettronica) Version() { println("invoice version 1.2") }

// Validate :
func (f FatturaElettronica) Validate() error {
	var err error
	for index, i := range f.FatturaElettronicaBody {

		if err = i.Validate(); err != nil {
			return fmt.Errorf("FatturaElettronica body:(%d) %s", index, err)
		}
	}

	if err = f.FatturaElettronicaHeader.Validate(); err != nil {
		return fmt.Errorf("FatturaElettronica %s", err)
	}
	return nil
}
