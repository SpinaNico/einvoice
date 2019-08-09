package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

//2.2.9.3 Dati Riferimento Sal
//Blocco da valorizzare nei casi di fattura per stato di avanzamento
//RiferimentoFase: formato numerico; lunghezza massima di 3 caratteri.
type datiSAL struct {
	//RiferimentoFase: fase dello stato avanzamento cui la fatturasi riferisce.
	RiferimentoFase int `xml:"RiferimentoFase" json:"RiferimentoFase"`
}

func (c datiSAL) Validate() error {
	if c.RiferimentoFase > 999 {
		return fmt.Errorf("DatiSal %s", share.ErrorIncluded(0, 999))
	}
	return nil
}
