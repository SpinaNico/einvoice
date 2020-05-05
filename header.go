package invoice

import "fmt"

type datiTrasmissione struct {
	IDTrasmittente *iDFiscaleIVA `xml:"IdTrasmittente" json:"IdTrasmittente" validate:"required"`
	//ProgressivoInvio: progressivo che il soggetto trasmittente attribuisce
	//al file che inoltra al Sistema di Interscambio per una propria finalità di
	//identificazione univoca.
	ProgressivoInvio string `xml:"ProgressivoInvio" json:"ProgressivoInvio" validate:"required,max=10"`
	//FormatoTrasmissione: codice identificativo del tipo di trasmissione
	//che si sta effettuando e del relativo formato. Va sempre valorizzato
	//con “FPR12”.
	FormatoTrasmissione string `xml:"FormatoTrasmissione" json:"FormatoTrasmissione" validate:"required,oneof=FPR12 FPA12"`

	//CodiceDestinatario: identifica il canale telematico sul quale
	//recapitare la fattura; deve contenere un valore alfanumerico di 7
	//caratteri corrispondente a:
	//- uno dei codici che il Sistema di Interscambio attribuisce ai
	//	soggetti, con canale accreditato in ricezione, che ne abbiano
	//	fatto richiesta attraverso la funzione ‘Richiesta codici destinatario
	//	B2B’ presente sul sito www.fatturapa.gov.it;
	//- ‘0000000’, nei casi di fattura destinata ad un soggetto che riceve
	//	tramite PEC e questa sia stata indicata nel campo
	//	PECDestinatario;
	//- ‘0000000’, nei casi di fattura destinata ad un soggetto per il quale
	// 	non si conosce il canale telematico (PEC o altro) sul quale
	// 	recapitare il file.
	//- XXXXXXX’, in caso di fattura emessa verso soggetti non
	//	residenti, non stabiliti, non identificati in Italia, e inviata al
	//	Sistema di Interscambio al fine di trasmettere i dati.
	// - L'amministrazione publica ha il codice univoco di soli  6 caratteri
	//   corrisponde al numero di ufficio
	CodiceDestinatario   string                `xml:"CodiceDestinatario" json:"CodiceDestinatario" validate:"min=6,max=7"`
	ContattiTrasmittente *contattiTrasmittente `xml:"ContattiTrasmittente" json:"ContattiTrasmittente"`

	//PECDestinatario: indirizzo di Posta Elettronica Certificata al quale, se
	//valorizzato, viene recapitata la fattura nei casi in cui il valore di
	//CodiceDestinatario sia uguale a ‘0000000’ e non risulti registrato alcun
	//canale telematico associato alla partita IVA del cessionario/committente.
	PECDestinatario string `xml:"PECDestinatario" json:"PECDestinatario" validate:"omitempty,min=7,max=256,email"`
}

// Trovi descrizioni nella sezione 2.2.4.1 Dati Anagrafici
type datiAnagrafici struct {
	CodiceFiscale string        `xml:"CodiceFiscale" json:"CodiceFiscale" validate:"omitempty,min=11,max=16"`
	Anagrafica    *anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
	IDFiscaleIVA  *iDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`

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

type contatti struct {
	Telefono string `xml:"Telefono" json:"Telefono" validate:"omitempty,min=5,max=12"`
	//todo: Validate Fax
	Fax   string `xml:"Fax" json:"Fax"`
	Email string `xml:"Email" json:"Email" validate:"omitempty,email,max=255"`
}

type contattiTrasmittente struct {
	Telefono string `xml:"Telefono" json:"Telefono" validate:"omitempty,min=5,max=12"`
	Email    string `xml:"Email" json:"Email" validate:"omitempty,email,max=255"`
}

// CedentePrestatore rappresent your company
type CedentePrestatore struct {
	DatiAnagrafici *datiAnagrafici `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           *indirizzoType  `xml:"Sede" json:"Sede"`
	Contatti       *contatti       `xml:"Contatti" json:"Contatti"`
}

//TODO:
//Se io NON sono un residente, devo
//valorizzare i blocchi: StabileOrganizzazione, RappresentanteFiscale

// CessionarioCommittente this struct rappresent client/customers
type CessionarioCommittente struct {
	DatiAnagrafici *datiAnagrafici `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           *indirizzoType  `xml:"Sede" json:"Sede"`

	// Da valorizzare solo se il cessionario non è un residente
	// quindi per l'operazione deve obbligatoriamente indicare uno stabile
	// in questa struttura deve indicare lo stabile (che risiede nel territorio italiano)
	StabileOrganizzazione *indirizzoType         `xml:"StabileOrganizzazione" json:"StabileOrganizzazione"`
	RappresentanteFiscale *rappresentanteFiscale `xml:"RappresentanteFiscale" json:"RappresentanteFiscale"`
}

type rappresentanteFiscale struct {
	anagrafica
	IDFiscaleIva *iDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
}
type terzoIntermediarioOSoggettoEmittente struct {
	IDFiscaleIVA  *iDFiscaleIVA `xml:"IdFiscaleIVA" json:"IdFiscaleIVA"`
	CodiceFiscale string        ` xml:"CodiceFiscale" json:"CodiceFiscale" validate:"omitempty,min=11,max=16"`
	Anagrafica    *anagrafica   `xml:"Anagrafica" json:"Anagrafica"`
}

//FatturaElettronicaHeader A tree of structures that represents the Header of an Italian electronic bill
type FatturaElettronicaHeader struct {
	DatiTrasmissione                     *datiTrasmissione                     `xml:"DatiTrasmissione" json:"DatiTrasmissione"`
	CedentePrestatore                    *CedentePrestatore                    `xml:"CedentePrestatore" json:"CedentePrestatore"`
	RappresentanteFiscale                *rappresentanteFiscale                `xml:"RappresentanteFiscale" json:"RappresentanteFiscale" validate:"omitempty"`
	CessionarioCommittente               *CessionarioCommittente               `xml:"CessionarioCommittente" json:"CessionarioCommittente"`
	TerzoIntermediarioOSoggettoEmittente *terzoIntermediarioOSoggettoEmittente `xml:"TerzoIntermediarioOSoggettoEmittente" json:"TerzoIntermediarioOSoggettoEmittente" validate:"omitempty"`
	SoggettoEmittente                    string                                `xml:"SoggettoEmittente" json:"SoggettoEmittente" validate:"omitempty,len=2,oneof=CC CZ"`
}

// Validate Check the correctness of the header according to Italian SDi
func (v FatturaElettronicaHeader) Validate() error {
	validate := getValidator()
	if err := validate.Struct(v); err != nil {
		return fmt.Errorf("FatturaElettronicaHeader %s", err)
	}
	return nil
}
