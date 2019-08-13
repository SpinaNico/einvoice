package body

import (
	"fmt"
	share "github.com/SpinaNico/go-struct-invoice/share"
)

type datiAnagraficiVettore struct {
	Anagrafica   share.Anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA share.IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}

// Validate ...
func (f datiAnagraficiVettore) Validate() error {
	var err error

	if err = f.Anagrafica.Validate(); err != nil {
		return fmt.Errorf("DatiAnagraficiVettore %s", err)
	}

	if err = f.IDFiscaleIVA.Validate(); err != nil {
		return fmt.Errorf("DatiAnagraficiVettore %s", err)
	}
	return nil
}
