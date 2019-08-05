package body

import share "invoice/share"

type datiAnagraficiVettore struct {
	Anagrafica   share.Anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA share.IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}

// Validate ...
func (f datiAnagraficiVettore) Validate() error {
	return nil
}