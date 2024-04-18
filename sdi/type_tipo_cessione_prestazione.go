package sdi

type TipoCessionePrestazione string

const SC TipoCessionePrestazione = "SC"
const PR TipoCessionePrestazione = "PR"
const AB TipoCessionePrestazione = "AB"
const AC TipoCessionePrestazione = "AC"

var TipiCessionePrestazione = map[string]string{
	"SC": "Sconto",
	"PR": "Premio",
	"AB": "Abbuono",
	"AC": "Spesa accessoria",
}
