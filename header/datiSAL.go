package header

//2.2.9.3 Dati Riferimento Sal
//Blocco da valorizzare nei casi di fattura per stato di avanzamento
//RiferimentoFase: formato numerico; lunghezza massima di 3 caratteri.
type datiSAL struct {
	//RiferimentoFase: fase dello stato avanzamento cui la fatturasi riferisce.
	RiferimentoFase int `xml:"RiferimentoFase" json:"RiferimentoFase"`
}

func (c datiSAL) Validate() error {
	return nil
}
