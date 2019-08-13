package body

import (
	"fmt"

	"github.com/SpinaNico/go-struct-invoice/share"
)

type datiTrasporto struct {
	DatiAnagraficiVettore datiAnagraficiVettore `xml:"DatiAnagraficiVettore" json:"DatiAnagraficiVettore"`

	//formato alfanumerico; lunghezza massima di 80  caratteri.
	MezzoTrasporto string `xml:"MezzoTrasporto" json:"MezzoTrasporto"`
	// formato alfanumerico; lunghezza massima di 100 caratteri.
	CausaleTrasporto string `xml:"CausaleTrasporto" json:"CausaleTrasporto"`
	//NumeroColli: formato numerico; lunghezza massima di 4 caratteri.
	NumeroColli int `xml:"NumeroColli" json:"NumeroColli"`
	//Descrizione: formato alfanumerico; lunghezza massima di 100 caratteri.
	Descrizione string `xml:"Descrizione" json:"Descrizione"`
	//UnitaMisuraPeso: formato alfanumerico; lunghezza massima di 10 caratteri.
	UnitaMisuraPeso string `xml:"UnitaMisuraPeso" json:"UnitaMisuraPeso"`
	//PesoLordo: formato numerico La sua lunghezza va da 4 a 7 caratteri.
	PesoLordo decimale2 `xml:"PesoLordo" json:"PesoLordo"`
	// PesoNetto: formato numerico  La sua lunghezza va da 4 a 7 caratteri.
	PesoNetto decimale2 `xml:"PesoNetto" json:"PesoNetto"`
	//DataOraRitiro: la data e ora, devono essere in questo formato YYYY-MM- DDTHH:MM:SS.
	DataOraRitiro dataOra `xml:"DataOraRitiro" json:"DataOraRitiro"`
	// DataInizioTrasporto: la data deve essere  di questo formato YYYY-MM-DD.
	DataInizioTrasporto data `xml:"DataInizioTrasporto" json:"DataInizioTrasporto"`

	// TipoResa: codifica del termine di resa (Incoterms) espresso
	// secondo lo standard ICC-Camera di Commercio Internazionale
	// (formato alfanumerico di 3 caratteri)
	//
	//  **EXW**: FRANCO FABBRICA "Franco Fabbrica"
	//  significa che il venditore effettua la consegna mettendo
	//  la merce a disposizione del compratore nei propri
	//  locali o in altro luogo convenuto (stabilimento,
	//  fabbrica, magazzino, etc). Il venditore non ha l'obbligo
	//  di caricare la merce sul veicolo di prelevamento,
	//  n� di sdoganarla all'esportazione, nel caso in
	//
	//  **FCA**: FRANCO VETTORE "Franco Vettore"
	//  significa che il venditore effettua la consegna
	//  rimettendo la merce al vettore o ad altra persona designata
	//  dal compratore nei propri locali o in altro luogo
	//  convenuto. Si raccomanda alle parti di specificare il pi�
	//  chiaramente possibile il punto nel luogo di consegna
	//  convenuto, poich� il risciho passa al compratore in tale
	//  punto. Ove le parti desiderino che la merce sia
	//  consegnata nei locali del venditore, esse devono indicare
	//  l'indirizzo di tali locali quale luogo di consegna
	//  convenuto. Se, invece, le parti desiderano che la merce sia
	//  consegnata in altro luogo, esse devono indicare un
	//  diverso, specifico luogo di consegna. FCA richiede che
	//  il venditore, se il caso, sdogani la merce
	//
	//  **CPT**: TRASPORTO PAGATO FINO A "Trasporto pagato
	//  fino a" significa che il venditore effettua la
	//  consegna rimettendo la merce al vettore o ad altra
	//  persona designata dallo stesso venditore in un luogo
	//  concordato (se tale luogo � stato concordato tra le parti) e
	//  che il venditore deve stipulare il contratto di
	//  trasporto e sopportare le spese necessarie per l'invio
	//  della merce al luogo di destinazione contenuto. Il
	//  venditore adempie la sua obbligazione di effettuare la
	//  consegna quando rimette la merce al vettore e non quando
	//  la merce arriva al luogo di destinazione. CPT
	//  richiede che il venditore, se il caso, sdogani la merce
	//
	//  **CIP**: TRASPORTO E ASSICURAZIONE PAGATI FINO A
	//  "Trasporto e Assicurazione Pagati fino a" significa che il
	//  venditore effettua la consegna rimettendo la merce al
	//  vettore o ad altra persona da lui stesso designata in un
	//  luogo concordato (se tale luogo � stato concordato
	//  fra le parti) e che il venditore deve stipulare il
	//  contratto di trasporto e sopportare le spese necessarie
	//  per l'invio della merce al luogo di destinazione
	//  convenuto. Il venditore provvede anche ad una copertura
	//  assicurativa contro il rischio del compratore di perdita o
	//  danni alla merce durante il trasporto. Il compratore
	//  deve tenere presente che secondo la regola CIP il
	//  venditore � obbligato ad ottenere soltanto una copertura
	//  assicurativa minima. Ove il compratore desideri avere una
	//  protezione assicurativa pi� ampia, dovr� accordarsi
	//  espressamente con il venditore o provvedere direttamente ad
	//  un'assicurazione integrativa. Il venditore adempie la sua
	//  obbligazione di effettuare la consegna quando rimette la
	//  merce al vettore e non quando la merce arriva al luogo
	//  di destinazione. Si raccomanda alle parti di
	//  specificare il pi� chiaramente possibile il punto nel luogo
	//  di destinazione convenuto, poich� le spese fino a
	//  tale punto sono a carico del venditore. CIP richiede
	//  che il venditore, se il caso, sdogani la merce
	//
	//  **DAT**: RESO AL TERMINAL "Reso al Terminal"
	//  significa che il venditore effettua la consegna mettendo
	//  la merce, una volta scaricata dal mezzo di
	//  trasporto di arrivo, a disposizione del compratore al
	//  terminal convenuto nel porto o luogo di destinazione
	//  concordato. "Terminal" include ogni luogo, coperto o
	//  scoperto, come una banchina, un magazzino, un piazzale
	//  per container, un terminal stradale, ferroviario
	//  o aeroportuale. Il venditore sopporta tutti i
	//  rischi connessi al trasporto ed alla scaricazione
	//  della merce al terminal nel porto o luogo di
	//  destinazione convenuto. Si raccomanda alle parti di
	//  specificare il pi� chiaramente possibile il terminal e, se
	//  possibile, un punto specifico all'interno del terminal
	//  nel porto o luogo di destinazione convenuto poich�
	//  i rischi fino a tale punto sono a carico del
	//  venditore. DAT richiede che il venditore, se il caso,
	//
	//  **DAP**: RESO A LUOGO DI DESTINAZIONE "Reso a Luogo
	//  di Destinazione" significa che il venditore
	//  effettua la consegna mettendo la merce a disposizione
	//  del compratore sul mezzo di trasporto di arrivo
	//  pronta per la scaricazione nel luogo di destinazione
	//  convenuto. Il venditore sopporta tutti i rischi connessi
	//  al trasporto della merce al luogo convenuto. Si
	//  raccomanda alle poarti di specificare il pi� chiaramente
	//  possibile il punto nel luogo di destinazione convenuto,
	//  poich� i rischi fino a tale punto sono a carico del
	//  venditore. DAP richiede che il venditore, se il caso,
	//  sdogani la merce all'esportazione ma NON
	//
	//  **DDP**: RESO SDOGANATO La presente regola pu�
	//  essere utilizzata indipendentemente dal modo di
	//  trasporto scelto e pu� essere usata anche nel caso in cui si
	//  utilizzi pi� di un modo di trasporto. "Reso sdoganato"
	//  significa che il venditore effettua la consegna mettendo
	//  la merce a disposizione del compratore,
	//  sdoganata all'importazione, sul mezzo di trasporto di
	//  arrivo pronta per la scaricazione nel luogo di
	//  destinazione convenuto. Il venditore supporta tutte le
	//  spese ed i rischi connessi al trasporto della merce al
	//  luogo di destinazione ed ha l'obbligo di sdoganare la
	//  merce non solo all'esportazione ma anche
	//  all'importazione, di pagare eventuali diritti sia di
	//  esportazione sia d'importazione ed espletare tutte le
	//  formalit� doganali. Il DDP comporta l livello massimo di
	//
	//  **FAS**: FRANCO LUNGO BORDO La presente regola pu�
	//  essere utilizzata esclusivamente in caso di
	//  trasporto marittimo o per vie d'acqua interne. "Franco
	//  Lungo Bordo" significa che il venditore effettua la
	//  consegna mettendo la merce sottobordo della nave (ad es.
	//  su una banchina o una chiatta) designata dal
	//  compratore nel porto d'imbarco convenuto. Il rischio di
	//  perdita o di danni alla merce passa quando la merce �
	//  sottobordo della nave ed il compratore sopporta tutte le
	//  spese da tale momento in avanti. Si raccomanda alle
	//  parti di specificare il pi� chiaramente possibile il
	//  punto di caricazione nel porto d'imbarco convenuto,
	//  poch� le spese ed i rischi fino a tale punto sono a
	//  carico del venditore e tali spese e gli oneri di
	//  movimentazione connessi possono variare a seconda degli usi
	//  del porto. Se la merce � containerizzata, � d'uso
	//  per il venditore consegnare la merce al vettore al
	//  terminal e non sottobordo della nave. In tali
	//  situazioni, la regola FAS risulterebbe inappropriata e si
	//
	//  **FOB**: FRANCO A BORDO La presente regola pu�
	//  essere utilizzata esclusivamente in caso di
	//  trasporto marittimo o per vie d'acqua interne. "Franco a
	//  Bordo" significa che il venditore effettua la
	//  consegna mettendo la merce a bordo della nave designata
	//  dal compratore nel porto d'imbarco convenuto o
	//  procurando la merce gi� cos� consegnata. Il rischio di
	//  perdita o di danni alla merce passa quando la merce � a
	//  bordo della nave ed il compratore sopporta tutte le
	//  spese da tale momento in avanti. FOB richiede che il
	//  venditore, se il caso, sdogani la merce
	//
	//  **CFR**: COSTO E NOLO La presente regola pu� essere
	//  utilizzata esclusivamente in caso di trasporto marittimo
	//  o per vie d'acqua interne. "Costo e Nolo"
	//  significa che il venditore effettua la consegna mettendo
	//  la merce a bordo della nave o procurando la merce
	//  gi� cos� consegnata. Il rischio di perdita o di
	//  danni alla merce passa quando la merce � a bordo della
	//  nave. Il venditore deve stipulare il contratto di
	//  trasporto e sopportare le spese necessarie per l'invio
	//  della merce al porto di destinazione convenuto.
	//  Quando si utilizzano CPT, CIP, CFR o CIF, il venditore
	//  adempie la sua obbligazione di effettuare la consegna
	//  quando rimette la merce al vettore secondo quanto
	//  specificato dalla regola scelta e non quando la merce arriva
	//  al luogo di destinazione. Ne deriva che il
	//  passaggio del rischio ed il trasferimento delle spese
	//  avvengono in luoghi diversi. Mentre il contratto
	//  specificher� sempre un porto di destinazione, esso potrebbe
	//  non specificare il porto d'imbarco, ove il rischio
	//  passa al compratore. Se il porto d'imbarco presenta
	//  un particolare interesse per il compratore, si
	//  raccomanda alle parti di specificarlo il pi� chiaramente
	//  possibile nel contratto. CFR richiede che il venditore,
	//
	//  **CIF**: COSTO, ASSICURAZIONE E NOLO La presente
	//  regola può essere utilizzata esclusivamente in caso
	//  di trasporto marittimo o per via d'acqua interne.
	//  "Costo, Assicurazione e Nolo" significa che il
	//  venditore effettuala consegna mettendo la merce a bordo
	//  della nava o procurando la merce gi� consegnata. Il
	//  rischio di perdita o di danni alla merce passa quando la
	//  merce � a bordo della nave. Il venditore deve
	//  stipulare il contratto di trasporto e sopportare le spese
	//  necessarie per l'invio della merce al porto di
	//  destinazione convenuto. Il venditore provvede anche ad una
	//  copertura assicurativa contro il rischio del compratore
	//  di perdita o danni alla merce durante il
	//  trasporto. Il compratore deve tener presente che secondo
	//  la regola CIF il venditore � obbligato ad ottenere
	//  soltanto una copertuira assicurativa minima. Ove il
	//  compratore desideri avere una protezione assicurativa
	//  pi� ampia, dovr� accordarsi espressamente con il
	//  venditore o provvedere direttamente ad un
	//  'assicurazione integrativa. Quando si utilizzano
	//  CPT,CIP,CFR o CIF, il venditore adempie la sua obbligazione
	//  di effetture la consegna quando rimette la merce
	//  al vettore secondo quanto specificato dalla
	//  regola scelta e non quando la merce arriva al luogo di
	TipoResa string `xml:"TipoResa" json:"TipoResa"`

	//IndirizzoResa che si compone degli stessi campi previsti per
	//l’elemento Sede del CedentePrestatore contenuti nel tipo
	IndirizzoResa share.IndirizzoType `xml:"IndirizzoResa" json:"IndirizzoResa"`

	// YYYY-MM-DDTHH:MM:SS.
	DataOraConsegna dataOra `xml:"DataOraConsegna" json:"DataOraConsegna"`
}

