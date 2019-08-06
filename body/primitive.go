package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
	"strconv"
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
