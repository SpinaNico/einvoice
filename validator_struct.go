package einvoice

import (
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

func validateCodiceArticolo(d validator.StructLevel) {
	data := d.Current().Interface().(codiceArticolo)
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
	data := d.Current().Interface().(datiTrasmissione)
	if data.CodiceDestinatario == "0000000" {
		if data.PECDestinatario == "" {
			d.ReportError(data.PECDestinatario, "PECDestinatario", "", "required", "")
		}
	}

}

// if the person making the invoice is a foreigner, the Italian office must be indicated.
// This validator also checks whether the Stabile Organization is in the
// Italian territory so the Nation value is "IT"
func cessionarioCommittenteValidate(d validator.StructLevel) {

	data := d.Current().Interface().(CessionarioCommittente)
	if data.Sede == nil {
		d.ReportError(data.Sede, "Sede", "cessionarioCommittenteValidate", "required", "")
	} else if data.Sede.Nazione != "IT" {

		if data.StabileOrganizzazione != nil {
			if *data.StabileOrganizzazione == (indirizzoType{}) {
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
}

func validateDatiBeniServizi(d validator.StructLevel) {
	data := d.Current().Interface().(datiBeniServizi)
	for _, linea := range data.DettaglioLinee {
		if linea.NumeroLinea == 0 {
			d.ReportError(linea.NumeroLinea, "NumeroLinea", "", "required", "")
		}

		v, _ := strconv.ParseFloat(linea.AliquotaIVA, 32)

		if v == 0 {
			if linea.Natura == "" {
				d.ReportError(linea.NumeroLinea, "NumeroLinea", "", ivaZeroNatureWrong, "")
			}
		}
	}
}
