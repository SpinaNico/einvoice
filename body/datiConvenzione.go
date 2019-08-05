package body

//DatiConvenzione (dati relativi alla convenzione collegata alla fattura)
type datiConvenzione struct {
	//	RiferimentoNumeroLinea: numero della linea o delle linee di
	//dettaglio della fattura alle quali si riferisce la convenzione così
	//come identificata dai tre elementi successivi (IdDocumento,
	//Data, NumItem); nel caso in cui la convenzione si riferisce
	//all’intera fattura, questo elemento non deve essere
	//valorizzato.
	RiferimentoNumeroLinea string `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea"`

	//IdDocumento: numero della convenzione associata alla
	//fattura o alla linea/linee di fattura indicate nell’elemento
	//RiferimentoNumeroLinea.
	IDDocumento string `xml:"IdDocumento" json:"IdDocumento"`

	//Data: data della convenzione associata alla fattura o alla
	//linea/linee di fattura indicate  nell’elemento RiferimentoNumeroLinea.
	Data string `xml:"Data" json:"Data"`

	//NumItem: identificativo della singola voce (linea di
	//convenzione) all'interno della convenzione associata alla
	//fattura o alla linea/linee di fattura indicate nell’elemento RiferimentoNumeroLinea.
	NumItem string `xml:"NumItem" json:"NumItem"`

	//CodiceCommessaConvenzione: codice della commessa o
	//della convenzione collegata alla fattura.
	CodiceCommessaConvenzione string `xml:"CodiceCommessaConvenzione" json:"CodiceCommessaConvenzione"`

	//CodiceCUP: codice gestito dal CIPE che caratterizza ogni
	//progetto di investimento pubblico (Codice Unitario Progetto).
	CodiceCUP string `xml:"CodiceCUP" json:"CodiceCUP"`

	//CodiceCIG: Codice Identificativo della Gara.
	CodiceCIG string `xml:"CodiceCIG" json:"CodiceCIG"`
}

// - RiferimentoNumeroLinea: formato numerico; lunghezza massima di 4 caratteri.
// - IdDocumento: formato alfanumerico; lunghezza massima di 20 caratteri.
// - Data: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
// - NumItem: formato alfanumerico; lunghezza massima di 20 caratteri.
// - CodiceCUP: formato alfanumerico; lunghezza massima di 15 caratteri.
// - CodiceCIG: formato alfanumerico; lunghezza massima di 15 caratteri.
func (c datiConvenzione) Validate() error {
	return nil
}
