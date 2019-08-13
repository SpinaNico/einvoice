package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiGeneraliDocumento struct {
	//TipoDocumento: tipologia del documento oggetto della trasmissione
	//(fattura, acconto/anticipo su fattura, acconto/anticipo su parcella ,
	// nota di credito, nota di debito, parcella, autofattura).
	// + TD01 Fattura
	// + TD02 Acconto/Anticipo su fattura
	// + TD03 Acconto/Anticipo su parcella
	// + TD04 Nota di Credito
	// + TD05 Nota di Debito
	// + TD06 Parcella
	// + TD20 Autofattura
	TipoDocumento tipoDocumento `xml:"TipoDocumento" json:"TipoDocumento"`

	// Divisa: tipo di valuta utilizzata per l'indicazione degli importi.
	Divisa divisa `xml:"Divisa" json:"Divisa"`

	// Data: data del documento.
	Data data `xml:"Data" json:"Data"`

	// Numero: numero progressivo attribuito dal cedente/prestatore
	// al documento. Deve contenere almeno un carattere numerico.
	// In caso contrario il file viene scartato con codice errore 00425.
	Numero string `xml:"Numero" json:"Numero"`
	// Campo che viene valorizzato in caso sia presente ritenuta.
	DatiRitenuta           datiRitenuta           `xml:"DatiRitenuta" json:"DatiRitenuta"`
	DatiBollo              datiBollo              `xml:"DatiBollo" json:"DatiBollo"`
	DatiCassaPrevidenziale datiCassaPrevidenziale `xml:"DatiCassaPrevidenziale" json:"DatiCassaPrevidenziale "`
	ScontoMaggiorazione    scontoMaggiorazione    `xml:"ScontoMaggiorazione" json:"ScontoMaggiorazione"`
	// max length 15 characters and minium length 4 (00.00)
	ImportoTotaleDocumento decimale2 `xml:"ImportoTotaleDocumento" json:"ImportoTotaleDocumento"`
	// max length 15 characters and minium length 4 (00.00)
	Arrotondamento decimale2 `xml:"Arrotondamento" json:"Arrotondamento"`
	// Causale: descrizione della causale del documento.
	Causale []string `xml:"Causale" json:"Causale"`
	//Art73: formato alfanumerico; lunghezza di 2 caratteri;
	// il valore ammesso è:
	// + SI
	// ---
	// documento emesso secondo modalità e termini stabiliti con DM ai sensi del’’art. 73 del DPR 633/72.
	Art73 string `xml:"Art73" json:"Art73"`
}

// Validate ...
func (f datiGeneraliDocumento) Validate() error {
	var err error

	if err = f.TipoDocumento.Validate(); err != nil {
		return fmt.Errorf("TipoDocumento %s", err)
	}
	if err = f.Divisa.Validate(); err != nil {
		return fmt.Errorf("Divisa %s", err)
	}
	if err = f.Data.Validate(); err != nil {
		return fmt.Errorf("Data %s", err)
	}
	if err = f.DatiRitenuta.Validate(); err != nil {
		return fmt.Errorf("DatiRitenuta %s", err)
	}
	if err = f.DatiBollo.Validate(); err != nil {
		return fmt.Errorf("DatiBollo %s", err)
	}
	if err = f.DatiCassaPrevidenziale.Validate(); err != nil {
		return fmt.Errorf("DatiCassaPrevidenziale %s", err)
	}
	if err = f.ScontoMaggiorazione.Validate(); err != nil {
		return fmt.Errorf("ScontoMaggiorazione %s", err)
	}
	if err = f.ImportoTotaleDocumento.Validate(); err != nil {
		return fmt.Errorf("ImportoTotaleDocumento %s", err)
	}
	if err = f.Arrotondamento.Validate(); err != nil {
		return fmt.Errorf("Arrotondamento %s", err)
	}
	// Causale
	// Art73
	// Numero
	if len(f.Numero) > 20 {
		return fmt.Errorf("Numero %s", share.ErrorMaxLength(20))
	}
	if f.Art73 != "SI" {
		return fmt.Errorf("Art73 %s", share.ErrorIncorrectValue(f.Art73))
	}
	n := 0
	for _, val := range f.Causale {
		n = n + len(val)
	}
	if n > 200 {
		return fmt.Errorf("Casuale %s", share.ErrorMaxLength(200))
	}
	return err
}
