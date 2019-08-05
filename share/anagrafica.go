package share

//Anagrafica ...
type Anagrafica struct {
	Denominazione string `xml:"Denominazione" json:"Denominazione"`

	// Nome Qual'ora è una persona fisica
	Nome string `xml:"Nome" json:"Nome"`

	// Cognome qual'ora è una persona fisica
	Cognome string `xml:"Cognome" json:"Cognome"`

	// Tiolo onorifico
	Titolo string `xml:"Titolo" json:"Titolo"`

	//CodEORI: numero del Codice EORI (Economic
	//Operator Registration and Identification) in base
	//al Regolamento (CE) n. 312 del 16 aprile 2009.
	//In vigore dal 1 Luglio 2009 tale codice identifica
	//gli operatori economici nei rapporti con le autorità
	//doganali sull'intero territorio dell'Unione Europea.
	CodEORI string `xml:"CodEORI" json:"CodEORI"`
}

// Validate
func (a Anagrafica) Validate() {

}
