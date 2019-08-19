package body

import (
	"fmt"
)

//FatturaElettronicaBody corpo della fattura
type FatturaElettronicaBody struct {
	DatiGenerali    datiGenerali    `xml:"DatiGenerali" json:"DatiGenerali"`
	DatiBeniServizi datiBeniServizi `xml:"DatiBeniServizi" json:"DatiBeniServizi"`
	DatiVeicolo     datiVeicolo     `xml:"DatiVeicolo" json:"DatiVeicolo"`
	DatiPagamento   datiPagamento   `xml:"DatiPagamento" json:"DatiPagamento"`
	Allegati        allegati        `xml:"Allegati" json:"Allegati"`
}

// Validate ...
func (f FatturaElettronicaBody) Validate() error {
	var err error

	if err = f.DatiGenerali.Validate(); err != nil {
		return fmt.Errorf("DatiGenerali %s", err)
	}

	if err = f.DatiBeniServizi.Validate(); err != nil {
		return fmt.Errorf("DatiBeniServizi %s", err)
	}

	if err = f.DatiVeicolo.Validate(); err != nil {
		return fmt.Errorf("DatiVeicolo %s", err)
	}

	if err = f.DatiPagamento.Validate(); err != nil {
		return fmt.Errorf("DatiPagamento %s", err)
	}

	if err = f.Allegati.Validate(); err != nil {
		return fmt.Errorf("Allegati %s", err)
	}
	return nil
}
