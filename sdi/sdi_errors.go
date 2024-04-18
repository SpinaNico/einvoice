package sdi

type ErrorSDICode int

const (
	NoSDIError ErrorSDICode = 0
)

// Raccolta di tutti gli errori SDI
var ErrorsMap = map[string]string{

	"00330": "l’indirizzo PEC indicato nel campo PECDestinatario  corrisponde ad una casella PEC del SdI",
	"00415": "DatiCassaPrevidenziale - L'elemento Ritenuta è stato valorizzato con SI, ma non è presente il blocco DatiRitenuta.",
	"00001": "Nome file non valido",
	"00002": "Nome file duplicato",
	"00404": "E’ stato già trasmesso un file con identico contenuto",
	"00003": "Le dimensioni del file superano quelle ammesse",
	"00102": "La firma elettronica apposta al file non risulta valida",
	"00100": "Certificato di firma scaduto",
	"00101": "Certificato di firma revocato",
	"00104": "La CA (Certification Authority) che ha emesso il certificato di firma non risulta nell’elenco delle CA affidabili",
	"00107": "Il certificato di firma non è valido ",
	"00103": "Alla firma elettronica apposta al file manca il riferimento temporale",
	"00105": "Il riferimento temporale associato alla firma elettronica apposta al file è successivo alla data di ricezione del file",
	"00106": "Il file compresso è vuoto oppure non è leggibile",
	"00200": "File non conforme al formato del problema: %s",
	"00201": "Non è possibile procedere con ulteriori controlli perché gli errori di formato presenti nel file superano il numero massimo previsto (50) ",
	"00400": `A fronte di un’aliquota pari a zero, la Natura non è stata indicata o non è stata 
			correttamente valorizzata (si ricorda che la natura N6 è compatibile con l'indicazione dell'aliquota 0 solo
			per le fatture emesse, mentre è compatibile solo con l'aliquota diversa da 0 per le fatture ricevute)
			(se nei dati di riepilogo l’aliquota IVA risulta essere uguale a zero, deve
			essere indicata la natura dell’operazione che giustifica il valore zero
			dell’aliquota. Per aliquota IVA uguale a zero la natura dell’operazione deve
			essere sempre valorizzata, con la sola esclusione, in caso di DTR, di N6).
			Per aliquota IVA diversa da zero la natura dell’operazione non deve essere
			mai valorizzata oppure, in caso di DTR, può essere valorizzata con N6)`,
	"00401": `A fronte di un’aliquota diversa da zero, è stata indicata una
			Natura non compatibile con l'operazione
			(se nei dati di riepilogo l’aliquota IVA risulta essere diversa da zero:
			o in caso di DTE la natura dell’operazione non deve essere
			valorizzata; o in caso di DTR l’unica natura compatibile con l’operazione può essere N6)`,
	"00403": `La data nel documento è futura al momento che SDI riceverà il messaggio`,
	"00411": "Hai valorizzato un dettaglioLinee in benieServizi con valore Ritenuta=SI ma non è presente il blocco DatiRitenuta obbligatorio",
	"00413": "Nella cassaPrevidenziale, non è presente la natura a fronte di una Aliquota 0",
	"00414": "Nalla cassaPrevidenziale, viene indicata la natura ma l'aliquota non è zero",

	"00417": "Non è presente un identificativo fiscale PIVA o CF",
	"00419": "DatiRiepilogo non presente in corrisponspondenza del blocco AliquotaIVA, per ogni AliquotaIVa deve essere presente il suo DatiRiepilogo, solo fatture ordinarie",
	"00420": `A fronte di EsigibilitaIVA uguale a S (Split-Payment), per
			Natura è stato indicato il valore N6
			(una operazione con natura N6 non può prevedere una modalità di
			versamento dell’IVA in regime di ‘split payment’)`,
	"00424": `L’aliquota non è indicata in termini percentuali
			(l’aliquota IVA va sempre espressa in termini percentuali; ad esempio
			un’aliquota del 10% va indicata con 10.00 e non con 0.10) `,
	"00431": `A fronte di Tipo Documento uguale a TD07 (fattura
			semplificata) o TD08 (nota di credito semplificata), gli Identificativi Fiscali e
			gli Altri Dati Identificativi del Cessionario/Committente non sono stati
			valorizzati (almeno uno dei due deve essere presente nel file)
			(per i documenti di tipo semplificato (tipo documento TD07 o TD08)
			contenuti in file di tipo DTE, è possibile inserire, per identificare la
			controparte (cessionario/committente), o gli identificativi fiscali o gli altri dati
			anagrafici o entrambi. Non è ammessa l’assenza sia degli uni che degli altri)`,
	"00432": `A fronte di Tipo Documento uguale a TD01 (fattura), TD04
			(nota di credito) o TD05 (Nota di debito), gli Identificativi Fiscali del
			Cessionario/Committente non sono stati indicati
			(per i documenti di tipo ordinario (tipo documento TD01 o TD04 o TD05)
			contenuti in file di tipo DTE, è obbligatorio inserire, per identificare la
			controparte (cessionario/committente), i suoi identificativi fiscali)`,
	"00433": `L’Imposta o l’Aliquota non sono state valorizzate (devono
			essere entrambe presenti nel file, a meno che il TipoDocumento sia uguale
			a TD07, fattura semplificata, o TD08, nota di credito semplificata)
			(solo per i documenti di tipo semplificato (tipo documento TD07 o TD08) è
			prevista la possibilità di indicare, in alternativa non esclusiva, l’aliquota IVA
			o l’imposta. In tutti gli altri casi devono essere presenti sia l’una che l’altra)`,
	"00434": `Imposta e Aliquota non coerenti
			(l’aliquota IVA e l’imposta presenti in un blocco di dati di riepilogo sono
			giudicate non coerenti quando:
			o l’aliquota IVA è pari a zero e l’imposta è diversa da zero
			o l’aliquota IVA è diversa da zero e l’imposta è pari a zero ed il valore
			zero di quest’ultima non è giustificabile con un basso valore
			dell’imponibile, cioè quando [aliquota * imponibile] / 100 > 0.00)`,
	"00435": `Detraibile e Deducibile non possono essere presenti
			contemporaneamente con riferimento agli stessi DatiRiepilogo
			(uno stesso importo non può essere riferito a spese che siano
			contemporaneamente deducibili e detraibili)`,
	"00436": `DataRegistrazione antecedente alla data del documento
			(la data di registrazione di un documento indicata nel file di tipo DTR non 
			può risultare anteriore alla data del documento stesso; questo controllo
			viene effettuato anche nel caso di tipo documento uguale a TD12 ma solo
			se presente il campo 3.2.3.1.2 <Data>)`,
	"00442": `La rettifica non è possibile perché i dati risultano già annullati
			(non è possibile rettificare il documento indicato nel file di rettifica in quanto
			risulta già annullato)`,
	"00443": `L’annullamento non è possibile perché i dati risultano già annullati
			(non è possibile annullare il documento o il file indicati nel file di
			annullamento in quanto sono già stati precedentemente annullati)`,
	"00444": `Il file originario indicato nel campo IdFile non esiste
			(nel file di rettifica/annullamento è stato indicato l’identificativo di un file che
			non esiste)`,
	"00445": `Il file indicato nel campo IdFile non è il file originario
			(il file di rettifica/annullamento deve rettificare/annullare il file originario e
			non file di rettifica/annullamento)`,
	"00446": `Posizione non trovata all’interno del file originario
			(la posizione del documento indicata nel file di rettifica/annullamento non è
			una posizione esistente all’interno del file indicato come file oggetto di
			rettifica/annullamento)`,
	"00447": `Un file di rettifica non deve contenere più di un documento
			(con un file di rettifica è possibile modificare un solo documento alla volta)`,
	"00448": `(controllo in vigore dal primo ottobre 2020)
			non è più ammesso il valore generico N2, N3 o N6 come
			codice natura dell’operazione (a partire dal primo ottobre 2020 non è più
			consentito utilizzare i codici natura ‘padre’ ma solo quelli di dettaglio,
			laddove previsti; in particolare non sono più utilizzabili i codici N2, N3 e N6)`,
	"00460": `Il Tipo Documento non è coerente con il Paese del
			Cedente/Prestatore
			(in un file di tipo DTR è possibile inserire dati relativi a fatture di acquisto
			intracomunitario di beni e/o servizi (tipo documento TD10 e TD11). Questo
			però è consentito solo se il paese della controparte (cedente/prestatore) è
			diverso da IT e rientra in uno di quelli previsti per questo tipo di operazioni
			(AT – BE – BG – CY – HR – DK – EE – FI – FR – DE – GB – EL – IE – LV
			– LT – LU – MT – NL – PL – PT – CZ – RO – SK – SI – ES – SE - HU))`,
	"00461": `Il Tipo Documento non è ammesso per le fatture emesse 
			(in un file di tipo DTE non è possibile inserire dati relativi a fatture di
			acquisto intracomunitario di beni e/o servizi (tipo documento TD10 e
			TD11))`,
	"00462": `La Data deve essere valorizzata (può non esserlo solo per
			TipoDocumento uguale a TD12, documento riepilogativo)
			(in un file di tipo DTR la data del documento deve essere sempre indicata,
			a meno che non si tratti di documento riepilogativo (TD12) nel qual caso è
			opzionale)`,
	"00464": `Gli Identificativi Fiscali del Cedente/Prestatore non sono stati
			valorizzati (tale circostanza è ammissibile solo a fronte di TipoDocumento
			uguale a TD12, documento riepilogativo)
			(in un file di tipo DTR gli identificativi fiscali del cedente/prestatore devono
			essere sempre indicati, a meno che non siano stati riportati soltanto dati di
			documenti riepilogativi (TD12))`,
	"00467": `Con riferimento allo stesso blocco Cessionario/Committente,
			sono stati riportati i dati di documenti riepilogativi insieme ad altri tipi di
			documento
			(in un file di tipo DTE non possono essere contemporaneamente presenti,
			per uno stessa controparte (cessionario/committente), dati relativi a
			documenti riepilogativi (TD12) e dati relativi ad altre tipologie di documento;
			la presenza di un documento TD12 non ammette, con riferimento alla
			stessa controparte (cessionario/committente), tipologie di documento
			diverse da TD12)`,
	"00468": `Con riferimento allo stesso blocco Cedente/Prestatore, sono
			stati riportati i dati di documenti riepilogativi insieme ad altri tipi di
			documento
			(in un file di tipo DTR non possono essere contemporaneamente presenti,
			per uno stessa controparte (cedente/prestatore), dati relativi a documenti
			riepilogativi (TD12) e dati relativi ad altre tipologie di documento; la
			presenza di un documento TD12 non ammette, con riferimento alla stessa
			controparte (cedente/prestatore), tipologie di documento diverse da TD12)`,
	"00469": `Il soggetto che ha firmato il file non corrisponde né al
			firmatario del file originario né al soggetto titolare dei dati contenuti nel file
			originario
			(il file di rettifica o di annullamento deve essere firmato dallo stesso
			soggetto che ha firmato il file originario oggetto di rettifica o annullamento 
			158
			oppure deve essere firmato dallo stesso soggetto (o suo incaricato) che,
			nel file originario oggetto di rettifica o annullamento, figura come
			cedente/prestatore (in caso di DTE) o come cessionario/committente (in
			caso di DTR))`,
	"00470": `E’ stata modificata la partita IVA e/o il codice fiscale del
			soggetto titolare dei dati contenuti nel file originario
			(il file di rettifica non può contenere identificativi fiscali del
			cedente/prestatore (in caso di DTE) o del cessionario/committente (in caso
			di DTR) diversi da quelli contenuti nel file originario oggetto di rettifica)`,
	"00301": "La partita IVA del Cedente/Prestatore non è valida",
	"00302": "Il Codice Fiscale del Cedente/Prestatore non è valido",
	"00303": "La partita IVA del Rappresentante Fiscale non è valida",
	"00305": "La partita IVA del Cessionario/Committente non è valida",
	"00306": "Il Codice Fiscale del Cessionario/Committente non è valido",
	"00320": "Il codice fiscale, non fa parte del gruppo IVA indicato",
	"00321": `In presenza di una partita IVA di gruppo IVA del Cedente/Prestatore occorre valorizzare il Codice Fiscale del 
			Cedente/Prestatore con quello del soggetto partecipante al gruppo`,
	"00322": "Il codice fiscale è presente ma non è partecipe al gruppo IVA",
	"00325": `In presenza di una partita IVA di gruppo IVA del
			Cessionario/Committente occorre valorizzare il Codice Fiscale del
			Cessionario/Committente con quello del soggetto partecipante al gruppo`,
	"00600": "Soggetto non autorizzato alla trasmissione",

	"00500": `Partita IVA del Cedente/Prestatore cessata in Anagrafe Tributaria (il controllo viene effettuato 
			solo per tipo documento diverso da TD12 oppure, se uguale a TD12, solo per dati fatture emesse (DTE))`,
	"00501": `Partita IVA del Cessionario/Committente cessata in Anagrafe Tributaria (il controllo viene effettuato 
			solo per tipo documento diverso da TD12 oppure, se uguale a TD12, solo per dati fatture ricevute (DTR))`,
	"00502": `Partita IVA del Rappresentante Fiscale cessata in Anagrafe Tributaria`,
	"00503": `La data del documento non è compatibile con il periodo di riferimento
			(il controllo produce segnalazione quando un file, ricevuto entro il termine
			ultimo stabilito per l’invio dei dati fattura relativi al periodo di riferimento,
			presenta il campo Data antecedente al termine iniziale del periodo stesso)`,
	"00504": `La data di registrazione del documento non è compatibile con
			il periodo di riferimento
			(il controllo produce segnalazione quando un file, ricevuto entro il termine
			ultimo stabilito per l’invio dei dati fattura relativi al periodo di riferimento,
			presenta il campo DataRegistrazione antecedente al termine iniziale del
			periodo stesso)`,
}
