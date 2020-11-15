package einvoice

type Allegati struct {
	NomeAttachment        string `xml:"NomeAttachment" json:"NomeAttachment" validate:"max=60"`
	AlgoritmoCompressione string `xml:"AlgoritmoCompressione" json:"AlgoritmoCompressione" validate:"max=10"`
	FormatoAttachment     string `xml:"FormatoAttachment" json:"FormatoAttachment" validate:"max=10"`
	DescrizioneAttachment string `xml:"DescrizioneAttachment" json:"DescrizioneAttachment" validate:"max=100"`
	Attachment            string `xml:"Attachment" json:"Attachment" validate:"base64"`
}

type AltriDatiGestionali struct {
	TipoDato          string `xml:"TipoDato" json:"TipoDato" validate:"max=10"`
	RiferimentoTesto  string `xml:"RiferimentoTesto" json:"RiferimentoTesto" validate:"max=60"`
	RiferimentoNumero string `xml:"RiferimentoNumero" json:"RiferimentoNumero" `
	RiferimentoData   string `xml:"RiferimentoData" json:"RiferimentoData" validate:"isDate"`
}

type CodiceArticolo struct {
	CodiceTipo   string `xml:"CodiceTipo" json:"CodiceTipo" validate:"max=35"`
	CodiceValore string `xml:"CodiceValore" json:"CodiceValore" validate:"max=35"`
}

type datiAnagraficiVettore struct {
	Anagrafica   *Anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA *IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}

type datiRiepilogo struct {
	//DatiIVA        *DatiIVA `xml:"DatiIVA" json:"DatiIVA"`
	AlliquotaIVA      float64 `xml:"AliquotaIVA" json:"AliquotaIVA" validiate:"isIva"`
	Natura            string  `xml:"Natura" json:"Natura" validate:"omitempty,isNatura"`
	SpeseAccessorie   float64 `xml:"SpeseAccessorie" json:"SpeseAccessorie"`
	Arrotondamento    float64 `xml:"Arrotondamento" json:"Arrotondamento"`
	ImponibileImporto float64 `xml:"ImponibileImporto" json:"ImponibileImporto"`
	Imposta           float32 `xml:"Imposta" json:"Imposta"`

	EsigibilitaIVA       string `xml:"EsigibilitaIVA" json:"EsigibilitaIVA" validate:"omitempty,oneof=I D S"`
	RiferimentoNormativo string `xml:"RiferimentoNormativo" json:"RiferimentoNormativo"`
}

type DettaglioLinee struct {
	NumeroLinea                int                    `xml:"NumeroLinea" json:"NumeroLinea" validate:"min=1,max=9999"`
	TipoCessionePrestazione    string                 `xml:"TipoCessionePrestazione" json:"TipoCessionePrestazione" validate:"isTCP"`
	CodiceArticolo             []*CodiceArticolo      `xml:"CodiceArticolo" json:"CodiceArticolo"`
	Descrizione                string                 `xml:"Descrizione" json:"Descrizione" validate:"required,max=1000"`
	Quantita                   string                 `xml:"Quantita" json:"Quantita" validate:"max=21,min=4"`
	UnitaMisura                string                 `xml:"UnitaMisura" json:"UnitaMisura" validate:"max=10"`
	DataInizioPeriodo          string                 `xml:"DataInizioPeriodo" json:"DataInizioPeriodo" validate:"isDate"`
	DataFinePeriodo            string                 `xml:"DataFinePeriodo" json:"DataFinePeriodo" validate:"isDate"`
	PrezzoUnitario             string                 `xml:"PrezzoUnitario" json:"PrezzoUnitario" validate:"min=4,max=21"`
	ScontoMaggiorazione        []*ScontoMaggiorazione `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	PrezzoTotale               string                 `xml:"PrezzoTotale" json:"PrezzoTotale" validate:"min=4,max=21"`
	AliquotaIVA                float64                `xml:"AliquotaIVA" json:"AliquotaIVA" validate:"isIva"`
	Ritenuta                   string                 `xml:"Ritenuta" json:"Ritenuta" validate:"omitempty,eq=SI"`
	Natura                     string                 `xml:"Natura" json:"Natura" validate:"isNatura"`
	RiferimentoAmministrazione string                 `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione" validate:"max=20"`
	AltriDatiGestionali        []*AltriDatiGestionali `xml:"AltriDatiGestionali" json:"AltriDatiGestionali"`
}

