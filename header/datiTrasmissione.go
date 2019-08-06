package header

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiTrasmissione struct {
	IDTrasmittente share.IDFiscaleIVA `xml:"IdTrasmittente" json:"IdTrasmittente"`
	//ProgressivoInvio: progressivo che il soggetto trasmittente attribuisce
	//al file che inoltra al Sistema di Interscambio per una propria finalità di
	//identificazione univoca.
	ProgressivoInvio progressivoInvio `xml:"ProgressivoInvio" json:"ProgressivoInvio"`
	//FormatoTrasmissione: codice identificativo del tipo di trasmissione
	//che si sta effettuando e del relativo formato. Va sempre valorizzato
	//con “FPR12”.
	FormatoTrasmissione formatoTrasmissione `xml:"FormatoTrasmissione" json:"FormatoTrasmissione"`
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
	CodiceDestinatario   codiceDestinatario   `xml:"CodiceDestinatario" json:"CodiceDestinatario"`
	ContattiTrasmittente contattiTrasmittente `xml:"ContattiTrasmittente" json:"ContattiTrasmittente"`

	//PECDestinatario: indirizzo di Posta Elettronica Certificata al quale, se
	//valorizzato, viene recapitata la fattura nei casi in cui il valore di
	//CodiceDestinatario sia uguale a ‘0000000’ e non risulti registrato alcun
	//canale telematico associato alla partita IVA del cessionario/committente.
	PECDestinatario pecDestinatario `xml:"PECDestinatario" json:"PECDestinatario"`
}

func (c datiTrasmissione) Validate() error {
	var err error
	err = c.IDTrasmittente.Validate()
	if err != nil {
		return fmt.Errorf("DatiTrasmissione %s", err)
	}

	err = c.PECDestinatario.Validate()
	if err != nil {
		return fmt.Errorf("DatiTrasmissione %s", err)
	}

	err = c.ProgressivoInvio.Validate()
	if err != nil {
		return fmt.Errorf("DatiTrasmissione %s", err)
	}

	err = c.CodiceDestinatario.Validate()
	if err != nil {
		return fmt.Errorf("DatiTrasmissione %s", err)
	}

	err = c.FormatoTrasmissione.Validate()
	if err != nil {
		return fmt.Errorf("DatiTrasmissione %s", err)
	}

	if len(string(c.PECDestinatario)) == 0 && string(c.CodiceDestinatario) == "0000000" {
		return fmt.Errorf("DatiTrasmissione: You can't give me a Pec, with the recipient code \"0000000\" you have to give me a code, or the PEC")
	}

	return nil
}
