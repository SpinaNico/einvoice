package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type dettaglioPagamento struct {
	//Beneficiario: formato alfanumerico; lunghezza massima di 200 caratteri.
	Beneficiario string `xml:"Beneficiario" json:"Beneficiario"`
	// ModalitaPagamento: formato alfanumerico; lunghezza di 4
	// caratteri; i valori ammessi sono i seguenti:
	// ---
	// + MP01 contanti
	// + MP02 assegno
	// + MP03 assegno circolare
	// + MP04 contanti presso Tesoreria
	// + MP05 bonifico
	// + MP06 vaglia cambiario
	// + MP07 bollettino bancario
	// + MP08 carta di pagamento
	// + MP09 RID
	// + MP10 RID utenze
	// + MP11 RID veloce
	// + MP12 Riba
	// + MP13 MAV
	// + MP14 quietanza erario stato
	// + MP15 giroconto su conti di contabilità speciale
	// + MP16 domiciliazione bancaria
	// + MP17 domiciliazione postale
	// + MP18 bollettino di c/c postale
	// + MP19 SEPA Direct Debit
	// + MP20 SEPA Direct Debit CORE
	// + MP21 SEPA Direct Debit B2B
	// + MP22 Trattenuta su somme già riscosse
	ModalitaPagamento               string `xml:"ModalitaPagamento" json:"ModalitaPagamento"`
	DataRiferimentoTerminiPagamento data   `xml:"DataRiferimentoTerminiPagamento" json:"DataRiferimentoTerminiPagamento"`
	// GiorniTerminiPagamento: formato numerico di lunghezza
	// massima pari a 3. Vale 0 (zero) per pagamenti a vista.
	GiorniTerminiPagamento int  `xml:"GiorniTerminiPagamento" json:"GiorniTerminiPagamento"`
	DataScadenzaPagamento  data `xml:"DataScadenzaPagamento" json:"DataScadenzaPagamento"`
	// lunghezza va da 4 a 15 caratteri.
	ImportoPagamento decimale2 `xml:"ImportoPagamento" json:"ImportoPagamento"`
	// formato alfanumerico; lunghezza massima di 20 caratteri.
	CodUfficioPostale string `xml:"CodUfficioPostale" json:"CodUfficioPostale"`
	// CFQuietanzante: formato alfanumerico; lunghezza di 16 caratteri.
	CFQuietanzante string `xml:"CFQuietanzante" json:"CFQuietanzante"`
	// NomeQuietanzante: formato alfanumerico; lunghezza massima di 60 caratteri.
	NomeQuietanzante string `xml:"NomeQuietanzante" json:"NomeQuietanzante"`
	// CognomeQuietanzante: formato alfanumerico; lunghezza massima di 60 caratteri.
	CognomeQuietanzante string `xml:"CognomeQuietanzante" json:"CognomeQuietanzante"`
	// formato alfanumerico; lunghezza che va da 2 a 10 caratteri. 95
	TitoloQuietanzante string `xml:"TitoloQuietanzante" json:"TitoloQuietanzante"`
	// formato alfanumerico; lunghezza massima di 80 caratteri.
	IstitutoFinanziario string `xml:"IstitutoFinanziario" json:"IstitutoFinanziario"`
	// formato alfanumerico; lunghezza che va da 15 a 34 caratteri.
	IBAN string `xml:"IBAN" json:"IBAN"`
	//formato numerico di 5 caratteri.
	ABI string `xml:"ABI" json:"ABI"`
	// formato numerico di 5 caratteri.
	CAB string `xml:"CAB" json:"CAB"`
	//formato alfanumerico; lunghezza che va da 8 a 11 caratteri.
	BIC string `xml:"BIC" json:"BIC"`
	// lunghezza va da 4 a 15 caratteri.
	ScontoPagamentoAnticipato     decimale2 `xml:"ScontoPagamentoAnticipato" json:"ScontoPagamentoAnticipato"`
	DataLimitePagamentoAnticipato data      `xml:"DataLimitePagamentoAnticipato" json:"DataLimitePagamentoAnticipato"`
	// lunghezza va da 4 a 15 caratteri.
	PenalitaPagamentiRitardati decimale2 `xml:"PenalitaPagamentiRitardati" json:"PenalitaPagamentiRitardati"`
	DataDecorrenzaPenale       data      `xml:"DataDecorrenzaPenale" json:"DataDecorrenzaPenale"`
	//CodicePagamento: formato alfanumerico; lunghezza massima di 60 caratteri.
	CodicePagamento string `xml:"CodicePagamento" json:"CodicePagamento"`
}

