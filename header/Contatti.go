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

	err = c.Email.Validate()
	if err != nil {
		return fmt.Errorf("Contatti %s", err)
	}

	err = c.Telefono.Validate()
	if err != nil {
		return fmt.Errorf("Contatti %s", err)
	}
	return nil
}
