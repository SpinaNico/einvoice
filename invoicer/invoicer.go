package invoicer

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spinanico/einvoice/sdi"
)

func New() Invoicer {
	return &invoicerImple{}
}

type invoicerImple struct {
}

func (i *invoicerImple) CreateEmptyInvoice(uniqueNumber string, namespace string) Invoice {

	header := createHeader()

	if namespace == "" {
		namespace = "ns"
	}

	header.DatiTrasmissione.ProgressivoInvio = uniqueNumber

	return &invoiceSt{
		fat: &sdi.FatturaElettronica{
			Versione:                 sdi.FPR12,
			FatturaElettronicaHeader: header,

			FatturaElettronicaBody: []*sdi.FatturaElettronicaBody{
				createBody(),
			},
			XMLName: xml.Name{
				Local: fmt.Sprintf("%s:FatturaElettronica", namespace),
				//Space: sdi.SpaceValue,
			},
			Xmlns: xml.Attr{
				Name: xml.Name{
					Local: fmt.Sprintf("xmlns:%s", namespace),
					//Space: sdi.SpaceValue,
				},
				Value: sdi.SpaceValue,
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

func (i *invoicerImple) SaveToFile(invoice Invoice, filePath string) error {

	bb, err := i.GetBytes(invoice)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bb, 0644)

}

func (i *invoicerImple) GetBytes(invoice Invoice) ([]byte, error) {

	data, err := xml.MarshalIndent(invoice.(*invoiceSt).fat, "", "    ")
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("%s\n%s", sdi.HeaderXMLInvoice, data)), nil

}

func (i *invoicerImple) GetRawInvoice(invoice Invoice) *sdi.FatturaElettronica {
	return invoice.(*invoiceSt).fat
}
