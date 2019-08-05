package header

import (
	"fmt"
)

type sede struct {
	Indirizzo    indirizzo    `xml:"Indirizzo" json:"Indirizzo"`
	NumeroCivico numeroCivico `xml:"NumeroCivico" json:"NumeroCivico"`
	// **IT** Codice avviamento postale
	// **EN** postal code for info:
	// [wikipedia](https://en.wikipedia.org/wiki/Codice_di_avviamento_postale)
	CAP       Cap       `xml:"CAP" json:"CAP"`
	Comune    comune    `xml:"Comune" json:"Comune"`
	Provincia provincia `xml:"Provincia" json:"Provincia"`
	Nazione   nazione   `xml:"Nazione" json:"Nazione"`
}

func (c sede) Validate() error {
	var err error
	err = c.CAP.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}

	err = c.Comune.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}

	err = c.Provincia.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}

	err = c.Indirizzo.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}

	err = c.Nazione.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}

	err = c.NumeroCivico.Validate()
	if err != nil {
		return fmt.Errorf("Sede %s", err)
	}
	return nil
}
