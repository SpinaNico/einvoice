package invoice

import "fmt"

type allegati struct {
	NomeAttachment        string `xml:"NomeAttachment" json:"NomeAttachment" validate:"max=60"`
	AlgoritmoCompressione string `xml:"AlgoritmoCompressione" json:"AlgoritmoCompressione" validate:"max=10"`
	FormatoAttachment     string `xml:"FormatoAttachment" json:"FormatoAttachment" validate:"max=10"`
	DescrizioneAttachment string `xml:"DescrizioneAttachment" json:"DescrizioneAttachment" validate:"max=100"`
	Attachment            string `xml:"Attachment" json:"Attachment" validate:"base64"`
}

type altriDatiGestionali struct {
	TipoDato          string `xml:"TipoDato" json:"TipoDato" validate:"max=10"`
	RiferimentoTesto  string `xml:"RiferimentoTesto" json:"RiferimentoTesto" validate:"max=60"`
	RiferimentoNumero string `xml:"RiferimentoNumero" json:"RiferimentoNumero" validate:"isPrice"`
	RiferimentoData   string `xml:"RiferimentoData" json:"RiferimentoData" validate:"isDate"`
}

type codiceArticolo struct {
	CodiceTipo   string `xml:"CodiceTipo" json:"CodiceTipo" validate:"max=35"`
	CodiceValore string `xml:"CodiceValore" json:"CodiceValore" validate:"max=35"`
}

type datiAnagraficiVettore struct {
	Anagrafica   *anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA *iDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}

type datiBeniServizi struct {
	DettaglioLinee []*dettaglioLinee `xml:"DettaglioLinee" json:"DettaglioLinee"`
	DatiRiepilogo  *datiRiepilogo    `xml:"DatiRiepilogo" json:"DatiRiepilogo"`
}

type datiBollo struct {
	BolloVirtuale string `xml:"BolloVirtuale" json:"BolloVirtuale" validate:"eq=SI"`
	ImportoBollo  string `xml:"ImportoBollo" json:"ImportoBollo" validate:"isPrice,max=15"`
}

type datiCassaPrevidenziale struct {
	TipoCassa                  string `xml:"TipoCassa" json:"TipoCassa" validate:"oneof=TC01 TC02 TC03 TC04 TC05 TC06 TC07 TC08 TC09 TC10 TC11 TC12 TC13 TC14 TC15 TC16 TC17 TC18 TC19 TC20 TC21 TC22"`
	AlCassa                    string `xml:"AlCassa" json:"AlCassa" validate:"isPrice"`
	ImportoContributoCassa     string `xml:"ImportoContributoCassa" json:"ImportoContributoCassa" validate:"isPrice,max=15"`
	ImponibileCassa            string `xml:"ImponibileCassa" json:"ImponibileCassa" validate:"isPrice,max=15"`
	AliquotaIVA                string `xml:"AliquotaIVA" json:"AliquotaIVA" validate:"isPrice,max=6,min=4"`
	Ritenuta                   string `xml:"Ritenuta" json:"Ritenuta" validate:"eq=SI,len=2"`
	Natura                     string `xml:"Natura" json:"Natura" validate:"isNatura"`
	RiferimentoAmministrazione string `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione" validate:"max=20"`
}

type dataBlok struct {
	RiferimentoNumeroLinea    int    `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea" validate:"min=0,max=9999"`
	IDDocumento               string `xml:"IdDocumento" json:"IdDocumento" validate:"max=20"`
	Data                      string `xml:"Data" json:"Data" validate:"omitempty,isDate"`
	NumItem                   string `xml:"NumItem" json:"NumItem" validate:""`
	CodiceCommessaConvenzione string `xml:"CodiceCommessaConvenzione" json:"CodiceCommessaConvenzione" validate:"max=100"`
	CodiceCUP                 string `xml:"CodiceCUP" json:"CodiceCUP" validate:"max=15"`
	CodiceCIG                 string `xml:"CodiceCIG" json:"CodiceCIG" validate:"max=15"`
}
type datiConvenzione struct{ dataBlok }
type datiOrdineAcquisto struct{ dataBlok }
type datiContratto struct{ dataBlok }
type datiFattureCollegate struct{ dataBlok }
type datiRicezione struct{ dataBlok }

type datiDDT struct {
	NumeroDDT              string `xml:"NumeroDDT" json:"NumeroDDT" validate:"max=20"`
	DataDDT                string `xml:"DataDDT" json:"DataDTT" validate:"isDate"`
	RiferimentoNumeroLinea int    `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea" validate:"min=0,max=9999"`
}

