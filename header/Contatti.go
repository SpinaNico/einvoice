package header

type contatti struct {
	Telefono string `xml:"Telefono" json:"Telefono"`
	Fax      string `xml:"Fax" json:"Fax"`
	Email    string `xml:"Email" json:"Email"`
}

func (c contatti) Validate() error {
	return nil
}
