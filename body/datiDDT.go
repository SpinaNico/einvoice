package body

//2.2.9.4 Dati Ddt (Documento Di Trasporto)
// DatiDDT (nei casi in cui sia presente un documento di trasporto
//collegato alla fattura, casi di fatturazione differita, vanno valorizzati i
//seguenti elementi per ogni documento di trasporto)
type datiDDT struct {
	//NumeroDDT: numero del Documento Di Trasporto.
	NumeroDDT int `xml:"NumeroDDT" json:"NumeroDDT"`

	//DataDDT: data del Documento Di Trasporto
	DataDDT data `xml:"DataDDT" json:"DataDTT"`

	//RiferimentoNumeroLinea: numero della linea o delle linee di
	//dettaglio della fattura alle quali si riferisce il DDT (così come
	//identificato dagli elementi NumeroDDT e DataDDT); nel caso
	//in cui il documento di trasporto si riferisce all’intera fattura,
	//questo elemento non deve essere valorizzato.
	RiferimentoNumeroLinea int `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea"`
}

func (c datiDDT) Validate() error {
	return nil
}