type DatiBeniServizi struct {
	DettaglioLinee []*DettaglioLinee `xml:"DettaglioLinee" json:"DettaglioLinee" `
	DatiRiepilogo  []*datiRiepilogo  `xml:"DatiRiepilogo" json:"DatiRiepilogo"`
}

type DatiBollo struct {
	BolloVirtuale string `xml:"BolloVirtuale" json:"BolloVirtuale" validate:"eq=SI"`
	ImportoBollo  string `xml:"ImportoBollo" json:"ImportoBollo" validate:"max=15"`
}

type DatiCassaPrevidenziale struct {
	TipoCassa                  string  `xml:"TipoCassa" json:"TipoCassa" validate:"isTC"`
	AlCassa                    string  `xml:"AlCassa" json:"AlCassa"`
	ImportoContributoCassa     string  `xml:"ImportoContributoCassa" json:"ImportoContributoCassa" validate:"max=15"`
	ImponibileCassa            string  `xml:"ImponibileCassa" json:"ImponibileCassa" validate:"max=15"`
	AliquotaIVA                float64 `xml:"AliquotaIVA" json:"AliquotaIVA" validate:"isIva"`
	Ritenuta                   string  `xml:"Ritenuta" json:"Ritenuta" validate:"eq=SI,len=2"`
	Natura                     string  `xml:"Natura" json:"Natura" validate:"isNatura"`
	RiferimentoAmministrazione string  `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione" validate:"max=20"`
}

type DatiDocumentiCorrelatiType struct {
	RiferimentoNumeroLinea    int    `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea" validate:"min=0,max=9999"`
	IDDocumento               string `xml:"IdDocumento" json:"IdDocumento" validate:"max=20"`
	Data                      string `xml:"Data" json:"Data" validate:"omitempty,isDate"`
	NumItem                   string `xml:"NumItem" json:"NumItem" validate:""`
	CodiceCommessaConvenzione string `xml:"CodiceCommessaConvenzione" json:"CodiceCommessaConvenzione" validate:"max=100"`
	CodiceCUP                 string `xml:"CodiceCUP" json:"CodiceCUP" validate:"max=15"`
	CodiceCIG                 string `xml:"CodiceCIG" json:"CodiceCIG" validate:"max=15"`
}
type datiConvenzione = DatiDocumentiCorrelatiType
type datiOrdineAcquisto = DatiDocumentiCorrelatiType
type datiContratto = DatiDocumentiCorrelatiType
type datiFattureCollegate = DatiDocumentiCorrelatiType
type datiRicezione = DatiDocumentiCorrelatiType

type DatiDDT struct {
	NumeroDDT              string `xml:"NumeroDDT" json:"NumeroDDT" validate:"max=20"`
	DataDDT                string `xml:"DataDDT" json:"DataDTT" validate:"isDate"`
	RiferimentoNumeroLinea int    `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea" validate:"min=0,max=9999"`
}

type DatiGeneraliDocumento struct {
	TipoDocumento          string                  `xml:"TipoDocumento" json:"TipoDocumento" validate:"isTD"`
	Divisa                 string                  `xml:"Divisa" json:"Divisa" validate:"len=3"`
	Data                   string                  `xml:"Data" json:"Data" validate:"isDate,noFuture"`
	Numero                 string                  `xml:"Numero" json:"Numero"`
	DatiRitenuta           []*DatiRitenuta         `xml:"DatiRitenuta,omitempty" json:"DatiRitenuta,omitempty" validate:"omitempty"`
	DatiBollo              *DatiBollo              `xml:"DatiBollo" json:"DatiBollo" validate:"omitempty"`
	DatiCassaPrevidenziale *DatiCassaPrevidenziale `xml:"DatiCassaPrevidenziale" json:"DatiCassaPrevidenziale" `
	ScontoMaggiorazione    *ScontoMaggiorazione    `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione" `
	ImportoTotaleDocumento float64                 `xml:"ImportoTotaleDocumento" json:"ImportoTotaleDocumento" validate:""`
	Arrotondamento         float64                 `xml:"Arrotondamento" json:"Arrotondamento" validate:""`
	Causale                []string                `xml:"Causale" json:"Causale"`
	Art73                  string                  `xml:"Art73" json:"Art73" validate:"omitempty,eq=SI"`
}

