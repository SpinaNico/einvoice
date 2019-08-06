package body

import (
	"fmt"
)

// DatiContratto (dati relativi al contratto dal quale scaturisce la
// cessione/prestazione oggetto del documento fattura)
type datiContratto struct {
	//	RiferimentoNumeroLinea: numero della linea o delle linee di
	//dettaglio della fattura alle quali si riferisce il contratto così
	//come identificato dai tre elementi successivi (IdDocumento,
	//Data, NumItem); nel caso in cui il contratto si riferisce all’intera
	//fattura, questo elemento non deve essere valorizzato.
	RiferimentoNumeroLinea numeroLinea `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea"`

	//IdDocumento: numero del contratto associato alla fattura o
	//allalinea/linee di fattura indicate nell’elemento RiferimentoNumeroLinea
	IDDocumento idDocumento `xml:"IdDocumento" json:"IdDocumento"`

	//Data: data del contratto associato alla fattura o alla linea/linee
	//di fattura indicate nell’elemento RiferimentoNumeroLinea.
	Data data `xml:"Data" json:"Data"`

	// NumItem: identificativo della singola voce (linea di contratto)
	//all'interno del contratto associata alla fattura o alla linea/linee
	//di fattura indicate nell’elemento RiferimentoNumeroLinea.
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

func (c datiContratto) Validate() error {
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
