package header

type datiTrasmissione struct {
	IDTrasmittente idTrasmittente `xml:"IdTrasmittente" json:"IdTrasmittente"`
	//ProgressivoInvio: progressivo che il soggetto trasmittente attribuisce
	//al file che inoltra al Sistema di Interscambio per una propria finalità di
	//identificazione univoca.
	ProgressivoInvio string `xml:"ProgressivoInvio" json:"ProgressivoInvio"`
	//FormatoTrasmissione: codice identificativo del tipo di trasmissione
	//che si sta effettuando e del relativo formato. Va sempre valorizzato
	//con “FPR12”.
	FormatoTrasmissione string `xml:"FormatoTrasmissione" json:"FormatoTrasmissione"`
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
	CodiceDestinatario   string               `xml:"CodiceDestinatario" json:"CodiceDestinatario"`
	ContattiTrasmittente contattiTrasmittente `xml:"ContattiTrasmittente" json:"ContattiTrasmittente"`

	//PECDestinatario: indirizzo di Posta Elettronica Certificata al quale, se
	//valorizzato, viene recapitata la fattura nei casi in cui il valore di
	//CodiceDestinatario sia uguale a ‘0000000’ e non risulti registrato alcun
	//canale telematico associato alla partita IVA del cessionario/committente.
	PECDestinatario string `xml:"PECDestinatario" json:"PECDestinatario"`
}

func (c datiTrasmissione) Validate() error {
	return nil
}
