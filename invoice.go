package einvoice

// FatturaElettronica The structure of a multi body electronic invoice, if you need only one body,
//just insert a single element in the FatturaElettronicaBod slice...
type FatturaElettronica struct {
	FatturaElettronicaHeader FatturaElettronicaHeader `xml:"FatturaElettronicaHeader" json:"FatturaElettronicaHeader"`
	FatturaElettronicaBody   []FatturaElettronicaBody `xml:"FatturaElettronicaBody" json:"FatturaElettronicaBody"`
	// ds:Signature todo: firma digitale da implementare
	Signature string `xml:"ds:Signature" json:"Signature"`
}

// Version The reference version of the electronic invoice returns
func (f FatturaElettronica) Version() string { return "1.6.1" }
