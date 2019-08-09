package header

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type regimeFiscale string

func (rf regimeFiscale) Validate() error {
	// RegimeFiscale
	RF := strings.ToUpper(string(rf))
	rf = regimeFiscale(RF)
	if len(string(rf)) != 4 {
		return fmt.Errorf("(RegimeFiscale) %s", share.ErrorEgual(4))
	}
	if string(string(rf)[0]) != "R" && string(string(rf)[1]) != "F" {
		return fmt.Errorf("(RegimeFiscale) must start with the \"RF\" characters")
	}
	regimeFiscale := make(map[string]string)
	for i := 1; i < 20; i++ {
		regimeFiscale[fmt.Sprintf("RF%d", i)] = "true"
	}
	_, exists := regimeFiscale[string(rf)]
	if exists == false {
		return fmt.Errorf("(RegimeFiscale) %s", share.ErrorIncorrectValue(string(rf)))
	}

	return nil
}

type codiceFiscale string

func (c codiceFiscale) Validate() error {
	cl := len(string(c))
	if !(cl >= 11 && cl <= 16) {
		return fmt.Errorf("(CodiceFiscale) %s", share.ErrorIncluded(11, 16))
	}
	return nil
}

type telefono string

func (c telefono) Validate() error {
	if !(len(string(c)) >= 5 && len(string(c)) <= 12) {
		return fmt.Errorf("(Telefono): %s", share.ErrorIncluded(5, 12))

	}
	return nil
}

type email string

func (c email) Validate() error {
	if !(len(string(c)) >= 7 && len(string(c)) <= 256) {
		return fmt.Errorf("%s", share.ErrorIncluded(7, 256))
	}
	m, _ := regexp.Match(`\w*@\w*\.\w*`, []byte(c))
	if m == false {
		return fmt.Errorf("%s", share.ErrorIncorrectValue(string(c)))
	}
	return nil
}

type pecDestinatario string

func (c pecDestinatario) Validate() error {
	if !(len(string(c)) >= 7 && len(string(c)) <= 256) {
		return fmt.Errorf("(PECDestinatario): %s", share.ErrorIncluded(7, 256))

	}
	return nil
}

type progressivoInvio string

func (c progressivoInvio) Validate() error {
	if len(string(c)) > 10 {
		return fmt.Errorf("(ProgressivoInvio): %s", share.ErrorMaxLength(10))
	}

	return nil
}

type codiceDestinatario string

func (c codiceDestinatario) Validate() error {
	if len(string(c)) != 7 {
		return fmt.Errorf("(CodiceDestinatario): %s", share.ErrorEgual(7))
	}
	return nil
}

type formatoTrasmissione string

func (c formatoTrasmissione) Validate() error {
	if string(c) != "FPR12" {
		return fmt.Errorf("(FormatoTrasmissione): is not FPR12")
	}
	return nil
}
