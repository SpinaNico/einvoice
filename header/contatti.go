package header

import (
	"fmt"
)

type contatti struct {
	Telefono telefono `xml:"Telefono" json:"Telefono"`
	//todo: Validate Fax
	Fax   string `xml:"Fax" json:"Fax"`
	Email email  `xml:"Email" json:"Email"`
}

func (c contatti) Validate() error {
	var err error

	if err = c.Email.Validate(); err != nil {
		return fmt.Errorf("Email %s", err)
	}

	if err = c.Telefono.Validate(); err != nil {
		return fmt.Errorf("Telefono %s", err)
	}
	return nil
}