type DatiGenerali struct {
	DatiGeneraliDocumento *DatiGeneraliDocumento `xml:"DatiGeneraliDocumento" json:"DatiGeneraliDocumento"`
	DatiOrdineAcquisto    *datiOrdineAcquisto    `xml:"DatiOrdineAcquisto,omitempity" json:"DatiOrdineAcquisto"`
	DatiContratto         *datiContratto         `xml:"DatiContratto" json:"DatiContratto"`
	DatiConvenzione       *datiConvenzione       `xml:"DatiConvenzione" json:"DatiConvenzione"`
	DatiRicezione         *datiRicezione         `xml:"DatiRicezione" json:"DatiRicezione"`
	DatiFattureCollegate  *datiFattureCollegate  `xml:"DatiFattureCollegate" json:"DatiFattureCollegate"`
	DatiSAL               *DatiSAL               `xml:"DatiSAL" json:"DatiSAL"`
	DatiDDT               *DatiDDT               `xml:"DatiDDT" json:"DatiDDT"`
	DatiTrasporto         *DatiTrasporto         `xml:"DatiTrasporto" json:"DatiTrasporto" validate:"omitempty"`
	FatturaPrincipale     *FatturaPrincipale     `xml:"FatturaPrincipale" json:"FatturaPrincipale"`
}

type DatiIVA struct {
	AliquotaIVA float64 `xml:"Aliquota" json:"Aliquota" validate:"isIva"`
	Imposta     float64 `xml:"Imposta" json:"Imposta" `
}

type DatiPagamento struct {
	CondizioniPagamento string              `xml:"CondizioniPagamento" json:"CondizioniPagamento" validate:"isCP"`
	DettaglioPagamento  *DettaglioPagamento `xml:"DettaglioPagamento" json:"DettaglioPagamento"`
}

type DatiRitenuta struct {
	TipoRitenuta     string  `xml:"TipoRitenuta" json:"TipoRitenuta" validate:"isTR"`
	ImportoRitenuta  float64 `xml:"ImportoRitenuta" json:"ImportoRitenuta" `
	AliquotaRitenuta float64 `xml:"AliquotaRitenuta" json:"AliquotaRitenuta"`
	CausalePagamento string  `xml:"CausalePagamento" json:"CausalePagamento" validate:"isCP"`
}

type DatiSAL struct {
	RiferimentoFase int `xml:"RiferimentoFase" json:"RiferimentoFase" validate:"min=0,max=999"`
}

type DatiTrasporto struct {
	DatiAnagraficiVettore *datiAnagraficiVettore `xml:"DatiAnagraficiVettore" json:"DatiAnagraficiVettore"`
	MezzoTrasporto        string                 `xml:"MezzoTrasporto" json:"MezzoTrasporto" validate:"omitempty,max=80"`
	CausaleTrasporto      string                 `xml:"CausaleTrasporto" json:"CausaleTrasporto" validate:"omitempty,max=100"`
	NumeroColli           int                    `xml:"NumeroColli" json:"NumeroColli" validate:"omitempty,max=9999,min=0"`
	Descrizione           string                 `xml:"Descrizione" json:"Descrizione" validate:"omitempty,max=100"`
	UnitaMisuraPeso       string                 `xml:"UnitaMisuraPeso" json:"UnitaMisuraPeso" validate:"omitempty,max=10"`
	PesoLordo             string                 `xml:"PesoLordo" json:"PesoLordo" validate:"omitempty,min=4,max=7"`
	PesoNetto             string                 `xml:"PesoNetto" json:"PesoNetto" validate:"omitempty,min=4,max=7"`
	DataOraRitiro         string                 `xml:"DataOraRitiro" json:"DataOraRitiro" validate:"omitempty,isDateTime"`
	DataInizioTrasporto   string                 `xml:"DataInizioTrasporto" json:"DataInizioTrasporto" validate:"omitempty,isDate"`
	TipoResa              string                 `xml:"TipoResa" json:"TipoResa" validate:"omitempty,len=3,oneof=EXW FCA CPT CIP DAT DAP DDP FAS FOB CFR CIF"`
	IndirizzoResa         *IndirizzoType         `xml:"IndirizzoResa" json:"IndirizzoResa" `
	DataOraConsegna       string                 `xml:"DataOraConsegna" json:"DataOraConsegna" validate:"isDateTime"`
}

type DatiVeicolo struct {
	Data           string `xml:"Data" json:"Data" validate:"isDate"`
	TotalePercorso string `xml:"TotalePercorso" json:"Tot alePercorso" validate:"max=15"`
}

