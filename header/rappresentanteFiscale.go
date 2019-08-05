package header

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
)

type rappresentanteFiscale struct {
	share.Anagrafica
	IDFiscaleIva share.IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}

func (a rappresentanteFiscale) Validate() error {
	var err error
	err = a.IDFiscaleIva.Validate()
	if err != nil {
		return fmt.Errorf("RappresentanteFiscale %s", err)
	}
	if len(a.Denominazione) > 80 {
		return fmt.Errorf("RappresentanteFiscale (Denominazione): %s", share.ErrorMaxLength(80))
	}
	if len(a.Nome) > 60 {
		return fmt.Errorf("RappresentanteFiscale (Nome): %s", share.ErrorMaxLength(60))
	}
	if len(a.Cognome) > 60 {
		return fmt.Errorf("RappresentanteFiscale (Nome): %s", share.ErrorMaxLength(60))
	}
	l := len(a.Titolo)
	if !(l >= 2 && l <= 10) {
		return fmt.Errorf("RappresentanteFiscale (Titolo) %s", share.ErrorIncluded(2, 10))
	}

	cl := len(a.CodEORI)
	if !(cl >= 13 && cl <= 17) {
		return fmt.Errorf("RappresentanteFiscale (CodEORI) %s", share.ErrorIncluded(13, 17))
	}

	if a.Denominazione != "" && (a.Nome != "" || a.Cognome != "") {
		return fmt.Errorf("RappresentanteFiscale: you cannot write the field name surname if you have indicated a denomination")
	}
	return nil
}
