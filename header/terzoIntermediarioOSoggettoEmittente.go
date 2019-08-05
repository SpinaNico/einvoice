package header

import share "invoice/share"

type terzoIntermediarioOSoggettoEmittente struct {
	IDFiscaleIVA  share.IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
	CodiceFiscale string             ` xml:"CodiceFiscale" json:"CodiceFiscale"`
	Anagrafica    share.Anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
}

func (c terzoIntermediarioOSoggettoEmittente) Validate() error {
	return nil
}
