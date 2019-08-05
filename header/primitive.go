package header

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
	"strconv"
	"strings"
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

// Cap postal code
type Cap string

// Validate ...
func (c Cap) Validate() error {
	// Cap
	if len(string(c)) != 5 {
		return fmt.Errorf(" (CAP): %s", share.ErrorEgual(5))
	}
	_, err := strconv.Atoi(string(c))
	if err != nil {
		return fmt.Errorf("(CAP): %s", share.ErrorMustBeInt())
	}

	return nil
}

type numeroCivico string

func (c numeroCivico) Validate() error {
	if len(string(c)) > 8 {
		return fmt.Errorf("(NumeroCivico): %s", share.ErrorMaxLength(8))
	}
	return nil
}

type comune string

func (c comune) Validate() error {
	if len(string(c)) > 60 {
		return fmt.Errorf("(Comune): %s", share.ErrorMaxLength(60))
	}
	return nil
}

type provincia string

func (c provincia) Validate() error {
	if len(string(c)) != 2 {
		return fmt.Errorf("(Provincia) %s", share.ErrorEgual(2))
	}
	return nil
}

type nazione string

func (c nazione) Validate() error {
	if len(string(c)) != 2 {
		return fmt.Errorf("(Nazione) %s", share.ErrorEgual(2))
	}
	c = nazione(strings.ToUpper(string(c)))
	return nil
}

type indirizzo string

func (c indirizzo) Validate() error {
	if len(string(c)) > 60 {
		return fmt.Errorf("(Indirizzo): %s", share.ErrorMaxLength(60))
	}

	return nil
}
