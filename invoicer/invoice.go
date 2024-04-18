package invoicer

import (
	"github.com/spinanico/einvoice/sdi"
)

type invoiceSt struct {
	fat    *sdi.FatturaElettronica
	cursor int
}

func (i *invoiceSt) SetCursor(cursor int) Invoice {
	i.cursor = cursor

	if i.cursor < 0 {
		i.cursor = 0
	}

	if i.cursor >= len(i.fat.FatturaElettronicaBody) {
		i.cursor = len(i.fat.FatturaElettronicaBody) - 1
	}

	return i
}

func (i *invoiceSt) AddProduct(product *ProductParams) Invoice {
	i.checkBody()

	body := i.fat.FatturaElettronicaBody[i.cursor]

	if body.DatiBeniServizi.DettaglioLinee == nil {
		body.DatiBeniServizi.DettaglioLinee = make([]*sdi.DettaglioLinee, 0)
	}

	body.DatiBeniServizi.DettaglioLinee = append(body.DatiBeniServizi.DettaglioLinee, &sdi.DettaglioLinee{
		NumeroLinea:    len(body.DatiBeniServizi.DettaglioLinee) + 1,
		Quantita:       sdi.F64(product.Quantity),
		PrezzoUnitario: sdi.F64(product.Price),
		PrezzoTotale:   sdi.F64(product.Quantity * product.Price),
		AliquotaIVA:    22.0,
		Natura:         *product.Natura,
	})

	body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento = 0

	for _, line := range body.DatiBeniServizi.DettaglioLinee {
		body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento += line.PrezzoTotale
	}

	body.DatiBeniServizi.DatiRiepilogo = make([]*sdi.DatiRiepilogo, 0)

	for _, line := range body.DatiBeniServizi.DettaglioLinee {

		for _, r := range body.DatiBeniServizi.DatiRiepilogo {

			if r.AliquotaIVA == line.AliquotaIVA {
				r.ImponibileImporto += line.PrezzoTotale
				r.Imposta += line.PrezzoTotale * (line.AliquotaIVA / 100)
				r.AliquotaIVA = line.AliquotaIVA
				r.Natura = line.Natura
				break
			}

		}

	}

	return i
}

func (i *invoiceSt) SetCustomer(customer *CustomerParams) Invoice {

	i.checkHeader()

	var cessionarioCommittente = sdi.CessionarioCommittente{
		DatiAnagrafici: &sdi.DatiAnagrafici{
			CodiceFiscale: customer.PersonalTaxId,
			Anagrafica: &sdi.Anagrafica{
				Denominazione: customer.Denomination,
				Nome:          customer.FirstName,
				Cognome:       customer.LastName,
			},
		},
		Sede: &sdi.IndirizzoType{
			Indirizzo: customer.Address,
			CAP:       customer.ZipCode,
			Comune:    customer.City,
			Provincia: customer.Province,
			Nazione:   customer.Nation,
		},
	}

	i.fat.FatturaElettronicaHeader.CessionarioCommittente = &cessionarioCommittente

	return i
}

func (i *invoiceSt) SetSeller(seller *SellerParams) Invoice {

	i.checkHeader()

	var cedentePrestatore = sdi.CedentePrestatore{
		DatiAnagrafici: &sdi.DatiAnagrafici{
			CodiceFiscale: seller.PersonalTaxId,
			Anagrafica: &sdi.Anagrafica{
				Denominazione: seller.Denomination,
				Nome:          seller.FirstName,
				Cognome:       seller.LastName,
			},
		},
		Sede: &sdi.IndirizzoType{
			Indirizzo: seller.Address,
			CAP:       seller.ZipCode,
			Comune:    seller.City,
			Provincia: seller.Province,
			Nazione:   seller.Nation,
		},
	}

	i.fat.FatturaElettronicaHeader.CedentePrestatore = &cedentePrestatore

	return i
}

func (i *invoiceSt) checkBody() {
	if len(i.fat.FatturaElettronicaBody) == 0 {
		i.fat.FatturaElettronicaBody = append(i.fat.FatturaElettronicaBody, createBody())
		return
	} else {
		for _, body := range i.fat.FatturaElettronicaBody {
			if body.DatiGenerali == nil {
				body.DatiGenerali = &sdi.DatiGenerali{}
			}
			if body.DatiBeniServizi == nil {
				body.DatiBeniServizi = &sdi.DatiBeniServizi{}
			}
			if body.DatiPagamento == nil {
				body.DatiPagamento = &sdi.DatiPagamento{}
			}
		}
	}

}

func (i *invoiceSt) checkHeader() {
	if i.fat.FatturaElettronicaHeader == nil {
		i.fat.FatturaElettronicaHeader = createHeader()
	}

}

func createBody() *sdi.FatturaElettronicaBody {
	body := &sdi.FatturaElettronicaBody{
		DatiGenerali: &sdi.DatiGenerali{
			DatiGeneraliDocumento: &sdi.DatiGeneraliDocumento{},
		},
		DatiBeniServizi: &sdi.DatiBeniServizi{},
		DatiPagamento:   &sdi.DatiPagamento{},
	}
	return body
}

func createHeader() *sdi.FatturaElettronicaHeader {
	header := &sdi.FatturaElettronicaHeader{
		DatiTrasmissione: &sdi.DatiTrasmissione{},
		CedentePrestatore: &sdi.CedentePrestatore{
			DatiAnagrafici: &sdi.DatiAnagrafici{},
			Sede:           &sdi.IndirizzoType{},
		},
		CessionarioCommittente: &sdi.CessionarioCommittente{
			DatiAnagrafici: &sdi.DatiAnagrafici{},
			Sede:           &sdi.IndirizzoType{},
		},
	}
	return header
}
