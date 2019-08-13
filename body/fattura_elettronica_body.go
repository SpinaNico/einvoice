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
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}

	if err = f.DatiBeniServizi.Validate(); err != nil {
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}

	if err = f.DatiVeicolo.Validate(); err != nil {
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}

	if err = f.DatiPagamento.Validate(); err != nil {
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}

	if err = f.Allegati.Validate(); err != nil {
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}
	return nil
}
