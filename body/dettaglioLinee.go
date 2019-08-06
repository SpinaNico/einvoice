package body

type dettaglioLinee struct {
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

	//Quantita: formato numerico nel quale i decimali vanno separati dall'intero
	//con il carattere '.' (punto). La sua lunghezza va da 4 a 21 caratteri.
	Quantita float32 `xml:"Quantita" json:"Quantita"`

	// UnitaMisura: formato alfanumerico; lunghezza massima di 10 caratteri.
	UnitaMisura string `xml:"UnitaMisura" json:"UnitaMisura"`

	DataInizioPeriodo data `xml:"DataInizioPeriodo" json:"DataInizioPeriodo"`
	DataFinePeriodo   data `xml:"DataFinePeriodo" json:"DataFinePeriodo"`

	PrezzoUnitario decimale2 `xml:"PrezzoUnitario" json:"PrezzoUnitario"`

	ScontoMaggiorazione scontoMaggiorazione `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`

	PrezzoTotale decimale2 `xml:"PrezzoTotale" json:"PrezzoTotale"`
	AliquotaIVA  decimale2 `xml:"AliquotaIVA" json:"AliquotaIVA"`

	Ritenuta                   string              `xml:"Ritenuta" json:"Ritenuta"`
	Natura                     natura              `xml:"Natura" json:"Natura"`
	RiferimentoAmministrazione string              `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione"`
	AltriDatiGestionali        altriDatiGestionali `xml:"AltriDatiGestionali" json:"AltriDatiGestionali"`
}

// Validate ...
func (f dettaglioLinee) Validate() error {
	return nil
}
