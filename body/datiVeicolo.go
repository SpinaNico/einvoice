package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiVeicolo struct {
	// Data: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
	Data data `xml:"Data" json:"Data"`
	// TotalePercorso: formato alfanumerico; lunghezza massima di 15 caratteri.
	TotalePercorso string `xml:"TotalePercorso" json:"TotalePercorso"`
}

// Validate ...
func (f datiVeicolo) Validate() error {

	if err := f.Data.Validate(); err != nil {
		return fmt.Errorf("Data %s", err)
	}

	if len(f.TotalePercorso) > 15 {
		return fmt.Errorf("TotalePercorso %s", share.ErrorMaxLength(15))
	}

	return nil
}
