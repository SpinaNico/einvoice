package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
)

type altriDatiGestionali struct {
	TipoDato          string    `xml:"TipoDato" json:"TipoDato"`
	RiferimentoTesto  string    `xml:"RiferimentoTesto" json:"RiferimentoTesto"`
	RiferimentoNumero decimale2 `xml:"RiferimentoNumero" json:"RiferimentoNumero"`
	RiferimentoData   data      `xml:"RiferimentoData" json:"RiferimentoData"`
}

func (c altriDatiGestionali) Validate() error {
	var err error

	if len(c.TipoDato) > 10 {
		return fmt.Errorf("TipoDato %s", share.ErrorMaxLength(10))
	}

	if len(c.RiferimentoTesto) > 60 {
		return fmt.Errorf("RiferimentoTesto %s", share.ErrorMaxLength(60))
	}

	if err = c.RiferimentoData.Validate(); err != nil {
		return fmt.Errorf("RiferimentoData %s", err)
	}

	if len(string(c.RiferimentoNumero)) > 21 && len(string(c.RiferimentoNumero)) < 4 {
		return fmt.Errorf("RiferimentoNumero %s", share.ErrorIncluded(4, 21))
	}
	if err = c.RiferimentoNumero.Validate(); err != nil {
		return fmt.Errorf("RiferimentoNumero %s", err)
	}
	return nil
}
