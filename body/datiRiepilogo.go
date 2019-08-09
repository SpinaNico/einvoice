package body

import (
	"fmt"
	"strings"

	"github.com/SpinaNico/go-struct-invoice/share"
)

// DatiRiepilogo (blocco obbligatorio che può ripetersi più volte
// per ogni fattura fino a un massimo di 1000 occorrenze. Ogni
// ripetizione conterrà le informazioni aggregate per ciascun
// valore di aliquota IVA applicata alle operazioni elencate nel
// documento e, nel caso di imposta a zero, per ciascun motivo
// di esclusione, come riportato nell'elemento Natura.
type datiRiepilogo struct {
	// ImponibileImporto: ammontare dei beni ceduti e dei
	// servizi resi. Nei casi di documento (fattura/nota di
	// credito/nota di debito) ordinario contiene:
	//  + o la base imponibile alla quale applicare l’IVA secondo l’aliquota indicata.
	//  +  o l’importo (per le operazioni per le quali il
	// cedente/prestatore non deve dettagliare l'imposta in fattura).
	// Per le fatture semplificate, contiene l'importo risultante
	// dalla somma di imponibile ed imposta.
	ImponibileImporto decimale2 `xml:"ImponibileImporto" json:"ImponibileImporto"`

	DatiIVA datiIVA `xml:"DatiIVA" json:"DatiIVA"`
	Natura  natura  `xml:"Natura" json:"Natura"`

	//Detraibile : La sua lunghezzava da 4 a 6 caratteri.
	Detraibile decimale2 `xml:"Detraibile" json:"Detraibile"`
	//Deducibile: formato alfanumerico; lunghezza di 2 caratteri; i
	// valori ammessi sono i seguenti:
	// + SI spesa deducibile.
	Deducibile string `xml:"Deducibile" json:"Deducibile"`

	// EsigibilitaIVA: formato alfanumerico; lunghezza di 1 carattere;
	// i valori ammessi sono i seguenti:
	//
	// + I IVA ad esigibilità immediata
	// + D	IVA ad esigibilità differita
	// + S	scissione dei pagamenti
	// ---
	// EsigibilitaIVA: codice che esprime il regime di esigibilità dell’IVA (differita o immediata).
	EsigibilitaIVA string `xml:"EsigibilitaIVA" json:"EsigibilitaIVA"`
}

// Validate ...
func (f datiRiepilogo) Validate() error {

	if err := f.ImponibileImporto.Validate(); err != nil {
		return fmt.Errorf("ImponibileImporto %s", err)
	}

	if err := f.DatiIVA.Validate(); err != nil {
		return fmt.Errorf("DatiIVA %s", err)
	}

	if err := f.Natura.Validate(); err != nil {
		return fmt.Errorf("Natura %s", err)
	}

	f.Deducibile = strings.ToUpper(f.Deducibile)
	if f.Deducibile != "SI" {
		return fmt.Errorf("Deducibile %s", share.ErrorIncorrectValue(f.Deducibile))
	}

	f.EsigibilitaIVA = strings.ToUpper(f.EsigibilitaIVA)
	switch f.EsigibilitaIVA {
	case "I":
		break
	case "D":
		break
	case "S":
		break
	default:
		return fmt.Errorf("EsigibilitaIVA %s", share.ErrorIncorrectValue(f.EsigibilitaIVA))
	}

	if len(f.Detraibile) < 4 || len(f.Detraibile) > 6 {
		return fmt.Errorf("Detraibile %s", share.ErrorIncluded(4, 6))
	}

	return nil
}
