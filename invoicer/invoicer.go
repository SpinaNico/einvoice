package invoicer

import (
	"encoding/xml"
	"os"

	"github.com/spinanico/einvoice/sdi"
)

func New() Invoicer {
	return &invoicerImple{}
}

type invoicerImple struct {
}

func (i *invoicerImple) CreateEmptyInvoice(uniqueNumber string) Invoice {

	return &invoiceSt{
		fat: &sdi.FatturaElettronica{
			FatturaElettronicaHeader: createHeader(),

			FatturaElettronicaBody: []*sdi.FatturaElettronicaBody{
				createBody(),
			},
		},
	}
}

func (i *invoicerImple) FromFile(filePath string) (Invoice, error) {

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return i.FromBytes(file)

}

func (i *invoicerImple) FromBytes(data []byte) (Invoice, error) {

	fat := &sdi.FatturaElettronica{}

	err := xml.Unmarshal(data, fat)

	if err != nil {
		return nil, err
	}

	return &invoiceSt{
		fat: fat,
	}, nil

}
