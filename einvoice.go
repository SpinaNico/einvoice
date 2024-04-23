package einvoice

import "github.com/spinanico/einvoice/invoicer"

func NewInvoicer() invoicer.Invoicer {
	return invoicer.New()
}
