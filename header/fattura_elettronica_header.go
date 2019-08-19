package header

//FatturaElettronicaHeader :
type FatturaElettronicaHeader struct {
	/// NO DOC
	DatiTrasmissione datiTrasmissione `xml:"DatiTrasmissione" json:"DatiTrasmissione"`
	/// NO DOC
	CedentePrestatore cedentePrestatore `xml:"CedentePrestatore" json:"CedentePrestatore"`
	/// NO DOC
	RappresentanteFiscale rappresentanteFiscale `xml:"RappresentanteFiscale" json:"RappresentanteFiscale"`
	/// NO DOC
	CessionarioCommittente cessionarioCommittente `xml:"CessionarioCommittente" json:"CessionarioCommittente"`

	/// NO DOC
	TerzoIntermediarioOSoggettoEmittente terzoIntermediarioOSoggettoEmittente `xml:"TerzoIntermediarioOSoggettoEmittente" json:"TerzoIntermediarioOSoggettoEmittente"`

	// Nei casi di documenti emessi da un soggetto diverso dal cedente/prestatore va
	// valorizzato l’elemento seguente.
	// SoggettoEmittente: codice che sta ad indicare se la fattura è stata
	// emessa da parte del cessionario/committente ovvero da parte di un
	// terzo per conto del cedente/prestatore.
	// può assumere il Valore: "CC" o "TZ"
	SoggettoEmittente string `xml:"SoggettoEmittente" json:"SoggettoEmittente" validate:"omitempty,len=2,oneof=CC CZ"`
}
