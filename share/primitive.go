package share

import (
	"fmt"
	"strconv"
	"strings"
)

// Cap postal code
type Cap string

// Validate ...
func (c Cap) Validate() error {
	// Cap
	if len(string(c)) != 5 {
		return fmt.Errorf(" (CAP): %s", ErrorEgual(5))
	}
	_, err := strconv.Atoi(string(c))
	if err != nil {
		return fmt.Errorf("(CAP): %s", ErrorMustBeInt())
	}

	return nil
}

type numeroCivico string

func (c numeroCivico) Validate() error {
	if len(string(c)) > 8 {
		return fmt.Errorf("(NumeroCivico): %s", ErrorMaxLength(8))
	}
	return nil
}

type comune string

func (c comune) Validate() error {
	if len(string(c)) > 60 {
		return fmt.Errorf("(Comune): %s", ErrorMaxLength(60))
	}
	return nil
}

type provincia string

func (c provincia) Validate() error {
	if len(string(c)) != 2 {
		return fmt.Errorf("(Provincia) %s", ErrorEgual(2))
	}
	return nil
}

type nazione string

func (c nazione) Validate() error {
	if len(string(c)) != 2 {
		return fmt.Errorf("(Nazione) %s", ErrorEgual(2))
	}
	c = nazione(strings.ToUpper(string(c)))
	return nil
}

type indirizzo string

func (c indirizzo) Validate() error {
	if len(string(c)) > 60 {
		return fmt.Errorf("(Indirizzo): %s", ErrorMaxLength(60))
	}

	return nil
}
