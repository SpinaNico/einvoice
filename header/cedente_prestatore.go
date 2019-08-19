package header

import (
	share "github.com/SpinaNico/go-struct-invoice/share"
)

type cedentePrestatore struct {
	DatiAnagrafici datiAnagrafici      `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           share.IndirizzoType `xml:"Sede" json:"Sede"`
	Contatti       contatti            `xml:"Contatti" json:"Contatti"`
}
