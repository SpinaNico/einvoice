package einvoice

import (
	"fmt"
	"testing"
)

const mockInvoiceName = "IT05692280877_0001"

func errorFormat(code string) string {
	return fmt.Sprintf("non è presente errore %s ", code)
}

func TestBaseSDI(t *testing.T) {

	t.Run("NomenclaturaFile", func(t *testing.T) {

		var f FatturaElettronica
		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00001") {
			t.Error(errorFormat("00001"))
		}
	})

	t.Run("Dimensione FIle ", func(t *testing.T) {
		t.Skip("Questo test è skippato è per controllare l'errore 00003")
	})

	t.Run("Signature firma elettronica sulla fattura", func(t *testing.T) {
		// XAdES-Bes.

	})

	t.Run("00400 DatiBeniServizi errore controllo sulla natura e aliquota iva 0", func(t *testing.T) {
		f := FatturaElettronica{
			FatturaElettronicaBody: []*FatturaElettronicaBody{
				{
					DatiBeniServizi: &DatiBeniServizi{
						DettaglioLinee: []*DettaglioLinee{
							{
								AliquotaIVA: 0,
								Natura:      "",
							},
						},
					},
				},
			},
		}

		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00400") {
			t.Error(errorFormat("00400"))
		}
	})
	t.Run("00413 DatiCassaPrevidenziale errore controllo sulla natura presente con aliquota iva 0 dei dati cassa prevedenziale", func(t *testing.T) {
		f := FatturaElettronica{
			FatturaElettronicaBody: []*FatturaElettronicaBody{
				{
					DatiGenerali: &DatiGenerali{

						DatiGeneraliDocumento: &DatiGeneraliDocumento{
							TipoDocumento: "TD01",
							DatiCassaPrevidenziale: &DatiCassaPrevidenziale{
								AliquotaIVA: 0,
								Natura:      "",
								TipoCassa:   "TC22",
							},
						},
					},
				},
			},
		}

		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00413") {
			t.Error(errorFormat("00413"))
		}
	})

	t.Run("00401 DatiBeniServizi Natura Presenta ma IVA diversa da zero", func(t *testing.T) {
		f := FatturaElettronica{
			FatturaElettronicaBody: []*FatturaElettronicaBody{
				{
					DatiBeniServizi: &DatiBeniServizi{
						DettaglioLinee: []*DettaglioLinee{
							{
								AliquotaIVA: 22,
								Natura:      "N1",
							},
						},
					},
				},
			},
		}

		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00401") {
			t.Error(errorFormat("00401"))
		}
	})

	t.Run("00414 DatiCassaPrevidenziale  Natura presenta ma iva diversa da zero CassaPrevidenziale", func(t *testing.T) {
		f := FatturaElettronica{
			FatturaElettronicaBody: []*FatturaElettronicaBody{
				{
					DatiGenerali: &DatiGenerali{

						DatiGeneraliDocumento: &DatiGeneraliDocumento{
							TipoDocumento: "TD01",
							DatiCassaPrevidenziale: &DatiCassaPrevidenziale{
								AliquotaIVA: 22,
								Natura:      "N1",
								TipoCassa:   "TC22",
							},
						},
					},
				},
			},
		}

		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00414") {
			t.Error(errorFormat("00414"))
		}
	})

	t.Run("00403 controllo sulla data del documento, prendendo come data per SDI oggi ", func(t *testing.T) {
		var f = placeholderFattura()
		f.FatturaElettronicaBody[0].DatiGenerali.DatiGeneraliDocumento.Data = "2080-01-01"
		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00403") {
			t.Error(errorFormat("00403"))
		}
	})

	t.Run("00411  controllo del blocco dati ritenuta in caso di ritenuta=SI su dettaglio linee", func(t *testing.T) {

		var f = placeholderFattura()
		f.FatturaElettronicaBody[0].DatiBeniServizi = &DatiBeniServizi{
			DettaglioLinee: []*DettaglioLinee{
				{
					Ritenuta: "SI",
				},
			},
		}

		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00411") {
			t.Error(errorFormat("00411"))
		}
	})

	t.Run("00415 DatiCassaPrevidenziale ritenuta=SI ma manca il blocco dati ritenuta", func(t *testing.T) {
		var f = placeholderFattura()
		f.FatturaElettronicaBody[0].DatiGenerali = &DatiGenerali{
			DatiGeneraliDocumento: &DatiGeneraliDocumento{
				DatiCassaPrevidenziale: &DatiCassaPrevidenziale{
					Ritenuta: "SI",
				},
			},
		}
		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00415") {
			t.Error(errorFormat("00415"))
		}
	})

	t.Run("00417 Non è presente almeno uno tra CodiceFisclae IdFiscaleIVA", func(t *testing.T) {

		var f = placeholderFattura()
		err := f.SDIValidation(mockInvoiceName)
		if err != nil && !err.ErrorCodeIsPresent("00417") {
			t.Error(errorFormat("00417"))
		}
	})

}

func placeholderFattura() *FatturaElettronica {
	return &FatturaElettronica{
		FatturaElettronicaHeader: FatturaElettronicaHeader{
			CessionarioCommittente: &CessionarioCommittente{},
		},
		FatturaElettronicaBody: []*FatturaElettronicaBody{
			{
				DatiGenerali: &DatiGenerali{

					DatiGeneraliDocumento: &DatiGeneraliDocumento{
						TipoDocumento: "TD01",
						DatiCassaPrevidenziale: &DatiCassaPrevidenziale{
							AliquotaIVA: 22,
							Natura:      "N1",
							TipoCassa:   "TC22",
						},
					},
				},
			},
		},
	}
}