type DettaglioPagamento struct {
	Beneficiario                    string  `xml:"Beneficiario" json:"Beneficiario" validate:"max=200"`
	ModalitaPagamento               string  `xml:"ModalitaPagamento" json:"ModalitaPagamento" validate:"oneof=MP01 MP02 MP03 MP04 MP05 MP06 MP07 MP08 MP09 MP10 MP11 MP12 MP13 MP14 MP15 MP16 MP17 MP18 MP19 MP20 MP21 MP22"`
	DataRiferimentoTerminiPagamento string  `xml:"DataRiferimentoTerminiPagamento" json:"DataRiferimentoTerminiPagamento" validate:"omitempty,isDate"`
	GiorniTerminiPagamento          int     `xml:"GiorniTerminiPagamento" json:"omitempty,GiorniTerminiPagamento" validate:"omitempty,min=0,max=999"`
	DataScadenzaPagamento           string  `xml:"DataScadenzaPagamento" json:"DataScadenzaPagamento" validate:"omitempty,isDate"`
	ImportoPagamento                float32 `xml:"ImportoPagamento" json:"ImportoPagamento" validate:"min=4,max=15"`
	CodUfficioPostale               string  `xml:"CodUfficioPostale" json:"CodUfficioPostale" validate:"omitempty,max=20"`
	CognomeQuietanzante             string  `xml:"CognomeQuietanzante" json:"CognomeQuietanzante" validate:"omitempty,max=60"`
	NomeQuietanzante                string  `xml:"NomeQuietanzante" json:"NomeQuietanzante" validate:"omitempty,max=60"`
	CFQuietanzante                  string  `xml:"CFQuietanzante" json:"CFQuietanzante" validate:"omitempty,max=16"`
	TitoloQuietanzante              string  `xml:"TitoloQuietanzante" json:"TitoloQuietanzante" validate:"omitempty,max=10,min=2"`
	IstitutoFinanziario             string  `xml:"IstitutoFinanziario" json:"IstitutoFinanziario" validate:"omitempty,max=80"`

	IBAN string `xml:"IBAN" json:"IBAN" validate:"omitempty,min=15,max=34,required_with_all=ABI CAB BIC"`
	ABI  string `xml:"ABI" json:"ABI" validate:"omitempty,len=5,required_with_all=IBAN CAB BIC"`
	CAB  string `xml:"CAB" json:"CAB" validate:"omitempty,len=5,required_with_all=ABI IBAN BIC"`
	BIC  string `xml:"BIC" json:"BIC" validate:"omitempty,min=8,max=11,required_with_all=ABI CAB IBAN"`

	ScontoPagamentoAnticipato     float32 `xml:"ScontoPagamentoAnticipato" json:"ScontoPagamentoAnticipato" validate:"omitempty"`
	DataLimitePagamentoAnticipato string  `xml:"DataLimitePagamentoAnticipato" json:"DataLimitePagamentoAnticipato" validate:"omitempty,isDate"`
	PenalitaPagamentiRitardati    float32 `xml:"PenalitaPagamentiRitardati" json:"PenalitaPagamentiRitardati" validate:"omitempty"`
	DataDecorrenzaPenale          string  `xml:"DataDecorrenzaPenale" json:"DataDecorrenzaPenale" validate:"omitempty,isDate"`
	CodicePagamento               string  `xml:"CodicePagamento" json:"CodicePagamento" validate:"max=60"`
}

type FatturaPrincipale struct {
	NumeroFatturaPrincipale string `xml:"NumeroFatturaPrincipale" json:"NumeroFatturaPrincipale" validate:"max=20"`
	DataFatturaPrincipale   string `xml:"DataFatturaPrincipale" json:"DataFatturaPrincipale" validate:"isDate"`
}

type ScontoMaggiorazione struct {
	Tipo        string  `xml:"Tipo" json:"Tipo" validate:"oneof=SC MG"`
	Percentuale float32 `xml:"Percentuale" json:"Percentuale"`
	Importo     float32 `xml:"Importo" json:"Importo"`
}

//FatturaElettronicaBody represents the body of an electronic invoice in all its parts
type FatturaElettronicaBody struct {
	DatiGenerali    *DatiGenerali    `xml:"DatiGenerali" json:"DatiGenerali"`
	DatiBeniServizi *DatiBeniServizi `xml:"DatiBeniServizi" json:"DatiBeniServizi"`
	DatiVeicolo     *DatiVeicolo     `xml:"DatiVeicolo" json:"DatiVeicolo"`
	DatiPagamento   *DatiPagamento   `xml:"DatiPagamento" json:"DatiPagamento"`
	Allegati        []*Allegati      `xml:"Allegati" json:"Allegati" validate:"dive"`
}

// Validate Valid the body of the invoice Italian Electronics
