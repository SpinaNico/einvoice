package share

// IDFiscaleIVA ...
type IDFiscaleIVA struct {
	IDPaese  string `xml:"IdPaese" json:"IdPaese"`
	IDCodice string `xml:"IdCodice" json:"IdCodice"`
}

// Validate ...
func (a IDFiscaleIVA) Validate() error {
	// ....
	return nil
}

// Doc ...
func (a IDFiscaleIVA) Doc() string {
	return `
	Bla bla bla bla
	`
}
