package header

import (
	share "github.com/SpinaNico/go-struct-invoice/share"
)

// Trovi descrizioni nella sezione 2.2.4.1 Dati Anagrafici
type datiAnagrafici struct {
	CodiceFiscale string             `xml:"CodiceFiscale" json:"CodiceFiscale" validate:"omitempty,min=11,max=16"`
	Anagrafica    share.Anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA  share.IDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`

	// AlboProfessionale: formato alfanumerico; lunghezza massima di 60 caratteri.
	AlboProfessionale string `xml:"AlboProfessionale" json:"AlboProfessionale" validate:"omitempty,max=60"`
	// ProvinciaAlbo: formato alfanumerico; lunghezza di 2 caratteri.
	ProvinciaAlbo string `xml:"ProvinciaAlbo" json:"ProvinciaAlbo" validate:"omitempty,len=2"`
	// NumeroIscrizioneAlbo: formato alfanumerico; lunghezza massima di 60 caratteri.
	NumeroIscrizioneAlbo string `xml:"NumeroIscrizioneAlbo" json:"NumeroIscrizioneAlbo" validate:"omitempty,max=60"`
	// DataIscrizioneAlbo: la data deve essere rappresentata secondo il formato ISO 8601:2004, con la seguente precisione: YYYY-MM-DD.
	DataIscrizioneAlbo string `xml:"DataIscrizioneAlbo" json:"DataIscrizioneAlbo"`

	//RegimeFiscale: formato alfanumerico; lunghezza di 4 caratteri; i
	//valori ammessi sono i seguenti:
	// RF01 - Ordinario;
	// RF02 - Contribuenti minimi (art. 1, c.96-117, L. 244/2007);
	// RF04 - Agricoltura e attività connesse e pesca (artt. 34 e 34-bis, D.P.R. 633/1972);
	// RF05 - Vendita sali e tabacchi (art. 74, c.1, D.P.R. 633/1972);
	// RF06 - Commercio dei fiammiferi (art. 74, c.1, D.P.R. 633/1972);
	// RF07 - Editoria (art. 74, c.1, D.P.R. 633/1972);
	// RF08 - Gestione di servizi di telefonia pubblica (art. 74, c.1, D.P.R. 633/1972);
	// RF09 - Rivendita di documenti di trasporto pubblico e di sosta (art. 74, c.1, D.P.R. 633/1972);
	// RF10 - Intrattenimenti, giochi e altre attività di cui alla tariffa allegata al D.P.R. n. 640/72 (art. 74, c.6, D.P.R. 633/1972);
	// RF11 - Agenzie di viaggi e turismo (art. 74-ter, D.P.R. 633/1972);
	// RF12 - 633/1972); Agriturismo (art. 5, c.2, L. 413/1991);
	// RF13 - Vendite a domicilio (art. 25-bis, c.6, D.P.R. 600/1973);
	// RF14 - Rivendita di beni usati, di oggetti d’arte, d’antiquariato o da collezione (art. 36, D.L. 41/1995);
	// RF15 - Agenzie di vendite all’asta di oggetti d’arte, antiquariato o da collezione (art. 40-bis, D.L. 41/1995);
	// RF17 - IVA per cassa (art. 32-bis, D.L. 83/2012);
	// RF18 - Altro;
	// RF19 - Forfettario (art.1, c. 54-89, L. 190/2014)
	RegimeFiscale string `xml:"RegimeFiscale" json:"RegimeFiscale" validate:"omitempty,len=4,startswith=RF,regimeValidate"`
}

// func (c datiAnagrafici) Validate() error {
// 	var err error
// 	err = c.IDFiscaleIVA.Validate()
// 	if err != nil {
// 		return fmt.Errorf("DatiAnagrafici %s", err)
// 	}
// 	err = c.Anagrafica.Validate()
// 	if err != nil {
// 		return fmt.Errorf("DatiAnagrafici %s", err)
// 	}
// 	err = c.CodiceFiscale.Validate()
// 	if err != nil {
// 		return fmt.Errorf("DatiAnagrafici %s", err)
// 	}
// 	err = c.RegimeFiscale.Validate()
// 	if err != nil {
// 		return fmt.Errorf("DatiAnagrafici %s", err)
// 	}

// 	palen := len(c.ProvinciaAlbo)
// 	if palen != 0 && palen != 2 {
// 		return fmt.Errorf("DatiAnagrafici (ProvinciaAlbo) %s", share.ErrorEgual(2))
// 	}

// 	aplen := len(c.AlboProfessionale)
// 	if aplen != 0 && aplen > 60 {
// 		r= len(c.NumeroIscrizioneAlbo)
// 	if nialen != 0 && nialen > 60 {
// 		return fmt.Errorf("DatiAnagrafici (NumeroIscrizioneAlbo) %s", share.ErrorMaxLength(60))
// 	}

// 	return nil
// }
// eturn fmt.Errorf("DatiAnagrafici (AlboProfessionale) %s", share.ErrorMaxLength(60))
// 	}= len(c.NumeroIscrizioneAlbo)
// 	if nialen != 0 && nialen > 60 {
// 		return fmt.Errorf("DatiAnagrafici (NumeroIscrizioneAlbo) %s", share.ErrorMaxLength(60))
// 	}

// 	return nil
// }

// = len(c.NumeroIscrizioneAlbo)
// 	if nialen != 0 && nialen > 60 {
// 		return fmt.Errorf("DatiAnagrafici (NumeroIscrizioneAlbo) %s", share.ErrorMaxLength(60))
// 	}

// 	return nil
// }

// 	niale= len(c.NumeroIscrizioneAlbo)
// 	if nialen != 0 && nialen > 60 {
// 		return fmt.Errorf("DatiAnagrafici (NumeroIscrizioneAlbo) %s", share.ErrorMaxLength(60))
// 	}

// 	return nil
// }
// n := len(c.NumeroIscrizioneAlbo)
// 	if nialen != 0 && nialen > 60 {
// 		return fmt.Errorf("DatiAnagrafici (NumeroIscrizioneAlbo) %s", share.ErrorMaxLength(60))
// 	}

// 	return nil
// }
