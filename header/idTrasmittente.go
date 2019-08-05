package header

//IdTrasmittente (identificativo fiscale del soggetto trasmittente
type idTrasmittente struct {
	//IdPaese: codice del paese assegnante l’identifcativo fiscale al soggetto trasmittente.
	IDPaese string `xml:"IdPaese" json:"IdPaese"`

	//IdCodice: numero di identificazione fiscale del trasmittente
	//(per i soggetti stabiliti nel territorio dello Stato Italiano
	//corrisponde al Codice Fiscale; per i non residenti si fa
	//riferimento all’identificativo fiscale assegnato dall’autorità del
	//paese di residenza). In caso di IdPaese uguale a IT, il sistema
	//ne verifica la presenza in Anagrafe Tributaria: se non esiste
	//come codice fiscale, il file viene scartato con codice errore	00300.
	IDCodice string `xml:"IdCodice" json:"IdCodice"`
}

func (c idTrasmittente) Validate() error {
	return nil
}
