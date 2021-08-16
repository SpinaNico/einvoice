package einvoice

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

func registerAllValidatorStructLevel(validate *validator.Validate) {
	validate.RegisterStructValidation(datiTrasmissioneValidate, DatiTrasmissione{})
	validate.RegisterStructValidation(validateCessionarioCommitente, CessionarioCommittente{})
	validate.RegisterStructValidation(validateDatiBeniServizi, DatiBeniServizi{})
	validate.RegisterStructValidation(validateDatiCassaPrevidenziale, DatiCassaPrevidenziale{})
	validate.RegisterStructValidation(validateFatturaElettronicaBody, FatturaElettronicaBody{})
}

func validateFatturaElettronicaBody(d validator.StructLevel) {
	data := d.Current().Interface().(FatturaElettronicaBody)

	if data.DatiBeniServizi != nil {
		for _, linea := range data.DatiBeniServizi.DettaglioLinee {
			if linea.Ritenuta == "SI" {
				if len(data.DatiGenerali.DatiGeneraliDocumento.DatiRitenuta) == 0 {
					// ERROR 00411
					d.ReportError(linea.Ritenuta, "DatiBeniServizi", "Ritenuta", ritenutaSiWithoutDatiRitenuta, linea.Ritenuta)
				}
			}
		}
	}

	if data.DatiGenerali != nil {
		if data.DatiGenerali.DatiGeneraliDocumento != nil {
			if data.DatiGenerali.DatiGeneraliDocumento.DatiCassaPrevidenziale != nil {
				if data.DatiGenerali.DatiGeneraliDocumento.DatiCassaPrevidenziale.Ritenuta == "SI" {
					if len(data.DatiGenerali.DatiGeneraliDocumento.DatiRitenuta) == 0 {
						// ERROR 00415
						d.ReportError("DatiCassaPrevidenziale", "DatiCassaPrevidenziale", "Ritenuta", ritenutaSiWithoutDatiRitenuta, "")
					}
				}
			}
		}
	}

}

func validateCodiceArticolo(d validator.StructLevel) {
	data := d.Current().Interface().(CodiceArticolo)
	/*
		CodiceValore: valore del codice articolo corrispondente alla tipologia. Se l’elemento CodiceTipo è uguale a “CARB”, l’elemento CodiceValore deve contenere uno dei seguenti codici:
		- “27101245” (per vendita di Benzina senza piombo ottani => 95 e < 98);
		- “27101249” (per vendita di Benzina senza piombo ottani => 98) ;
		- “27101943” (per vendita di Olii da gas aventi tenore, in peso, di zolfo inferiore o uguale a 0,001%);
		- “27102011” (per vendita di Olio da gas denaturato tenore in peso di zolfo nell'olio da gas =< 0,001%)

		riportati nella tabella di riferimento per i prodotti energetici TA13 – pubblicata sul sito dell’Agenzia delle Dogane
	*/
	if data.CodiceTipo == "CARB" {
		if data.CodiceValore != "27101245" &&
			data.CodiceValore != "27101249" &&
			data.CodiceValore != "27101943" &&
			data.CodiceValore != "27102011" {
			d.ReportError(data.CodiceValore, "CodiceValore", "CodiceValore", carbWrongArticleCode, data.CodiceValore)
		}
	}
}

// Struct Level control
// This control at the structure level is essential, check that if the
// target code is "0000000" the pec is not empty
func datiTrasmissioneValidate(d validator.StructLevel) {
	data := d.Current().Interface().(DatiTrasmissione)
	if data.CodiceDestinatario == "0000000" {
		if data.PECDestinatario == "" {
			d.ReportError(data.PECDestinatario, "PECDestinatario", "", "required", "")
		}
	}

}

// if the person making the invoice is a foreigner, the Italian office must be indicated.
// This validator also checks whether the Stabile Organization is in the
// Italian territory so the Nation value is "IT"
func validateCessionarioCommitente(d validator.StructLevel) {

	data := d.Current().Interface().(CessionarioCommittente)
	if data.Sede == nil {
		d.ReportError(data.Sede, "Sede", "validateCessionarioCommitente", "required", "")
	} else if data.Sede.Nazione != "IT" {

		if data.StabileOrganizzazione != nil {
			if *data.StabileOrganizzazione == (IndirizzoType{}) {
				d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "required", "")
			} else {
				if data.StabileOrganizzazione.Nazione != "IT" {
					d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "eq", "")
				}
			}
		} else {
			d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "required", "")
		}
	}

	if data.DatiAnagrafici != nil {
		if data.DatiAnagrafici.IDFiscaleIVA != nil {
			if data.DatiAnagrafici.IDFiscaleIVA.IDCodice == "" && data.DatiAnagrafici.CodiceFiscale == "" {
				d.ReportError("", "DatiAnagrafici", "", notExistsFiscalID, "")
			}
		} else if data.DatiAnagrafici.CodiceFiscale == "" {
			d.ReportError("", "DatiAnagrafici", "", notExistsFiscalID, "")
		}
	} else {
		d.ReportError("", "DatiAnagrafici", "", notExistsFiscalID, "")
	}

}

type aliquote struct {
	value  float64
	nature string
}

func validateDatiBeniServizi(d validator.StructLevel) {
	data := d.Current().Interface().(DatiBeniServizi)
	aliquotes := []*aliquote{}
	for _, linea := range data.DettaglioLinee {
		if linea.NumeroLinea == 0 {
			d.ReportError(linea.NumeroLinea, "NumeroLinea", "", "required", "")
		}

		if linea.AliquotaIVA == 0 {
			if linea.Natura == "" {
				d.ReportError("DatiBeniServizi", "DatiBeniServizi", "", ivaZeroNatureWrong, "")
			}
		} else {
			if linea.Natura != "" {
				d.ReportError("DatiBeniServizi", "DatiBeniServizi", "", ivaNotZeroNaturePresent, "")
			}
		}
		var ok = false
		for _, a := range aliquotes {
			if fmt.Sprintf("%.2f", a.value) == fmt.Sprintf("%.2f", linea.AliquotaIVA) {
				ok = true
				if linea.AliquotaIVA == 0 {
					if a.nature != linea.Natura {
						ok = false
					}
				}
			}
		}
		if ok == false {
			aliquotes = append(aliquotes, &aliquote{value: linea.AliquotaIVA, nature: linea.Natura})
		}
		ok = false
		for _, a := range aliquotes {
			for _, r := range data.DatiRiepilogo {
				if r.AlliquotaIVA == a.value && r.Natura == a.nature {
					ok = true
				}
			}
			if ok == false {
				d.ReportError("", "DatiRiepilogo", "DatiBeniServizi", aliquoteNotOk, fmt.Sprintf("%v", a))
			}
		}
	}

}

func validateDatiCassaPrevidenziale(d validator.StructLevel) {
	data := d.Current().Interface().(DatiCassaPrevidenziale)

	if data.AliquotaIVA == 0 {
		if data.Natura == "" {
			d.ReportError("DatiCassaPrevidenziale", "DatiCassaPrevidenziale", "", ivaZeroNatureWrong, "")
		}
	} else {
		if data.Natura != "" {
			d.ReportError("DatiCassaPrevidenziale", "DatiCassaPrevidenziale", "", ivaNotZeroNaturePresent, "")
		}
	}

}
