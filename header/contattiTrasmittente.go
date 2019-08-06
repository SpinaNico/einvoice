package header

import (
	"fmt"
)

type contattiTrasmittente struct {
	Telefono telefono `xml:"Telefono" json:"Telefono"`
	Email    email    `xml:"Email" json:"Email"`
}

func (c contattiTrasmittente) Validate() error {
	var err error

	if err = c.Telefono.Validate(); err != nil {
		return fmt.Errorf("ContattiTrasmittente %s", err)
	}

	if err = c.Email.Validate(); err != nil {
		return fmt.Errorf("ContattiTrasmittente %s", err)
	}
	return nil
}
