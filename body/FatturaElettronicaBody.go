package body

//FatturaElettronicaBody corpo della fattura
type FatturaElettronicaBody struct {
	DatiGenerali    datiGenerali    `xml:"DatiGenerali" json:"DatiGenerali"`
	DatiBeniServizi datiBeniServizi `xml:"DatiBeniServizi" json:"DatiBeniServizi"`
	DatiPagamento   datiPagamento   `xml:"DatiPagamento" json:"DatiPagamento"`
	DatiVeicolo     datiVeicolo     `xml:"DatiVeicolo" json:"DatiVeicolo"`
	Allegati        allegati        `xml:"Allegati" json:"Allegati"`
}

// Validate ...
func (f FatturaElettronicaBody) Validate() error {
	return nil
}
