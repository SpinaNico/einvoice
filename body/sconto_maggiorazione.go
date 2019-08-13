package body

type scontoMaggiorazione struct {

	//Tipo: indica se si tratta di sconto o di maggiorazione.
	Tipo string `xml:"Tipo" json:"Tipo"`
	//Percentuale: percentuale di sconto o di maggiorazione.
	Percentuale float32 `xml:"Percentuale" json:"Percentuale"`
	//Importo: importo dello sconto o della maggiorazione.
	Importo float32 `xml:"Importo" json:"Importo"`
}

// Validate ...
func (f scontoMaggiorazione) Validate() error {
	return nil
}
