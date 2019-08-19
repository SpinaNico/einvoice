package header

type contattiTrasmittente struct {
	Telefono string `xml:"Telefono" json:"Telefono" validate:"omitempty,min=5,max=12"`
	Email    string `xml:"Email" json:"Email" validate:"omitempty,email,max=255"`
}
