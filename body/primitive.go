package body

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SpinaNico/go-struct-invoice/share"
)

// YYYY-MM-DD.
type data string

func (c data) Validate() error {
	var err error
	if len(string(c)) != 10 {
		return fmt.Errorf("(Data) the format must be YYYY-MM-DD name of format: ISO 8601 2004 ")
	}
	return err
}

//DataOraConsegna la data deve essere rappresentata secondo il
//formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DDTHH:MM:SS.
type dataOra string

func (c dataOra) Validate() error {
	matched, _ := regexp.Match(`\d\d\d-\d\d-\d\d\d\d\d:\d\d:\d\d`, []byte(c))
	if !matched {
		return fmt.Errorf("%s", share.ErrorIncorrectValue(string(c)))
	}
	return nil
}

type decimale2 string

func (c decimale2) Validate() error {

	s := string(c)
	val, err := strconv.ParseFloat(s, 32)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	c = decimale2(fmt.Sprintf("%.2f", val))
	return nil
}

type numeroLinea int

func (c numeroLinea) Validate() error {

	if int(c) > 9999 {
		return fmt.Errorf(share.ErrorIncluded(1, 9999))
	}

	return nil
}

type codiceCup string

func (c codiceCup) Validate() error {

	if len(string(c)) > 15 {
		return fmt.Errorf(share.ErrorMaxLength(15))
	}
	return nil

}

type codiceCig string

func (c codiceCig) Validate() error {

	if len(string(c)) > 15 {
		return fmt.Errorf(share.ErrorMaxLength(15))
	}
	return nil

}

type idDocumento string

func (c idDocumento) Validate() error {
	if len(string(c)) > 20 {
		return fmt.Errorf("%s", share.ErrorMaxLength(20))
	}
	return nil
}

type codiceCommessaConvenzione string

func (c codiceCommessaConvenzione) Validate() error {
	if len(string(c)) > 100 {
		return fmt.Errorf("%s", share.ErrorMaxLength(100))
	}
	return nil
}

type numItem string

func (c numItem) Validate() error {
	if len(string(c)) > 20 {
		return fmt.Errorf("%s", share.ErrorMaxLength(20))
	}
	return nil
}

// Natura:
// **Regole**
// codice che esprime la natura della non
// imponibilità del contributo cassa.
// Deve essere presente nel solo caso in cui l’elemento AliquotaIVA vale zero.
// 40 Se è presente a fronte di un valore dell’elemento
// AliquotaIVA diverso da zero, il file viene scartato con
// codice errore 0414.
// ---
// valori possibili:
// + N1 *escluse ex art.15*
// + N2 *non soggette*
// + N3 *non imponibili*
// + N4 *esenti*
// + N5 *regime del margine / IVA non esposta in fattura*
// + N6 *inversione contabile (per le operazioni in reverse
//		charge ovvero nei casi di autofatturazione per
//		acquisti extra UE di servizi ovvero per
//		importazioni di beni nei soli casi previsti)*
// + N7 *IVA assolta in altro stato UE (vendite a distanza
//		ex art. 40 commi 3 e 4 e art. 41 comma 1 lett. b,
// 		DL 331/93; prestazione di servizi di telecomunicazioni,
//		tele-radiodiffusione ed elettronici ex art. 7-sexies lett. f, g,
// 		DPR 633/72 e art. 74-sexies, DPR 633/72)
type natura string

func (c natura) Validate() error {
	if len(string(c)) != 2 {
		return fmt.Errorf("Natura %s", share.ErrorEgual(2))
	}
	values := make(map[string]string)
	for i := 1; i < 8; i++ {
		values[fmt.Sprintf("N%d", i)] = "true"
	}
	if _, ok := values[string(c)]; ok == false {
		return fmt.Errorf("%s", share.ErrorIncorrectValue(string(c)))
	}
	return nil
}

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
type tipoDocumento string

func (c tipoDocumento) Validate() error {

	if len(string(c)) != 4 {
		return fmt.Errorf("%s", share.ErrorEgual(4))
	}
	c = tipoDocumento(strings.ToUpper(string(c)))
	switch string(c) {
	case "TD01":
		return nil
	case "TD02":
		return nil
	case "TD03":
		return nil
	case "TD04":
		return nil
	case "TD05":
		return nil
	case "TD06":
		return nil
	case "TD20":
		return nil
	default:
		return fmt.Errorf("%s", share.ErrorIncorrectValue(string(c)))
	}

}

type divisa string

func (c divisa) Validate() error {
	if len(string(c)) != 3 {
		return fmt.Errorf("%s", share.ErrorEgual(3))
	}
	c = divisa(strings.ToUpper(string(c)))
	return nil
}
