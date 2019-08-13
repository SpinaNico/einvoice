package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
	"strings"
)

//DatiBollo (i seguenti elementi vanno valorizzati nei casi in cui sia prevista l’imposta di bollo)
type datiBollo struct {
	// BolloVirtuale: indica l’assolvimento dell’imposta di bollo ai sensi del decreto MEF 17 giugno 2014 (bollovirtuale).
	BolloVirtuale string `xml:"BolloVirtuale" json:"BolloVirtuale"` // ?
	// ImportoBollo: importo dell’imposta di bollo.
	ImportoBollo decimale2 `xml:"ImportoBollo" json:"ImportoBollo"` // ?
}

// Validate ...
func (f datiBollo) Validate() error {
	var err error

	f.BolloVirtuale = strings.ToUpper(f.BolloVirtuale)
	if f.BolloVirtuale != "SI" {
		return fmt.Errorf("DatiBollo (BolloVirtuale) %s only valid value is \"SI\" ", share.ErrorIncorrectValue(f.BolloVirtuale))
	}
	if err = f.ImportoBollo.Validate(); err != nil {
		return fmt.Errorf("ImportoBollo %s", err)
	}
	if len(string(f.ImportoBollo)) > 15 {
		return fmt.Errorf("ImportoBollo %s", share.ErrorMaxLength(15))
	}

	return nil
}
