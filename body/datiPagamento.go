package body

type datiPagamento struct {
	CondizioniPagamento string             `xml:"CondizioniPagamento" json:"CondizioniPagamento"`
	DettaglioPagamento  dettaglioPagamento `xml:"DettaglioPagamento" json:"DettaglioPagamento"`
}

// Validate ...
func (f datiPagamento) Validate() error {
	return nil
}
