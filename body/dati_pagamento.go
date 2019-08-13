package body

import (
	"fmt"
	"strings"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiPagamento struct {
	CondizioniPagamento string             `xml:"CondizioniPagamento" json:"CondizioniPagamento"`
	DettaglioPagamento  dettaglioPagamento `xml:"DettaglioPagamento" json:"DettaglioPagamento"`
}

// Validate ...
func (f datiPagamento) Validate() error {
	f.CondizioniPagamento = strings.ToUpper(f.CondizioniPagamento)
	switch f.CondizioniPagamento {
	case "TP01":
		break
	case "TP02":
		break
	case "TP03":
		break
	default:
		return fmt.Errorf("CondizioniPagamento %s", share.ErrorIncorrectValue(f.CondizioniPagamento))
	}
	if err := f.DettaglioPagamento.Validate(); err != nil {
		return fmt.Errorf("DettaglioPagamento %s", err)
	}
	return nil
}
