package body

import (
	"fmt"
	"strings"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type dettaglioLinee struct {
	// formato numerico; lunghezza massima di 4 caratteri.
	NumeroLinea int `xml:"NumeroLinea" json:"NumeroLinea"`

	//TipoCessionePrestazione: formato alfanumerico; lunghezza di 2
	// caratteri; i ivalori ammessi sono:
	// + SC Sconto
	// + PR Premio
	// + AB Abbuono
	// + AC Spesa accessoria
	TipoCessionePrestazione string `xml:"TipoCessionePrestazione" json:"TipoCessionePrestazione"`

	CodiceArticolo codiceArticolo `xml:"CodiceArticolo" json:"CodiceArticolo"`

	//Descrizione: formato alfanumerico; lunghezza massima di 1000 caratteri.
	Descrizione string `xml:"Descrizione" json:"Descrizione"`

	//Quantita: formato numerico La sua lunghezza va da 4 a 21 caratteri.
	Quantita decimale2 `xml:"Quantita" json:"Quantita"`

	// UnitaMisura: formato alfanumerico; lunghezza massima di 10 caratteri.
	UnitaMisura string `xml:"UnitaMisura" json:"UnitaMisura"`

	DataInizioPeriodo data `xml:"DataInizioPeriodo" json:"DataInizioPeriodo"`
	DataFinePeriodo   data `xml:"DataFinePeriodo" json:"DataFinePeriodo"`
	// La sua lunghezza va da 4 a 21 caratteri.
	PrezzoUnitario decimale2 `xml:"PrezzoUnitario" json:"PrezzoUnitario"`

	ScontoMaggiorazione scontoMaggiorazione `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	// La sua lunghezza va da4 a 21 caratteri.
	PrezzoTotale decimale2 `xml:"PrezzoTotale" json:"PrezzoTotale"`
	// La sua lunghezza va da 4 a 6 caratteri.
	AliquotaIVA decimale2 `xml:"AliquotaIVA" json:"AliquotaIVA"`
	// only =="SI"
	Ritenuta string `xml:"Ritenuta" json:"Ritenuta"`
	Natura   natura `xml:"Natura" json:"Natura"`
	// RiferimentoAmministrazione: formato alfanumerico; lunghezza massima di 20 caratteri.
	RiferimentoAmministrazione string              `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione"`
	AltriDatiGestionali        altriDatiGestionali `xml:"AltriDatiGestionali" json:"AltriDatiGestionali"`
}

// Validate ...
func (f dettaglioLinee) Validate() error {

	if f.NumeroLinea > 9999 || f.NumeroLinea < 1 {
		return fmt.Errorf("NumeroLinea %s", share.ErrorIncluded(1, 9999))
	}
	f.TipoCessionePrestazione = strings.ToUpper(f.TipoCessionePrestazione)
	switch f.TipoCessionePrestazione {
	case "SC":
		break
	case "PR":
		break
	case "AB":
		break
	case "AC":
		break
	default:
		return fmt.Errorf("TipoCessionePrestazione %s", share.ErrorIncorrectValue(f.TipoCessionePrestazione))
	}

	if err := f.CodiceArticolo.Validate(); err != nil {
		return fmt.Errorf("CodiceArticolo %s", err)
	}

	if len(f.Descrizione) > 1000 {
		return fmt.Errorf("Descrizione %s", share.ErrorMaxLength(1000))
	}

	if len(f.Quantita) < 4 || len(f.Quantita) > 21 {
		return fmt.Errorf("Quantita %s", share.ErrorIncluded(4, 21))
	}

	if err := f.Quantita.Validate(); err != nil {
		return fmt.Errorf("Quantita %s", err)
	}

	if len(f.UnitaMisura) > 10 {
		return fmt.Errorf("UnitaMisura %s", share.ErrorMaxLength(10))
	}

	if err := f.DataInizioPeriodo.Validate(); err != nil {
		return fmt.Errorf("DataInizioPeriodo %s", err)
	}

	if err := f.DataFinePeriodo.Validate(); err != nil {
		return fmt.Errorf("DataFinePeriodo %s", err)
	}

	if err := f.PrezzoUnitario.Validate(); err != nil {
		return fmt.Errorf("PrezzoUnitario %s", err)
	}
	if len(f.PrezzoUnitario) < 4 || len(f.PrezzoUnitario) > 21 {
		return fmt.Errorf("PrezzoUnitario %s", share.ErrorIncluded(4, 21))
	}

	if err := f.ScontoMaggiorazione.Validate(); err != nil {
		return fmt.Errorf("ScontoMaggiorazione %s", err)
	}

	if err := f.PrezzoTotale.Validate(); err != nil {
		return fmt.Errorf("PrezzoTotale %s", err)
	}
	if len(f.PrezzoTotale) < 4 || len(f.PrezzoTotale) > 21 {
		return fmt.Errorf("PrezzoTotale %s", share.ErrorIncluded(4, 21))
	}

	if err := f.AliquotaIVA.Validate(); err != nil {
		return fmt.Errorf("AliquotaIVA %s", err)
	}
	if len(f.AliquotaIVA) < 4 || len(f.AliquotaIVA) > 6 {
		return fmt.Errorf("AliquotaIVA %s", share.ErrorIncluded(4, 6))
	}
	f.Ritenuta = strings.ToUpper(f.Ritenuta)
	if f.Ritenuta != "SI" {
		return fmt.Errorf("Ritenuta %s", share.ErrorIncorrectValue(f.Ritenuta))
	}

	if err := f.Natura.Validate(); err != nil {
		return fmt.Errorf("Natura %s", err)
	}

	if len(f.RiferimentoAmministrazione) > 20 {
		return fmt.Errorf("RiferimentoAmministrazione %s", share.ErrorMaxLength(20))
	}

	if err := f.AltriDatiGestionali.Validate(); err != nil {
		return fmt.Errorf("AltriDatiGestionali %s", err)
	}

	return nil
}
