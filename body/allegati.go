package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
)

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
func (c allegati) Validate() error {

	if len(c.NomeAttachment) > 60 {
		return fmt.Errorf("Allegati (NomeAttachment) %s", share.ErrorMaxLength(60))
	}

	if len(c.FormatoAttachment) > 10 {
		return fmt.Errorf("Allegati (FormatoAttachment) %s", share.ErrorMaxLength(10))
	}

	if len(c.DescrizioneAttachment) > 100 {
		return fmt.Errorf("Allegati (DescrizioneAttachment) %s", share.ErrorMaxLength(100))
	}
	// Attachment test?
	return nil
}