type datiGeneraliDocumento struct {
	TipoDocumento          string                  `xml:"TipoDocumento" json:"TipoDocumento" validate:"oneof=TD01 TD02 TD03 TD04 TD05 TD06 TD20"`
	Divisa                 string                  `xml:"Divisa" json:"Divisa" validate:"len=3"`
	Data                   string                  `xml:"Data" json:"Data" validate:"isDate"`
	Numero                 string                  `xml:"Numero" json:"Numero"`
	DatiRitenuta           *datiRitenuta           `xml:"DatiRitenuta,omitempty" json:"DatiRitenuta,omitempty" validate:"omitempty"`
	DatiBollo              *datiBollo              `xml:"DatiBollo" json:"DatiBollo" validate:"omitempty"`
	DatiCassaPrevidenziale *datiCassaPrevidenziale `xml:"DatiCassaPrevidenziale" json:"DatiCassaPrevidenziale" `
	ScontoMaggiorazione    *scontoMaggiorazione    `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione" `
	ImportoTotaleDocumento string                  `xml:"ImportoTotaleDocumento" json:"ImportoTotaleDocumento" validate:"omitempty,isPrice,max=15,min=4"`
	Arrotondamento         string                  `xml:"Arrotondamento" json:"Arrotondamento" validate:"omitempty,isPrice,max=15,min=4"`
	Causale                []string                `xml:"Causale" json:"Causale"`
	Art73                  string                  `xml:"Art73" json:"Art73" validate:"omitempty,eq=SI"`
}

type datiGenerali struct {
	DatiGeneraliDocumento *datiGeneraliDocumento `xml:"DatiGeneraliDocumento" json:"DatiGeneraliDocumento"`
	DatiOrdineAcquisto    *datiOrdineAcquisto    `xml:"DatiOrdineAcquisto,omitempity" json:"DatiOrdineAcquisto"`
	DatiContratto         *datiContratto         `xml:"DatiContratto" json:"DatiContratto"`
	DatiConvenzione       *datiConvenzione       `xml:"DatiConvenzione" json:"DatiConvenzione"`
	DatiRicezione         *datiRicezione         `xml:"DatiRicezione" json:"DatiRicezione"`
	DatiFattureCollegate  *datiFattureCollegate  `xml:"DatiFattureCollegate" json:"DatiFattureCollegate"`
	DatiSAL               *datiSAL               `xml:"DatiSAL" json:"DatiSAL"`
	DatiDDT               *datiDDT               `xml:"DatiDDT" json:"DatiDDT"`
	DatiTrasporto         *datiTrasporto         `xml:"DatiTrasporto" json:"DatiTrasporto" validate:"omitempty"`
	FatturaPrincipale     *fatturaPrincipale     `xml:"FatturaPrincipale" json:"FatturaPrincipale"`
}

type datiIVA struct {
	AliquotaIVA string `xml:"Aliquota" json:"Aliquota" validate:"isPrice"`
	Imposta     string `xml:"Imposta" json:"Imposta" validate:"isPrice"`
}

type datiPagamento struct {
	CondizioniPagamento string              `xml:"CondizioniPagamento" json:"CondizioniPagamento" validate:"oneof=TP01 TP02 TP03"`
	DettaglioPagamento  *dettaglioPagamento `xml:"DettaglioPagamento" json:"DettaglioPagamento"`
}

type datiRiepilogo struct {
	ImponibileImporto string   `xml:"ImponibileImporto" json:"ImponibileImporto" validate:"isPrice"`
	DatiIVA           *datiIVA `xml:"DatiIVA" json:"DatiIVA"`
	Natura            string   `xml:"Natura" json:"Natura" validate:"omitempty,isNatura"`
	Detraibile        string   `xml:"Detraibile" json:"Detraibile" validate:"omitempty,isPrice,min=4,max=6"`
	Deducibile        string   `xml:"Deducibile" json:"Deducibile" validate:"omitempty,eq=SI"`
	EsigibilitaIVA    string   `xml:"EsigibilitaIVA" json:"EsigibilitaIVA" validate:"omitempty,oneof=I D S"`
}

type datiRitenuta struct {
	TipoRitenuta     string `xml:"TipoRitenuta" json:"TipoRitenuta" validate:"oneof=RT01 RT02"`
	ImportoRitenuta  string `xml:"ImportoRitenuta" json:"ImportoRitenuta" validate:"isPrice,min=4,max=15"`
	AliquotaRitenuta string `xml:"AliquotaRitenuta" json:"AliquotaRitenuta" validate:"isPrice,min=4,max=6"`
	CausalePagamento string `xml:"CausalePagamento" json:"CausalePagamento" validate:"oneof=CU ZO M2"`
}

