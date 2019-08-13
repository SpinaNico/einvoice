package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
)

type codiceArticolo struct {
	//CodiceTipo: formato alfanumerico; lunghezza massima di 35 caratteri.
	CodiceTipo string `xml:"CodiceTipo" json:"CodiceTipo"`

	//CodiceValore: formato alfanumerico; lunghezza massima di 35 caratteri.
	CodiceValore string `xml:"CodiceValore" json:"CodiceValore"`
}

func (c codiceArticolo) Validate() error {

	if len(c.CodiceTipo) > 35 {
		return fmt.Errorf("CodiceTipo %s", share.ErrorMaxLength(35))
	}
	if len(c.CodiceValore) > 35 {
		return fmt.Errorf("CodiceValore %s", share.ErrorMaxLength(35))
	}

	return nil
}
