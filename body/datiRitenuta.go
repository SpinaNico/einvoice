package body

type datiRitenuta struct {
	// TipoRitenuta: tipologia di ritenuta (persone fisiche o persone giuridiche).
	TipoRitenuta string `xml:"TipoRitenuta" json:"TipoRitenuta"`

	// ImportoRitenuta: importo della ritenuta.
	ImportoRitenuta float32 `xml:"ImportoRitenuta" json:"ImportoRitenuta"`

	// 	AliquotaRitenuta: aliquota (espressa in percentuale %) della ritenuta.
	AliquotaRitenuta float32 `xml:"AliquotaRitenuta" json:"AliquotaRitenuta"`

	// CausalePagamento: codice della causale del
	// pagamento (il codice corrisponde a quello utilizzato per
	// la compilazione del modello CU; per il codice M2
	// indicare il valore M; per il codice ZO indicare il valore Z).
	CausalePagamento string `xml:"CausalePagamento" json:"CausalePagamento"`
}

// Validate ...
func (f datiRitenuta) Validate() error {
	return nil
}
