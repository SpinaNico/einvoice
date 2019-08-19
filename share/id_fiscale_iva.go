package share

import (
	"fmt"
)

// IDFiscaleIVA :
// **IT** 	Questa struttura definisce il luogo della partita iva, richiede il codice
//	 		paese (example: "it") e la partita iva (exmaple: 00000000)
// **EN** 	This structure defines the location of the VAT number,
// 			requires the country code (example: "it") and the VAT number (exmaple: 00000000)
type IDFiscaleIVA struct {
	// IDPaese: deve essere espresso secondo lo standard
	// ISO 3166-1 alpha-2 è utilizzato per indicare i luoghi.
	// esempio: Italy 		->	IT
	// esempio: Australia 	-> 	AU
	// esempio: Grece 		->	GR
	// maggiori informazioni: [wikipedia](https://it.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	IDPaese string `xml:"IdPaese" json:"IdPaese" validate:"omitempty,len=2"`
	// IDCodice : Lunghezza massima 28 caratteri alfanumerici.
	IDCodice string `xml:"IdCodice" json:"IdCodice" validate:"omitempty,max=28"`
}

// Validate :
// **IT** valida la struttura, IDFiscaleIVA
// **EN** valid the struct type IDFiscaleIva
func (a IDFiscaleIVA) Validate() error {
	// ....
	if len(a.IDCodice) > 28 {
		return fmt.Errorf("IDFiscaleIVA, (IdCodice): %s", ErrorMaxLength(28))
	}
	if len(a.IDPaese) != 2 {
		return fmt.Errorf("IDPaese %s", ErrorEgual(2))
	}
	return nil
}
