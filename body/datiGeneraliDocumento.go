package body

type datiGeneraliDocumento struct {
	//TipoDocumento: tipologia del documento oggetto della trasmissione
	//(fattura, acconto/anticipo su fattura, acconto/anticipo su parcella ,
	// nota di credito, nota di debito, parcella, autofattura).
	TipoDocumento string `xml:"TipoDocumento" json:"TipoDocumento"`

	// Divisa: tipo di valuta utilizzata per l'indicazione degli importi.
	Divisa string `xml:"Divisa" json:"Divisa"`

	// Data: data del documento.
	Data string `xml:"Data" json:"Data"`

	// Numero: numero progressivo attribuito dal cedente/prestatore
	// al documento. Deve contenere almeno un carattere numerico.
	// In caso contrario il file viene scartato con codice errore 00425.
	Numero string `xml:"Numero" json:"Numero"`

	// Causale: descrizione della causale del documento.
	Causale []string `xml:"Causale" json:"Causale"`

	// Campo che viene valorizzato in caso sia presente ritenuta.
	DatiRitenuta           datiRitenuta           `xml:"DatiRitenuta" json:"DatiRitenuta"`
	DatiBollosia           datiBollo              `xml:"DatiBollosia" json:"DatiBollosia"`
	DatiCassaPrevidenziale datiCassaPrevidenziale `xml:"DatiCassaPrevidenziale" json:"DatiCassaPrevidenziale "`
	ScontoMaggiorazione    scontoMaggiorazione    `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	ImportoTotaleDocumento float32                `xml:"ImportoTotaleDocumento" json:"ImportoTotaleDocumento"`
	Arrotondamento         float32                `xml:"Arrotondamento" json:"Arrotondamento"`
	Art73                  string                 `xml:"Art73" json:"Art73"`
}

// Validate ...
func (f datiGeneraliDocumento) Validate() error {
	return nil
}
