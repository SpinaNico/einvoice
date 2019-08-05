package header

type cedentePrestatore struct {
	DatiAnagrafici datiAnagrafici `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           sede           `xml:"Sede" json:"Sede"`
}

func (c cedentePrestatore) Validate() error {
	return nil
}
