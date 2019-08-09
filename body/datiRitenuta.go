package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiRitenuta struct {
	// TipoRitenuta: tipologia di ritenuta (persone fisiche o persone giuridiche)
	// TipoRitenuta: formato alfanumerico; lunghezza di 4 caratteri;
	// ---
	// valori ammessi sono i seguenti:
	// + 	RT01 Ritenuta persone fisiche
	// +	RT02 Ritenuta persone giuridiche
	TipoRitenuta string `xml:"TipoRitenuta" json:"TipoRitenuta"`

	// ImportoRitenuta: importo della ritenuta
	// La sua lunghezza va da 4 a 15 caratteri.
	ImportoRitenuta decimale2 `xml:"ImportoRitenuta" json:"ImportoRitenuta"`

	// AliquotaRitenuta: aliquota (espressa in percentuale %) della ritenuta
	// formato numerico nel quale i decimali vanno separati dall'intero con il carattere,
	// ---
	// La sua lunghezza va da 4 a 6 caratteri.
	AliquotaRitenuta decimale2 `xml:"AliquotaRitenuta" json:"AliquotaRitenuta"`

	// CausalePagamento: codice della causale del
	// pagamento (il codice corrisponde a quello utilizzato per
	// la compilazione del modello CU; per il codice M2
	// indicare il valore M; per il codice ZO indicare il valore Z)
	// Lunghezza 2 caratteri
	CausalePagamento string `xml:"CausalePagamento" json:"CausalePagamento"`
}

// Validate ...
func (f datiRitenuta) Validate() error {

	if f.TipoRitenuta != "RT01" && f.TipoRitenuta != "RT02" {
		return fmt.Errorf("TipoRitenuta %s", share.ErrorIncorrectValue(f.TipoRitenuta))
	}

	if len(f.ImportoRitenuta) < 4 || len(f.ImportoRitenuta) > 15 {
		return fmt.Errorf("ImportoRitenuta %s", share.ErrorIncluded(4, 15))
	}
	if len(f.AliquotaRitenuta) < 4 || len(f.AliquotaRitenuta) > 6 {
		return fmt.Errorf("ImportoRitenuta %s", share.ErrorIncluded(4, 6))
	}
	if f.CausalePagamento != "CU" && f.CausalePagamento != "ZO" && f.CausalePagamento != "M2" {
		return fmt.Errorf("CausalePagamento %s", share.ErrorIncorrectValue(f.CausalePagamento))
	}
	return nil
}
