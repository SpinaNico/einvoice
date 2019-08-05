package body

type datiVeicolo struct {
	// Data: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
	Data string `xml:"Data" json:"Data"`
	// TotalePercorso: formato alfanumerico; lunghezza massima di 15 caratteri.
	TotalePercorso string `xml:"TotalePercorso" json:"TotalePercorso"`
}

// Validate ...
func (f datiVeicolo) Validate() error {
	return nil
}