type datiSAL struct {
	RiferimentoFase int `xml:"RiferimentoFase" json:"RiferimentoFase" validate:"min=0,max=999"`
}

type datiTrasporto struct {
	DatiAnagraficiVettore *datiAnagraficiVettore `xml:"DatiAnagraficiVettore" json:"DatiAnagraficiVettore"`
	MezzoTrasporto        string                 `xml:"MezzoTrasporto" json:"MezzoTrasporto" validate:"omitempty,max=80"`
	CausaleTrasporto      string                 `xml:"CausaleTrasporto" json:"CausaleTrasporto" validate:"omitempty,max=100"`
	NumeroColli           int                    `xml:"NumeroColli" json:"NumeroColli" validate:"omitempty,max=9999,min=0"`
	Descrizione           string                 `xml:"Descrizione" json:"Descrizione" validate:"omitempty,max=100"`
	UnitaMisuraPeso       string                 `xml:"UnitaMisuraPeso" json:"UnitaMisuraPeso" validate:"omitempty,max=10"`
	PesoLordo             string                 `xml:"PesoLordo" json:"PesoLordo" validate:"omitempty,isPrice,min=4,max=7"`
	PesoNetto             string                 `xml:"PesoNetto" json:"PesoNetto" validate:"omitempty,isPrice,min=4,max=7"`
	DataOraRitiro         string                 `xml:"DataOraRitiro" json:"DataOraRitiro" validate:"omitempty,isDateTime"`
	DataInizioTrasporto   string                 `xml:"DataInizioTrasporto" json:"DataInizioTrasporto" validate:"omitempty,isDate"`
	TipoResa              string                 `xml:"TipoResa" json:"TipoResa" validate:"omitempty,len=3,oneof=EXW FCA CPT CIP DAT DAP DDP FAS FOB CFR CIF"`
	IndirizzoResa         *indirizzoType         `xml:"IndirizzoResa" json:"IndirizzoResa" `
	DataOraConsegna       string                 `xml:"DataOraConsegna" json:"DataOraConsegna" validate:"isDateTime"`
}

type datiVeicolo struct {
	Data           string `xml:"Data" json:"Data" validate:"isDate"`
	TotalePercorso string `xml:"TotalePercorso" json:"TotalePercorso" validate:"max=15"`
}

type dettaglioLinee struct {
	NumeroLinea                int                  `xml:"NumeroLinea" json:"NumeroLinea" validate:"min=1,max=9999"`
	TipoCessionePrestazione    string               `xml:"TipoCessionePrestazione" json:"TipoCessionePrestazione" validate:"oneof=SC PR AB AC"`
	CodiceArticolo             *codiceArticolo      `xml:"CodiceArticolo" json:"CodiceArticolo"`
	Descrizione                string               `xml:"Descrizione" json:"Descrizione" validate:"max=1000"`
	Quantita                   string               `xml:"Quantita" json:"Quantita" validate:"isPrice,max=21,min=4"`
	UnitaMisura                string               `xml:"UnitaMisura" json:"UnitaMisura" validate:"max=10"`
	DataInizioPeriodo          string               `xml:"DataInizioPeriodo" json:"DataInizioPeriodo" validate:"isDate"`
	DataFinePeriodo            string               `xml:"DataFinePeriodo" json:"DataFinePeriodo" validate:"isDate"`
	PrezzoUnitario             string               `xml:"PrezzoUnitario" json:"PrezzoUnitario" validate:"isPrice,min=4,max=21"`
	ScontoMaggiorazione        *scontoMaggiorazione `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	PrezzoTotale               string               `xml:"PrezzoTotale" json:"PrezzoTotale" validate:"isPrice,min=4,max=21"`
	AliquotaIVA                string               `xml:"AliquotaIVA" json:"AliquotaIVA" validate:"isPrice,min=4,max=6"`
	Ritenuta                   string               `xml:"Ritenuta" json:"Ritenuta" validate:"omitempty,eq=SI"`
	Natura                     string               `xml:"Natura" json:"Natura" validate:"isNatura"`
	RiferimentoAmministrazione string               `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione" validate:"max=20"`
	AltriDatiGestionali        *altriDatiGestionali `xml:"AltriDatiGestionali" json:"AltriDatiGestionali"`
}

