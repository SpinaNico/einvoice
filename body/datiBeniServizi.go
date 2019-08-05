package body

type datiBeniServizi struct {
	DettaglioLinee []dettaglioLinee `xml:"DettaglioLinee" json:"DettaglioLinee"`
	DatiRiepilogo  datiRiepilogo    `xml:"DatiRiepilogo" json:"DatiRiepilogo"`
}

// Validate ...
func (f datiBeniServizi) Validate() error {
	return nil
}
