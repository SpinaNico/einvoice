package body

type dettaglioLinee struct {
	NumeroLinea    int     `xml:"NumeroLinea" json:"NumeroLinea"`
	Descrizione    string  `xml:"Descrizione" json:"Descrizione"`
	Quantita       float32 `xml:"Quantita" json:"Quantita"`
	PrezzoUnitario float32 `xml:"PrezzoUnitario" json:"PrezzoUnitario"`
	PrezzoTotale   float32 `xml:"PrezzoTotale" json:"PrezzoTotale"`
	AliquotaIVA    float32 `xml:"AliquotaIVA" json:"AliquotaIVA"`
}

// Validate ...
func (f dettaglioLinee) Validate() error {
	return nil
}
