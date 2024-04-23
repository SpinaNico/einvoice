package invoicer

import (
	"time"

	"github.com/spinanico/einvoice/sdi"
)

type Invoicer interface {
	CreateEmptyInvoice(uniqueNumber string, ns string) Invoice
	FromFile(filePath string) (Invoice, error)
	FromBytes(data []byte) (Invoice, error)
	SaveToFile(Invoice, string) error
	GetBytes(Invoice) ([]byte, error)
}

type ProductParams struct {
	Name        string
	Quantity    float64
	Price       float64
	Natura      *sdi.Natura
	AliquotaIVA float64

	PeriodStart time.Time
	PeriodEnd   time.Time
}

type CustomerParams struct {
	NationTax     string
	CompanyTaxId  string
	PersonalTaxId string

	FirstName string
	LastName  string

	Denomination string

	Address       string
	AddressNumber string
	ZipCode       string
	City          string
	Province      string
	Nation        string
}

type SellerParams struct {
	CustomerParams

	RegimeFiscale sdi.RegimeFiscale
}

type TransmissionParams struct {
	TransmissionType sdi.FormatoTrasmissione

	Nation          string
	TaxId           string
	DestinationCode string
	PecAddress      string
	// if is empty, the UniqueTransmissionNumber will be not used
	UniqueTransmissionNumber string
}

type Invoice interface {
	AddProduct(product *ProductParams) Invoice
	SetTransmission(*TransmissionParams) Invoice
	SetCustomer(*CustomerParams) Invoice
	SetSeller(*SellerParams) Invoice
	SetInvoiceData(td sdi.TipoDocumento, number string, date time.Time) Invoice
	//Default is EUR
	SetCurrency(currency string) Invoice
	SetCursor(cursor int) Invoice

	Number() string
	Date() time.Time

	Total() float64
	TotalTax() float64

	FileName() string
}
