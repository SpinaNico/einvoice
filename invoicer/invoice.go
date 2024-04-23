package invoicer

import (
	"fmt"
	"time"

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

	var nature sdi.Natura

	if product.Natura != nil {
		nature = *product.Natura
	}

	body.DatiBeniServizi.DettaglioLinee = append(body.DatiBeniServizi.DettaglioLinee, &sdi.DettaglioLinee{
		NumeroLinea:       len(body.DatiBeniServizi.DettaglioLinee) + 1,
		Quantita:          sdi.F64(product.Quantity),
		PrezzoUnitario:    sdi.F64(product.Price),
		PrezzoTotale:      sdi.F64(product.Quantity * product.Price),
		AliquotaIVA:       sdi.F64(product.AliquotaIVA),
		Natura:            nature,
		Descrizione:       product.Name,
		DataInizioPeriodo: sdi.Date{Time: product.PeriodStart},
		DataFinePeriodo:   sdi.Date{Time: product.PeriodEnd},
	})

	body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento = 0

	for _, line := range body.DatiBeniServizi.DettaglioLinee {
		body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento += line.PrezzoTotale
	}

	body.DatiBeniServizi.DatiRiepilogo = make([]*sdi.DatiRiepilogo, 0)

	for _, line := range body.DatiBeniServizi.DettaglioLinee {
		var found bool
		for _, r := range body.DatiBeniServizi.DatiRiepilogo {

			if r.AliquotaIVA == line.AliquotaIVA {
				r.ImponibileImporto += line.PrezzoTotale
				r.Imposta += line.PrezzoTotale * (line.AliquotaIVA / 100)
				r.AliquotaIVA = line.AliquotaIVA
				r.Natura = line.Natura
				found = true
				break
			}

		}

		if !found {
			body.DatiBeniServizi.DatiRiepilogo = append(body.DatiBeniServizi.DatiRiepilogo, &sdi.DatiRiepilogo{
				AliquotaIVA:       line.AliquotaIVA,
				ImponibileImporto: line.PrezzoTotale,
				Imposta:           line.PrezzoTotale * (line.AliquotaIVA / 100),
				Natura:            line.Natura,
			})

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
			IdFiscaleIVA: &sdi.IdFiscaleIVA{
				IdPaese:  customer.NationTax,
				IdCodice: customer.CompanyTaxId,
			},
		},
		Sede: &sdi.IndirizzoType{
			Indirizzo:    customer.Address,
			CAP:          customer.ZipCode,
			Comune:       customer.City,
			Provincia:    customer.Province,
			Nazione:      customer.NationTax,
			NumeroCivico: customer.AddressNumber,
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
			IdFiscaleIVA: &sdi.IdFiscaleIVA{
				IdPaese:  seller.NationTax,
				IdCodice: seller.CompanyTaxId,
			},
			RegimeFiscale: seller.RegimeFiscale,
		},
		Sede: &sdi.IndirizzoType{
			Indirizzo:    seller.Address,
			CAP:          seller.ZipCode,
			Comune:       seller.City,
			Provincia:    seller.Province,
			Nazione:      seller.NationTax,
			NumeroCivico: seller.AddressNumber,
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
		DatiPagamento:   nil,
		DatiVeicolo:     nil,
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

func (i *invoiceSt) Number() string {
	return i.fat.FatturaElettronicaBody[i.cursor].DatiGenerali.DatiGeneraliDocumento.Numero
}

func (i *invoiceSt) Date() time.Time {
	return i.fat.FatturaElettronicaBody[i.cursor].DatiGenerali.DatiGeneraliDocumento.Data.Time
}

func (i *invoiceSt) Total() float64 {
	var total float64

	for _, body := range i.fat.FatturaElettronicaBody {
		if body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento == 0 {
			for _, line := range body.DatiBeniServizi.DettaglioLinee {
				body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento += line.PrezzoTotale
			}

		}

		total += float64(body.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento)
	}

	return total
}

func (i *invoiceSt) TotalTax() float64 {

	var total float64

	for _, body := range i.fat.FatturaElettronicaBody {
		for _, r := range body.DatiBeniServizi.DatiRiepilogo {
			total += float64(r.Imposta)
		}
	}

	return total
}

func (i *invoiceSt) FileName() string {
	return fmt.Sprintf("%s%s_%05s",
		i.fat.FatturaElettronicaHeader.CedentePrestatore.DatiAnagrafici.IdFiscaleIVA.IdPaese,
		i.fat.FatturaElettronicaHeader.CedentePrestatore.DatiAnagrafici.IdFiscaleIVA.IdCodice,
		i.fat.FatturaElettronicaHeader.DatiTrasmissione.ProgressivoInvio) + ".xml"
}

func (i *invoiceSt) SetTransmission(transmission *TransmissionParams) Invoice {

	unique := transmission.UniqueTransmissionNumber

	if unique == "" {

		unique = i.fat.FatturaElettronicaHeader.DatiTrasmissione.ProgressivoInvio
	}

	if unique == "" {
		panic("UniqueTransmissionNumber is empty")
	}

	if transmission.TransmissionType == "" {
		transmission.TransmissionType = sdi.FPR12
	}

	i.fat.FatturaElettronicaHeader.DatiTrasmissione = &sdi.DatiTrasmissione{
		ProgressivoInvio:    unique,
		FormatoTrasmissione: transmission.TransmissionType,

		CodiceDestinatario: transmission.DestinationCode,
		PECDestinatario:    transmission.PecAddress,
		IdTrasmittente: &sdi.IdFiscaleIVA{
			IdPaese:  transmission.Nation,
			IdCodice: transmission.TaxId,
		},
	}

	i.fat.Versione = transmission.TransmissionType

	return i
}

func (i *invoiceSt) SetInvoiceData(td sdi.TipoDocumento, number string, date time.Time) Invoice {

	i.fat.FatturaElettronicaBody[i.cursor].DatiGenerali.DatiGeneraliDocumento = &sdi.DatiGeneraliDocumento{
		TipoDocumento: td,
		Numero:        number,
		Data:          sdi.Date{Time: date},
		Divisa:        "EUR",
	}

	return i
}

func (i *invoiceSt) SetCurrency(currency string) Invoice {
	i.fat.FatturaElettronicaBody[i.cursor].DatiGenerali.DatiGeneraliDocumento.Divisa = currency
	return i
}
