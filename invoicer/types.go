package invoicer

import "github.com/spinanico/einvoice/sdi"

type Invoicer interface {
	CreateEmptyInvoice(uniqueNumber string) Invoice
	FromFile(filePath string) (Invoice, error)
	FromBytes(data []byte) (Invoice, error)
}

type ProductParams struct {
	Name     string
	Quantity float64
	Price    float64
	Natura   *sdi.Natura
}

type CustomerParams struct {
	Nation        string
	CompanyTaxId  string
	PersonalTaxId string

	FirstName string
	LastName  string

	Denomination string

	Address  string
	ZipCode  string
	City     string
	Province string
}

type SellerParams struct {
	CustomerParams
}

type Invoice interface {
	AddProduct(product *ProductParams) Invoice
	SetCustomer(*CustomerParams) Invoice
	SetSeller(*SellerParams) Invoice

	SetCursor(cursor int) Invoice
}
