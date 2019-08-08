package body

type datiGeneraliDocumento struct {
	//TipoDocumento: tipologia del documento oggetto della trasmissione
	//(fattura, acconto/anticipo su fattura, acconto/anticipo su parcella ,
	// nota di credito, nota di debito, parcella, autofattura).
	// + TD01 Fattura
	// + TD02 Acconto/Anticipo su fattura
	// + TD03 Acconto/Anticipo su parcella
	// + TD04 Nota di Credito
	// + TD05 Nota di Debito
	// + TD06 Parcella
	// + TD20 Autofattura
	TipoDocumento tipoDocumento `xml:"TipoDocumento" json:"TipoDocumento"`

	// Divisa: tipo di valuta utilizzata per l'indicazione degli importi.
	Divisa divisa `xml:"Divisa" json:"Divisa"`

	// Data: data del documento.
	Data data `xml:"Data" json:"Data"`

	// Numero: numero progressivo attribuito dal cedente/prestatore
	// al documento. Deve contenere almeno un carattere numerico.
	// In caso contrario il file viene scartato con codice errore 00425.
	Numero string `xml:"Numero" json:"Numero"`
	// Campo che viene valorizzato in caso sia presente ritenuta.
	DatiRitenuta           datiRitenuta           `xml:"DatiRitenuta" json:"DatiRitenuta"`
	DatiBollo              datiBollo              `xml:"DatiBollo" json:"DatiBollo"`
	DatiCassaPrevidenziale datiCassaPrevidenziale `xml:"DatiCassaPrevidenziale" json:"DatiCassaPrevidenziale "`
	ScontoMaggiorazione    scontoMaggiorazione    `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	// max length 15 characters and minium length 4 (00.00)
	ImportoTotaleDocumento decimale2 `xml:"ImportoTotaleDocumento" json:"ImportoTotaleDocumento"`
	// max length 15 characters and minium length 4 (00.00)
	Arrotondamento decimale2 `xml:"Arrotondamento" json:"Arrotondamento"`
	// Causale: descrizione della causale del documento.
	Causale []string `xml:"Causale" json:"Causale"`
	Art73   string   `xml:"Art73" json:"Art73"`
}

// Validate ...
func (f datiGeneraliDocumento) Validate() error {
	return nil
}
