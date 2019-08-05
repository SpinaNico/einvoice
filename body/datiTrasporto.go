package body

type datiTrasporto struct {
	DatiAnagraficiVettore datiAnagraficiVettore `xml:"DatiAnagraficiVettore" json:"DatiAnagraficiVettore"`
	DataOraConsegna       string                `xml:"DataOraConsegna" json:"DataOraConsegna"`
}

// Validate ...
func (f datiTrasporto) Validate() error {
	return nil
}
