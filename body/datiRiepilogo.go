package body

type datiRiepilogo struct {
	AliquotaIVA       float32 `xml:"AliquotaIVA" json:"AliquotaIVA"`
	ImponibileImporto float32 `xml:"ImponibileImporto" json:"ImponibileImporto"`
	Imposta           float32 `xml:"Imposta" json:"Imposta"`
	EsigibilitaIVA    string  `xml:"EsigibilitaIVA" json:"EsigibilitaIVA"`
}

// Validate ...
func (f datiRiepilogo) Validate() error {
	return nil
}
