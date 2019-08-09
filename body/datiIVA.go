package body

import "fmt"

type datiIVA struct {
	AliquotaIVA decimale2 `xml:"Aliquota" json:"Aliquota"`
	Imposta     decimale2 `xml:"Imposta" json:"Imposta"`
}

func (c datiIVA) Validate() error {
	if err := c.AliquotaIVA.Validate(); err != nil {
		return fmt.Errorf("AliquotaIVA %s", err)
	}
	if err := c.Imposta.Validate(); err != nil {
		return fmt.Errorf("Imposta %s", err)
	}
	return nil
}