type dettaglioPagamento struct {
	Beneficiario                    string `xml:"Beneficiario" json:"Beneficiario" validate:"max=200"`
	ModalitaPagamento               string `xml:"ModalitaPagamento" json:"ModalitaPagamento" validate:"oneof=MP01 MP02 MP03 MP04 MP05 MP06 MP07 MP08 MP09 MP10 MP11 MP12 MP13 MP14 MP15 MP16 MP17 MP18 MP19 MP20 MP21 MP22"`
	DataRiferimentoTerminiPagamento string `xml:"DataRiferimentoTerminiPagamento" json:"DataRiferimentoTerminiPagamento" validate:"omitempty,isDate"`
	GiorniTerminiPagamento          int    `xml:"GiorniTerminiPagamento" json:"omitempty,GiorniTerminiPagamento" validate:"omitempty,min=0,max=999"`
	DataScadenzaPagamento           string `xml:"DataScadenzaPagamento" json:"DataScadenzaPagamento" validate:"omitempty,isDate"`
	ImportoPagamento                string `xml:"ImportoPagamento" json:"ImportoPagamento" validate:"isPrice,min=4,max=15"`
	CodUfficioPostale               string `xml:"CodUfficioPostale" json:"CodUfficioPostale" validate:"omitempty,max=20"`
	CFQuietanzante                  string `xml:"CFQuietanzante" json:"CFQuietanzante" validate:"omitempty,max=16"`
	NomeQuietanzante                string `xml:"NomeQuietanzante" json:"NomeQuietanzante" validate:"omitempty,max=60"`
	CognomeQuietanzante             string `xml:"CognomeQuietanzante" json:"CognomeQuietanzante" validate:"omitempty,max=60"`
	TitoloQuietanzante              string `xml:"TitoloQuietanzante" json:"TitoloQuietanzante" validate:"omitempty,max=10,min=2"`
	IstitutoFinanziario             string `xml:"IstitutoFinanziario" json:"IstitutoFinanziario" validate:"omitempty,max=80"`

	IBAN string `xml:"IBAN" json:"IBAN" validate:"omitempty,min=15,max=34,required_with_all=ABI CAB BIC"`
	ABI  int    `xml:"ABI" json:"ABI" validate:"omitempty,len=5,required_with_all=IBAN CAB BIC"`
	CAB  string `xml:"CAB" json:"CAB" validate:"omitempty,len=5,required_with_all=ABI IBAN BIC"`
	BIC  string `xml:"BIC" json:"BIC" validate:"omitempty,min=8,max=11,required_with_all=ABI CAB IBAN"`

	ScontoPagamentoAnticipato     string `xml:"ScontoPagamentoAnticipato" json:"ScontoPagamentoAnticipato" validate:"omitempty,isPrice,min=4,max=15"`
	DataLimitePagamentoAnticipato string `xml:"DataLimitePagamentoAnticipato" json:"DataLimitePagamentoAnticipato" validate:"omitempty,isDate"`
	PenalitaPagamentiRitardati    string `xml:"PenalitaPagamentiRitardati" json:"PenalitaPagamentiRitardati" validate:"omitempty,isPrice,min=4,max=15"`
	DataDecorrenzaPenale          string `xml:"DataDecorrenzaPenale" json:"DataDecorrenzaPenale" validate:"omitempty,isDate"`
	CodicePagamento               string `xml:"CodicePagamento" json:"CodicePagamento" validate:"max=60"`
}

type fatturaPrincipale struct {
	NumeroFatturaPrincipale string `xml:"NumeroFatturaPrincipale" json:"NumeroFatturaPrincipale" validate:"max=20"`
	DataFatturaPrincipale   string `xml:"DataFatturaPrincipale" json:"DataFatturaPrincipale" validate:"isDate"`
}

type scontoMaggiorazione struct {
	Tipo        string  `xml:"Tipo" json:"Tipo"`
	Percentuale float32 `xml:"Percentuale" json:"Percentuale"`
	Importo     float32 `xml:"Importo" json:"Importo"`
}

//FatturaElettronicaBody represents the body of an electronic invoice in all its parts
type FatturaElettronicaBody struct {
	DatiGenerali    *datiGenerali    `xml:"DatiGenerali" json:"DatiGenerali"`
	DatiBeniServizi *datiBeniServizi `xml:"DatiBeniServizi" json:"DatiBeniServizi"`
	DatiVeicolo     *datiVeicolo     `xml:"DatiVeicolo" json:"DatiVeicolo"`
	DatiPagamento   *datiPagamento   `xml:"DatiPagamento" json:"DatiPagamento"`
	Allegati        []*allegati      `xml:"Allegati" json:"Allegati" validate:"dive"`
}

// Validate Valid the body of the invoice Italian Electronics
func (v FatturaElettronicaBody) Validate() error {
	validate := getValidator()
	if err := validate.Struct(v); err != nil {
		return fmt.Errorf("FatturaElettronicaBody %s", err)
	}
	return nil
}
