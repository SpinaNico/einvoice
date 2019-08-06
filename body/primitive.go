package body

import (
	"fmt"
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
