package body

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
	return nil
}