// Validate ...
func (f dettaglioPagamento) Validate() error {
	// bene
	if len(f.Beneficiario) > 200 {
		return fmt.Errorf("Beneficiario %s", share.ErrorMaxLength(200))
	}

	// mp
	mpslice := share.IotaString2zerofill("MP", 1, 22)
	if _, ok := mpslice[f.ModalitaPagamento]; ok == false {
		return fmt.Errorf("ModalitaPagamento %s", share.ErrorIncorrectValue(f.ModalitaPagamento))
	}
	//dati term
	if err := f.DataRiferimentoTerminiPagamento.Validate(); err != nil {
		return fmt.Errorf("DataRiferimentoTerminiPagamento %s", err)
	}

	// giorni
	if f.GiorniTerminiPagamento > 999 || f.GiorniTerminiPagamento < 0 {
		return fmt.Errorf("GiorniTerminiPagamento %s", share.ErrorIncluded(0, 999))
	}

	//data scade
	if err := f.DataScadenzaPagamento.Validate(); err != nil {
		return fmt.Errorf("DataScadenzaPagamento %s", err)
	}
	// importo
	if err := f.ImportoPagamento.Validate(); err != nil {
		return fmt.Errorf("ImportoPagamento %s", err)
	}
	if len(string(f.ImportoPagamento)) > 15 || len(string(f.ImportoPagamento)) < 4 {
		return fmt.Errorf("ImportoPagamento %s", share.ErrorIncluded(4, 15))
	}
	// cod
	if len(f.CodUfficioPostale) > 20 {
		return fmt.Errorf("CodUfficioPostale %s", share.ErrorMaxLength(20))
	}
	// cognome
	if len(f.CognomeQuietanzante) > 60 {
		return fmt.Errorf("CognomeQuietanzante %s", share.ErrorMaxLength(60))
	}
	// nome
	if len(f.NomeQuietanzante) > 60 {
		return fmt.Errorf("NomeQuietanzante %s", share.ErrorMaxLength(60))
	}
	// cf
	if len(f.CFQuietanzante) != 16 {
		return fmt.Errorf("CFQuietanzante %s", share.ErrorEgual(16))
	}
	// titolo
	if len(f.TitoloQuietanzante) < 2 || len(f.TitoloQuietanzante) > 10 {
		return fmt.Errorf("TitoloQuietanzante %s", share.ErrorIncluded(2, 10))
	}
	// istituto
	if len(f.IstitutoFinanziario) > 80 {
		return fmt.Errorf("IstitutoFinanziario %s", share.ErrorMaxLength(80))
	}
	// iban
	if len(f.IBAN) < 15 || len(f.IBAN) > 34 {
		return fmt.Errorf("IBAN %s", share.ErrorIncluded(15, 34))
	}
	// abi
	if len(f.ABI) != 5 {
		return fmt.Errorf("ABI %s", share.ErrorEgual(5))
	}
	// cab
	if len(f.CAB) != 5 {
		return fmt.Errorf("CAB %s", share.ErrorEgual(5))
	}
	// bic
	if len(f.BIC) < 8 || len(f.BIC) > 11 {
		return fmt.Errorf("BIC %s", share.ErrorIncluded(8, 11))
	}

	// sconto
	if err := f.ScontoPagamentoAnticipato.Validate(); err != nil {
		return fmt.Errorf("ScontoPagamentoAnticipato %s", err)
	}
	if len(string(f.ScontoPagamentoAnticipato)) > 15 || len(string(f.ScontoPagamentoAnticipato)) < 4 {
		return fmt.Errorf("ScontoPagamentoAnticipato %s", share.ErrorIncluded(4, 15))
	}
	//DataLimitePagamentoAnticipato
	if err := f.DataLimitePagamentoAnticipato.Validate(); err != nil {
		return fmt.Errorf("DataLimitePagamentoAnticipato %s", err)
	}

	// pena
	if err := f.PenalitaPagamentiRitardati.Validate(); err != nil {
		return fmt.Errorf("PenalitaPagamentiRitardati %s", err)
	}

	if len(string(f.PenalitaPagamentiRitardati)) > 15 || len(string(f.PenalitaPagamentiRitardati)) < 4 {
		return fmt.Errorf("PenalitaPagamentiRitardati %s", share.ErrorIncluded(4, 15))
	}
	//DataDecorrenzaPenale
	if err := f.DataDecorrenzaPenale.Validate(); err != nil {
		return fmt.Errorf("DataDecorrenzaPenale %s", err)
	}

	if len(f.CodicePagamento) > 60 {
		return fmt.Errorf("CodicePagamento %s", share.ErrorMaxLength(60))
	}

	return nil
}
