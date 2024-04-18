package sdisender

import "github.com/spinanico/einvoice/invoicer"

type Sender interface {
	Send(i invoicer.Invoice) error
}
