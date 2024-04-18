package sdi

const HeaderXMLInvoice = `<?xml version="1.0" encoding="UTF-8"?>`

type FormatoTrasmissione string

const FPR12 = FormatoTrasmissione("FPR12")
const FPA12 = FormatoTrasmissione("FPA12")

type TipoSocieta string

const (
	SU = "SU"
	SM = "SM"
)

var TipiSocieta = map[TipoSocieta]string{
	SU: "La società è a socio unico",
	SM: "La società non è a socio unico",
}

type StatoLiquidazione string

const (
	LS = "LS"
	LN = "LN"
)

var StatiLiquidazione = map[StatoLiquidazione]string{
	LS: "La società è in stato di liquidazione",
	LN: "La società non è in stato di liquidazione",
}

type SoggettoEmittente string

const (
	CC = "CC"
	TZ = "TZ"
)

var SoggettiEmittente = map[SoggettoEmittente]string{
	CC: "Cessionario/committente",
	TZ: "Soggetto terzo",
}

type ScontoMaggiorazioneType string

const (
	Sconto        ScontoMaggiorazioneType = "SC"
	Maggiorazione ScontoMaggiorazioneType = "MG"
)

var ScontiMaggiorazione = map[ScontoMaggiorazioneType]string{
	Sconto:        "Sconto",
	Maggiorazione: "Maggiorazione",
}

type EsigibilitaIVA string

const (
	Immediata EsigibilitaIVA = "I"
	Differita EsigibilitaIVA = "D"
	Scissione EsigibilitaIVA = "S"
)

var VarieEsigibilitaIVA = map[EsigibilitaIVA]string{
	Immediata: "IVA ad esigibilità immediata",
	Differita: "IVA ad esigibilità differita",
	Scissione: "scissione dei pagamenti",
}

type ValidatorSDI interface {
	Validate() []*SDIError
}
