package body

type allegati struct {
	// NomeAttachment: formato alfanumerico; lunghezza massima di 60 caratteri.
	NomeAttachment string `xml:"NomeAttachment" json:"NomeAttachment"`
	// AlgoritmoCompressione: formato alfanumerico; lunghezza massima di 10 caratteri.
	AlgoritmoCompressione string `xml:"AlgoritmoCompressione" json:"AlgoritmoCompressione"`
	// FormatoAttachment: formato alfanumerico; lunghezza massima di 10 caratteri.
	FormatoAttachment string `xml:"FormatoAttachment" json:"FormatoAttachment"`
	// DescrizioneAttachment: formato alfanumerico; lunghezza massima di 100 caratteri.
	DescrizioneAttachment string `xml:"DescrizioneAttachment" json:"DescrizioneAttachment"`
	// Attachment: è in formato xs:base64Binary.
	Attachment string `xml:"Attachment" json:"Attachment"`
}

// Validate ...
func (f allegati) Validate() error {
	return nil
}
