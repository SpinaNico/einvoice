package body

type datiGenerali struct {
	DatiGeneraliDocumento datiGeneraliDocumento `xml:"DatiGeneraliDocumento" json:"DatiGeneraliDocumento"`
	DatiOrdineAcquisto    datiOrdineAcquisto    `xml:"DatiOrdineAcquisto" json:"DatiOrdineAcquisto"`
	DatiContratto         datiContratto         `xml:"DatiContratto" json:"DatiContratto"`
	DatiConvenzione       datiConvenzione       `xml:"DatiConvenzione" json:"DatiConvenzione"`
	DatiRicezione         datiRicezione         `xml:"DatiRicezione" json:"DatiRicezione"`
	//DatiFattureCollegate
	DatiSAL       datiSAL       `xml:"DatiSAL" json:"DatiSAL"`
	DatiDDT       datiDDT       `xml:"DatiDDT" json:"DatiDDT"`
	DatiTrasporto datiTrasporto `xml:"DatiTrasporto" json:"DatiTrasporto"`
	//FatturaPrincipale
}

// Validate ...
func (f datiGenerali) Validate() error {
	return nil
}
