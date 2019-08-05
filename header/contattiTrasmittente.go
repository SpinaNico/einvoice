package header

type contattiTrasmittente struct {
	Telefono string `xml:"Telefono" json:"Telefono"`
	Email    string `xml:"Email" json:"Email"`
}

func (c contattiTrasmittente) Validate() error {
	return nil
}
