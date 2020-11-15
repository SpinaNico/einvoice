package einvoice

type Anagrafica struct {
	Denominazione string `xml:"Denominazione" json:"Denominazione" validate:"omitempty,max=80,required_without_all=Nome Cognome"`
	Nome          string `xml:"Nome" json:"Nome" validate:"omitempty,max=60,required_with=Cognome"`
	Cognome       string `xml:"Cognome" json:"Cognome" validate:"omitempty,max=60,required_with=Nome"`
	Titolo        string `xml:"Titolo" json:"Titolo" validate:"omitempty,min=2,max=10"`
	CodEORI       string `xml:"CodEORI" json:"CodEORI" validate:"omitempty,min=13,max=17"`
}

type IDFiscaleIVA struct {
	IDPaese  string `xml:"IdPaese" json:"IdPaese" validate:"len=2"`
	IDCodice string `xml:"IdCodice" json:"IdCodice" validate:"max=28,isPIVA"`
}

type IndirizzoType struct {
	Indirizzo    string `xml:"Indirizzo" json:"Indirizzo" validate:"max=60,required"`
	NumeroCivico string `xml:"NumeroCivico" json:"NumeroCivico" validate:"required,max=8"`
	CAP          string `xml:"CAP" json:"CAP" validate:"len=5,isInteger,required"`
	Comune       string `xml:"Comune" json:"Comune" validate:"required,max=60"`
	Provincia    string `xml:"Provincia" json:"Provincia" validate:"required,len=2"`
	Nazione      string `xml:"Nazione" json:"Nazione" validate:"required,len=2"`
}