// Validate ...
func (f datiTrasporto) Validate() error {
	if err := f.DatiAnagraficiVettore.Validate(); err != nil {
		return fmt.Errorf("DatiAnagraficiVettore %s", err)
	}

	if len(f.MezzoTrasporto) > 80 {
		return fmt.Errorf("MezzoTrasporto %s", share.ErrorMaxLength(80))
	}
	if len(f.CausaleTrasporto) > 100 {
		return fmt.Errorf("CausaleTrasporto %s", share.ErrorMaxLength(100))
	}

	if f.NumeroColli > 9999 {
		return fmt.Errorf("NumeroColli %s", share.ErrorIncluded(1, 9999))
	}

	if len(f.Descrizione) > 100 {
		return fmt.Errorf("Descrizione %s", share.ErrorMaxLength(100))
	}

	if len(f.UnitaMisuraPeso) > 10 {
		return fmt.Errorf("UnitaMisuraPeso %s", share.ErrorMaxLength(10))
	}
	// PesoLordo
	if err := f.PesoLordo.Validate(); err != nil {
		return fmt.Errorf("PesoLordo %s", err)
	}
	if len(f.PesoLordo) < 4 || len(f.PesoLordo) > 7 {
		return fmt.Errorf("PesoLordo %s", share.ErrorIncluded(4, 7))
	}

	// PesoNetto
	if err := f.PesoNetto.Validate(); err != nil {
		return fmt.Errorf("PesoNetto %s", err)
	}
	if len(f.PesoNetto) < 4 || len(f.PesoNetto) > 7 {
		return fmt.Errorf("PesoNetto %s", share.ErrorIncluded(4, 7))
	}

	if err := f.DataOraRitiro.Validate(); err != nil {
		return fmt.Errorf("DataOraRitiro %s", err)
	}

	if err := f.DataInizioTrasporto.Validate(); err != nil {
		return fmt.Errorf("DataInizioTrasporto %s", err)
	}

	switch f.TipoResa {
	case "EXW":
		break
	case "FCA":
		break
	case "CPT":
		break
	case "CIP":
		break
	case "DAT":
		break
	case "DAP":
		break
	case "DDP":
		break
	case "FAS":
		break
	case "FOB":
		break
	case "CFR":
		break
	case "CIF":
		break
	default:
		return fmt.Errorf("TipoResa %s", share.ErrorIncorrectValue(f.TipoResa))
	}

	if err := f.IndirizzoResa.Validate(); err != nil {
		return fmt.Errorf("IndirizzoResa %s", err)
	}

	if err := f.DataOraConsegna.Validate(); err != nil {
		return fmt.Errorf("DataOraConsegna %s", err)
	}
	return nil
}
