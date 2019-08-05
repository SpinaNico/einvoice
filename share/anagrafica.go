package share

import (
	"fmt"
)

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

// Validate ...
func (a Anagrafica) Validate() error {
	if len(a.Denominazione) > 80 {
		return fmt.Errorf("Anagrafica (Denominazione): %s", ErrorMaxLength(80))
	}
	if len(a.Nome) > 60 {
		return fmt.Errorf("Anagrafica (Nome): %s", ErrorMaxLength(60))
	}
	if len(a.Cognome) > 60 {
		return fmt.Errorf("Anagrafica (Nome): %s", ErrorMaxLength(60))
	}
	l := len(a.Titolo)
	if !(l >= 2 && l <= 10) {
		return fmt.Errorf("Anagrafica (Titolo) %s", ErrorIncluded(2, 10))
	}

	cl := len(a.CodEORI)
	if !(cl >= 13 && cl <= 17) {
		return fmt.Errorf("Anagrafica (CodEORI) %s", ErrorIncluded(13, 17))
	}

	if a.Denominazione != "" && (a.Nome != "" || a.Cognome != "") {
		return fmt.Errorf("Anagrafica: you cannot write the field name surname if you have indicated a denomination")
	}

	return nil
}
