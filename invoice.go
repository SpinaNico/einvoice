package einvoice

import "fmt"

// FatturaElettronica The structure of a multi body electronic invoice, if you need only one body,
//just insert a single element in the FatturaElettronicaBod slice...
type FatturaElettronica struct {
	FatturaElettronicaHeader FatturaElettronicaHeader `xml:"FatturaElettronicaHeader" json:"FatturaElettronicaHeader"`
	FatturaElettronicaBody   []FatturaElettronicaBody `xml:"FatturaElettronicaBody" json:"FatturaElettronicaBody"`
	// ds:Signature todo: firma digitale da implementare
}

// Version The reference version of the electronic invoice returns
func (f FatturaElettronica) Version() string { return "1.2" }

// Validate : Valid all the structure in each part, header and all the
// bodies present, if an error exists it is reported In the case of an error within the invoice body slice.
// This method returns an error indicating the index of the error in square brackets
func (f FatturaElettronica) Validate() error {
	if err := f.FatturaElettronicaHeader.Validate(); err != nil {
		return err
	}
	for index, value := range f.FatturaElettronicaBody {
		if err := value.Validate(); err != nil {
			return fmt.Errorf("[%d] %s", index+1, err)
		}
	}
	return nil
}
