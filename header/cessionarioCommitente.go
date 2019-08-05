package header

type cessionarioCommittente struct {
	DatiAnagrafici datiAnagrafici `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           sede           `xml:"Sede" json:"Sede"`
}

func (c cessionarioCommittente) Validate() error {
	return nil
}
