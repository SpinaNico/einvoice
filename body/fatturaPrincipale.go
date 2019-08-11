package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type fatturaPrincipale struct {
	//NumeroFatturaPrincipale: formato alfanumerico; lunghezza massima di 20 caratteri.
	NumeroFatturaPrincipale string `xml:"NumeroFatturaPrincipale" json:"NumeroFatturaPrincipale"`
	//DataFatturaPrincipale: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
	DataFatturaPrincipale data `xml:"DataFatturaPrincipale" json:"DataFatturaPrincipale"`
}

func (c fatturaPrincipale) Validate() error {

	if len(c.NumeroFatturaPrincipale) > 20 {
		return fmt.Errorf("NumeroFatturaPrincipale %s", share.ErrorMaxLength(20))
	}

	if err := c.DataFatturaPrincipale.Validate(); err != nil {
		return fmt.Errorf("DataFatturaPrincipale %s", err)
	}

	return nil
}
