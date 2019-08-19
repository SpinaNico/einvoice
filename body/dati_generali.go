package body

import "fmt"

type datiGenerali struct {
	DatiGeneraliDocumento datiGeneraliDocumento `xml:"DatiGeneraliDocumento" json:"DatiGeneraliDocumento"`
	DatiOrdineAcquisto    datiOrdineAcquisto    `xml:"DatiOrdineAcquisto,omitempity" json:"DatiOrdineAcquisto"`
	DatiContratto         datiContratto         `xml:"DatiContratto" json:"DatiContratto"`
	DatiConvenzione       datiConvenzione       `xml:"DatiConvenzione" json:"DatiConvenzione"`
	DatiRicezione         datiRicezione         `xml:"DatiRicezione" json:"DatiRicezione"`
	DatiFattureCollegate  datiFattureCollegate  `xml:"DatiFattureCollegate" json:"DatiFattureCollegate"`
	DatiSAL               datiSAL               `xml:"DatiSAL" json:"DatiSAL"`
	DatiDDT               datiDDT               `xml:"DatiDDT" json:"DatiDDT"`
	DatiTrasporto         datiTrasporto         `xml:"DatiTrasporto" json:"DatiTrasporto"`
	FatturaPrincipale     fatturaPrincipale     `xml:"FatturaPrincipale" json:"FatturaPrincipale"`
}

// Validate ...
func (f datiGenerali) Validate() error {

	if err := f.DatiGeneraliDocumento.Validate(); err != nil {
		return fmt.Errorf("DatiGeneraliDocumento %s", err)
	}

	if err := f.DatiOrdineAcquisto.Validate(); err != nil {
		return fmt.Errorf("DatiOrdineAcquisto %s", err)
	}

	if err := f.DatiContratto.Validate(); err != nil {
		return fmt.Errorf("DatiContratto %s", err)
	}

	if err := f.DatiConvenzione.Validate(); err != nil {
		return fmt.Errorf("DatiConvenzione %s", err)
	}
	if err := f.DatiRicezione.Validate(); err != nil {
		return fmt.Errorf("DatiRicezione %s", err)
	}
	if err := f.DatiSAL.Validate(); err != nil {
		return fmt.Errorf("DatiSAL %s", err)
	}

	if err := f.DatiDDT.Validate(); err != nil {
		return fmt.Errorf("DatiDDT %s", err)
	}

	if err := f.DatiTrasporto.Validate(); err != nil {
		return fmt.Errorf("DatiTrasporto %s", err)
	}

	return nil
}
