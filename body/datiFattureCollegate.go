package body

import "fmt"

//DatiFattureCollegate (dati relativi alla fattura alla quale si collega il
//	documento in oggetto)
type datiFattureCollegate struct {
	//RiferimentoNumeroLinea: numero della linea o delle linee di
	//dettaglio del documento alle quali si riferisce la fattura
	//collegata così come identificata dai tre elementi successivi
	//(IdDocumento, Data, NumItem); nel caso in cui la fattura
	//collegata si riferisce all’intero documento, questo elemento
	//non deve essere valorizzato.
	RiferimentoNumeroLinea numeroLinea `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea"`

	//IdDocumento: numero della fattura collegata associata al
	//documento o alla linea/linee del documento indicate
	//nell’elemento RiferimentoNumeroLinea.
	IDDocumento idDocumento `xml:"IdDocumento" json:"IdDocumento"`

	//Data: data della fattura collegata associata al documento o
	//alla linea/linee del documento indicate nell’elemento
	//RiferimentoNumeroLinea. Non può essere mai successiva alla
	//data del documento in oggetto; in caso contrario il file viene
	//scartato con codice errore 00418.
	Data data `xml:"Data" json:"Data"`

	//NumItem: identificativo della singola voce (linea di fattura
	//collegata) all'interno della fattura collegata associata al
	//documento o alla linea/linee del documento indicate
	//nell’elemento RiferimentoNumeroLinea.
	NumItem numItem `xml:"NumItem" json:"NumItem"`

	//CodiceCommessaConvenzione: codice della commessa o
	//della convenzione collegata alla fattura.
	CodiceCommessaConvenzione codiceCommessaConvenzione `xml:"CodiceCommessaConvenzione" json:"CodiceCommessaConvenzione"`

	//CodiceCUP: codice gestito dal CIPE che caratterizza ogni
	//progetto di investimento pubblico (Codice Unitario Progetto).
	CodiceCUP codiceCup `xml:"CodiceCUP" json:"CodiceCUP"`

	//CodiceCIG: Codice Identificativo della Gara.
	CodiceCIG codiceCig `xml:"CodiceCIG" json:"CodiceCIG"`
}

// - RiferimentoNumeroLinea: formato numerico; lunghezza massima di 4 caratteri.
// - IdDocumento: formato alfanumerico; lunghezza massima di 20 caratteri.
// - Data: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
// - NumItem: formato alfanumerico; lunghezza massima di 20 caratteri.
// - CodiceCUP: formato alfanumerico; lunghezza massima di 15 caratteri.
// - CodiceCIG: formato alfanumerico; lunghezza massima di 15 caratteri.
func (c datiFattureCollegate) Validate() error {
	var err error

	if err = c.Data.Validate(); err != nil {
		return fmt.Errorf("Data %s", err)
	}

	if err = c.CodiceCIG.Validate(); err != nil {
		return fmt.Errorf("CodiceCIG %s", err)
	}

	if err = c.CodiceCUP.Validate(); err != nil {
		return fmt.Errorf("CodiceCUP %s", err)
	}
	if err = c.RiferimentoNumeroLinea.Validate(); err != nil {
		return fmt.Errorf("RiferimentoNumeroLinea %s", err)
	}

	if err = c.NumItem.Validate(); err != nil {
		return fmt.Errorf("NumItem %s", err)
	}

	if err = c.CodiceCommessaConvenzione.Validate(); err != nil {
		return fmt.Errorf("CodiceCommessaConvenzione %s", err)
	}

	if err = c.IDDocumento.Validate(); err != nil {
		return fmt.Errorf("IDDocumento %s", err)
	}

	return err
}
