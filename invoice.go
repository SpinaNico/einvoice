package invoice

import (
	body "invoice/body"
	header "invoice/header"
)

// FatturaElettronica Fattura elettronica
type FatturaElettronica struct {
	FatturaElettronicaHeader header.FatturaElettronicaHeader `xml:"FatturaElettronicaHeader" json:"header"`
	FatturaElettronicaBody   []body.FatturaElettronicaBody   `xml:"FatturaElettronicaBody" json:"body"`
}

// Version : print in Std out reference versione invoice.
func (f FatturaElettronica) Version() { println("invoice version 1.2") }

// Validate :
func (f FatturaElettronica) Validate() error {
	for _, i := range f.FatturaElettronicaBody {
		err := i.Validate()
		if err != nil {
			return err
		}
	}
	err := f.FatturaElettronicaHeader.Validate()
	if err != nil {
		return err
	}
	return nil
}
