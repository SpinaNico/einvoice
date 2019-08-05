package header

type sede struct {
	Indirizzo string `xml:"Indirizzo" json:"Indirizzo"`
	CAP       string `xml:"CAP" json:"CAP"`
	Comune    string `xml:"Comune" json:"Comune"`
	Provincia string `xml:"Provincia" json:"Provincia"`
	Nazione   string `xml:"Nazione" json:"Nazione"`
}

func (c sede) Validate() error { return nil }
