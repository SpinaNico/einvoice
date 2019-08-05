package body

//DatiBollo (i seguenti elementi vanno valorizzati nei casi in cui sia prevista l’imposta di bollo)
type datiBollo struct {
	// BolloVirtuale: indica l’assolvimento dell’imposta di bollo ai sensi del decreto MEF 17 giugno 2014 (bollovirtuale).
	BolloVirtuale string `xml:"BolloVirtuale" json:"BolloVirtuale"` // ?
	// ImportoBollo: importo dell’imposta di bollo.
	ImportoBollo float32 `xml:"ImportoBollo" json:"ImportoBollo"` // ?
}

// Validate ...
func (f datiBollo) Validate() error {
	return nil
}
