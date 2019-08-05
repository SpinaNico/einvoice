package body

type dettaglioPagamento struct {
	ModalitaPagamento     string  `xml:"ModalitaPagamento" json:"ModalitaPagamento"`
	DataScadenzaPagamento string  `xml:"DataScadenzaPagamento" json:"DataScadenzaPagamento"`
	ImportoPagamento      float32 `xml:"ImportoPagamento" json:"ImportoPagamento"`
}

// Validate ...
func (f dettaglioPagamento) Validate() error {
	return nil
}
