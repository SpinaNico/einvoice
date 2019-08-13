package header

import (
	"fmt"

	share "github.com/SpinaNico/go-struct-invoice/share"
)

type cedentePrestatore struct {
	DatiAnagrafici datiAnagrafici      `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           share.IndirizzoType `xml:"Sede" json:"Sede"`
	Contatti       contatti            `xml:"Contatti" json:"Contatti"`
}

func (c cedentePrestatore) Validate() error {
	var err error
	err = c.Contatti.Validate()
	if err != nil {
		return fmt.Errorf("CedentePrestatore %s", err)
	}

	err = c.Sede.Validate()
	if err != nil {
		return fmt.Errorf("CedentePrestatore %s", err)
	}

	err = c.DatiAnagrafici.Validate()
	if err != nil {
		return fmt.Errorf("CedentePrestatore %s", err)
	}

	return nil
}
