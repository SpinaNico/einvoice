package body

type fatturaPrincipale struct {
	//NumeroFatturaPrincipale: formato alfanumerico; lunghezza massima di 20 caratteri.
	NumeroFatturaPrincipale string `xml:"NumeroFatturaPrincipale" json:"NumeroFatturaPrincipale"`
	//DataFatturaPrincipale: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
	DataFatturaPrincipale data `xml:"DataFatturaPrincipale" json:"DataFatturaPrincipale"`
}

func (c fatturaPrincipale) Validate() error {
	return nil
}
