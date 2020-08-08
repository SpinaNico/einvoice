>This package deals with giving structures for the validation of electronic 
invoices in the Italian billing system

# This package is still experimental (not use in production)

[IT](https://www.agenziaentrate.gov.it/portale/documents/20143/2370834/Allegato+A+-+Specifiche+tecniche+vers+1.6_.pdf/a9917ec2-29a3-4f4a-a7d0-93af96fcaad5)

---

+ [Header](header.go)  is a file that shows all the structures necessary for the electronic invoice header to function. In these structures the clinician is defined as professional.

+ [Body](body.go) is a file containing all the structures for the correct functioning and validation of the FatturaElettronicaBody essential block in XML that represents the products and services listed on the invoice, also shows the data for payment and various and possible

+ [Share](share.go) this is a file containing structures shared between body and header, in particular the Master Data and tax data

---

#### How does it work?
import the following repository into your code
```go
import (
	invoice "github.com/SpinaNico/go-struct-invoice"
)
```
immediately after importing the code you will find the structure visible (publishes): `FatturaElettronica`
which is the following structure inside "invoice.go":

```go
type FatturaElettronica struct {
	FatturaElettronicaHeader FatturaElettronicaHeader
	FatturaElettronicaBody   []FatturaElettronicaBody 
}
```
**Note:** *I omitted the tags*
you also have two other large structures that are FatturaElettronicaHeader, FatturaElettronicaBody
---
the operation is as follows:

```go
	var f invoice.FatturaElettronica
	if err := f.Validate(); err != nil {
		log.Fatal(fmt.Errorf("ValidateInvoice: %s", err))
	}
```
