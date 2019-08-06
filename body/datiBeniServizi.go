package body

import "fmt"

type datiBeniServizi struct {
	DettaglioLinee []dettaglioLinee `xml:"DettaglioLinee" json:"DettaglioLinee"`
	DatiRiepilogo  datiRiepilogo    `xml:"DatiRiepilogo" json:"DatiRiepilogo"`
}

// Validate ...
func (f datiBeniServizi) Validate() error {
	var err error

	if err = f.DatiRiepilogo.Validate(); err != nil {
		return fmt.Errorf("DatiBeniServizi %s", err)
	}

	for index, value := range f.DettaglioLinee {
		if err = value.Validate(); err != nil {
			return fmt.Errorf("DatiBeniServizi DettaglioLinee (%d) %s", index, err)
		}
	}

	return nil
}
