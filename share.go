package invoice

type anagrafica struct {
	Denominazione string `xml:"Denominazione" json:"Denominazione" validate:"omitempty,max=80,required_without_all=Nome Cognome"`

	// Nome Qual'ora è una persona fisica
	Nome string `xml:"Nome" json:"Nome" validate:"omitempty,max=60,required_with=Cognome"`

	// Cognome qual'ora è una persona fisica
	Cognome string `xml:"Cognome" json:"Cognome" validate:"omitempty,max=60,required_with=Nome"`

	// Tiolo onorifico
	Titolo string `xml:"Titolo" json:"Titolo" validate:"omitempty,min=2,max=10"`

	//CodEORI: numero del Codice EORI (Economic
	//Operator Registration and Identification) in base
	//al Regolamento (CE) n. 312 del 16 aprile 2009.
	//In vigore dal 1 Luglio 2009 tale codice identifica
	//gli operatori economici nei rapporti con le autorità
	//doganali sull'intero territorio dell'Unione Europea.
	CodEORI string `xml:"CodEORI" json:"CodEORI" validate:"omitempty,min=13,max=17"`
}

// Validate ...
func (a anagrafica) Validate() error {
	if a.Denominazione != "" && (a.Nome != "" || a.Cognome != "") {
		//	return fmt.Errorf("Anagrafica: you cannot write the field name surname if you have indicated a denomination")
	}

	return nil
}

type iDFiscaleIVA struct {
	// IDPaese: deve essere espresso secondo lo standard
	// ISO 3166-1 alpha-2 è utilizzato per indicare i luoghi.
	// esempio: Italy 		->	IT
	// esempio: Australia 	-> 	AU
	// esempio: Grece 		->	GR
	// maggiori informazioni: [wikipedia](https://it.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	IDPaese string `xml:"IdPaese" json:"IdPaese" validate:"omitempty,len=2"`
	// IDCodice : Lunghezza massima 28 caratteri alfanumerici.
	IDCodice string `xml:"IdCodice" json:"IdCodice" validate:"omitempty,max=28"`
}

type indirizzoType struct {
	Indirizzo    string `xml:"Indirizzo" json:"Indirizzo" validate:"max=60"`
	NumeroCivico string `xml:"NumeroCivico" json:"NumeroCivico" validate:"max=8"`
	// **IT** Codice avviamento postale
	// **EN** postal code for info:
	// [wikipedia](https://en.wikipedia.org/wiki/Codice_di_avviamento_postale)
	CAP       string `xml:"CAP" json:"CAP" validate:"len=5,isInteger"`
	Comune    string `xml:"Comune" json:"Comune" validate:"max=60"`
	Provincia string `xml:"Provincia" json:"Provincia" validate:"len=2"`
	Nazione   string `xml:"Nazione" json:"Nazione" validate:"len=2"`
}
