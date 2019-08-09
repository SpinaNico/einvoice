package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiTrasporto struct {
	DatiAnagraficiVettore datiAnagraficiVettore `xml:"DatiAnagraficiVettore" json:"DatiAnagraficiVettore"`

	//formato alfanumerico; lunghezza massima di 80  caratteri.
	MezzoTrasporto string `xml:"MezzoTrasporto" json:"MezzoTrasporto"`
	// formato alfanumerico; lunghezza massima di 100 caratteri.
	CausaleTrasporto string `xml:"CausaleTrasporto" json:"CausaleTrasporto"`
	//NumeroColli: formato numerico; lunghezza massima di 4 caratteri.
	NumeroColli int `xml:"NumeroColli" json:"NumeroColli"`
	//Descrizione: formato alfanumerico; lunghezza massima di 100 caratteri.
	Descrizione string `xml:"Descrizione" json:"Descrizione"`
	//UnitaMisuraPeso: formato alfanumerico; lunghezza massima di 10 caratteri.
	UnitaMisuraPeso string `xml:"UnitaMisuraPeso" json:"UnitaMisuraPeso"`
	//PesoLordo: formato numerico La sua lunghezza va da 4 a 7 caratteri.
	PesoLordo decimale2 `xml:"PesoLordo" json:"PesoLordo"`
	// PesoNetto: formato numerico  La sua lunghezza va da 4 a 7 caratteri.
	PesoNetto decimale2 `xml:"PesoNetto" json:"PesoNetto"`
	//DataOraRitiro: la data e ora, devono essere in questo formato YYYY-MM- DDTHH:MM:SS.
	DataOraRitiro dataOra `xml:"DataOraRitiro" json:"DataOraRitiro"`
	// DataInizioTrasporto: la data deve essere  di questo formato YYYY-MM-DD.
	DataInizioTrasporto data `xml:"DataInizioTrasporto" json:"DataInizioTrasporto"`

	//TipoResa: codifica del termine di resa (Incoterms) espresso
	//secondo lo standard ICC-Camera di Commercio Internazionale
	//(formato alfanumerico di 3 caratteri)
	TipoResa string `xml:"TipoResa" json:"TipoResa"`

	//IndirizzoResa che si compone degli stessi campi previsti per
	//l’elemento Sede del CedentePrestatore contenuti nel tipo
	IndirizzoResa share.IndirizzoType `xml:"IndirizzoResa" json:"IndirizzoResa"`

	// YYYY-MM-DDTHH:MM:SS.
	DataOraConsegna dataOra `xml:"DataOraConsegna" json:"DataOraConsegna"`
}

// Validate ...
func (f datiTrasporto) Validate() error {
	if err := f.DatiAnagraficiVettore.Validate(); err != nil {
		return fmt.Errorf("DatiAnagraficiVettore %s", err)
	}
	if err := f.DataOraConsegna.Validate(); err != nil {
		return fmt.Errorf("DataOraConsegna %s", err)
	}
	return nil
}
