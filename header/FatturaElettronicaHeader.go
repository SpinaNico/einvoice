package header

//FATTURA ELETTRONICA HEADER – Dati di trasmissione										29
//FATTURA ELETTRONICA HEADER – Dati del cedente/prestatore								30
//FATTURA ELETTRONICA HEADER – Dati del rappresentante fiscale del cedente/prestatore	33
//FATTURA ELETTRONICA HEADER – Dati del cessionario/committente 						35
//FATTURA ELETTRONICA HEADER – Dati del terzo intermediario soggetto emittente			38
//FATTURA ELETTRONICA HEADER – Soggetto emittente										39

//FatturaElettronicaHeader :
//
// ---
// SoggettoEmittente: *codice che sta ad indicare se la fattura è stata
// emessa da parte del cessionario/committente ovvero da parte di un
// terzo per conto del cedente/prestatore.*
type FatturaElettronicaHeader struct {
	DatiTrasmissione                     datiTrasmissione                     `xml:"DatiTrasmissione" json:"DatiTrasmissione"`
	CedentePrestatore                    cedentePrestatore                    `xml:"CedentePrestatore" json:"CedentePrestatore"`
	CessionarioCommittente               cessionarioCommittente               `xml:"CessionarioCommittente" json:"CessionarioCommittente"`
	TerzoIntermediarioOSoggettoEmittente terzoIntermediarioOSoggettoEmittente `xml:"TerzoIntermediarioOSoggettoEmittente" json:"TerzoIntermediarioOSoggettoEmittente"`
	SoggettoEmittente                    string                               `xml:"SoggettoEmittente" json:"SoggettoEmittente"`
}

// Validate : Check data of header
func (c FatturaElettronicaHeader) Validate() error {
	return nil
}
