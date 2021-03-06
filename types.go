package einvoice

var TipoCessionePrestazione = map[string]string{
	"SC": "Sconto",
	"PR": "Premio",
	"AB": "Abbuono",
	"AC": "Spesa accessoria",
}
var TipoDocumento = map[string]string{
	"TD01": "Fattura",
	"TD02": "Acconto/Anticipo su fattura",
	"TD03": "Acconto/Anticipo su parcella",

	"TD04": "Nota di credito",
	"TD05": "Nota di debito",
	"TD06": "Parcella",

	"TD07": "Fattura semplificata",
	"TD08": "Nota di Credito semplificata",
	"TD10": "Fattura per acquisto intracomunitario beni",
	"TD11": "Fattura per acquisto intracomunitario servizi",
	"TD12": "Documento riepilogativo (art.6,DPR 695/1996)",

	"TD16": "Integrazione fattura reverse charge interno",
	"TD17": "Integrazione/autofattura per acquisto servizi dall’estero",
	"TD18": "Integrazione per acquisto di beni intracomunitari",
	"TD19": "Integrazione/autofattura per acquisto di beni ex art.17 c.2 DPR 633/72",
	"TD20": "Autofattura per regolarizzazione e integrazione delle fatture (art.6 c.8 d.lgs. 471/97 o art.46 c.5 D.L. 331/93)",
	"TD21": "Autofattura per splafonamento",
	"TD22": "Estrazione beni da Deposito IVA",
	"TD23": "Estrazione beni da Deposito IVA con versamento dell’IVA",
	"TD24": "Fattura differita di cui all’art.21, comma 4, lett. a)",
	"TD25": "Fattura differita di cui all’art.21, comma 4, terzo periodo lett. b)",
	"TD26": "Cessione di beni ammortizzabili e per passaggi interni (ex art.36 DPR 633/72)",
	"TD27": "Fattura per autoconsumo o per cessioni gratuite senza rivalsa",
}

var Natura = map[string]string{
	"N1":   "escluse ex art.15",
	"N2":   "non soggette",
	"N2.1": "non soggette ad IVA ai sensi degli artt. Da 7 a 7-septies del DPR 633/72",
	"N2.2": "non soggette – altri casi",
	"N3":   "non imponibili",
	"N3.1": "non imponibili – esportazioni",
	"N3.2": "non imponibili – cessioni intracomunitarie",
	"N3.3": "non imponibili – cessioni verso San Marino",
	"N3.4": "non imponibili – operazioni assimilate alle cessioni all’esportazione",
	"N3.5": "non imponibili – a seguito di dichiarazioni d’intento",
	"N3.6": "non imponibili – altre operazioni che non concorrono alla formazione del plafond",
	"N4":   "esenti",
	"N5":   "regime del margine / IVA non esposta in fattura",
	"N6":   "inversione contabile (per le operazioni in reverse charge ovvero nei casi di autofatturazione per acquisti extra UE di servizi ovvero per importazioni di beni nei soli casi previsti)",
	"N6.1": "inversione contabile – cessione di rottami e altri materiali di recupero",
	"N6.2": "inversione contabile – cessione di oro e argento puro",
	"N6.3": "inversione contabile – subappalto nel settore edile",
	"N6.4": "inversione contabile – cessione di fabbricati",
	"N6.5": "inversione contabile – cessione di telefoni cellulari",
	"N6.6": "inversione contabile – cessione di prodotti elettronici",
	"N6.7": "inversione contabile – prestazioni comparto edile e settori connessi",
	"N6.8": "inversione contabile – operazioni settore energetico",
	"N6.9": "inversione contabile – altri casi",
	"N7":   "IVA assolta in altro stato UE (vendite a distanza ex art. 40 commi 3 e 4 e art. 41 comma 1 lett. b, DL 331/93; prestazione di servizi di telecomunicazioni, tele-radiodiffusione ed elettronici ex art. 7-sexies lett. f, g, DPR 633/72 e art. 74-sexies, DPR 633/72)",
}

var CondizioniPagamento = map[string]string{
	"TP01": "pagamento a rate",
	"TP02": "pagamento completo",
	"TP03": "anticipo",
}

var MetodiPagamento = map[string]string{
	"MP01": "contanti",
	"MP02": "assegno",
	"MP03": "assegno circolare",
	"MP04": "contanti presso Tesoreria",
	"MP05": "bonifico",
	"MP06": "vaglia cambiario",
	"MP07": "bollettino bancario",
	"MP08": "carta di pagamento",
	"MP09": "RID",
	"MP10": "RID utenze",
	"MP11": "RID veloce",
	"MP12": "Riba",
	"MP13": "MAV",
	"MP14": "quietanza erario stato",
	"MP15": "giroconto su conti di contabilità speciale",
	"MP16": "domiciliazione bancaria",
	"MP17": "domiciliazione postale",
	"MP18": "bollettino di c/c postale",
	"MP19": "SEPA Direct Debit",
	"MP20": "SEPA Direct Debit CORE",
	"MP21": "SEPA Direct Debit B2B",
	"MP22": "Trattenuta su somme già riscosse",
	"MP23": "PagoPA",
}

var RegimeFiscale = map[string]string{
	"RF01": "Ordinario;",
	"RF02": "Contribuenti minimi (art. 1, c.96-117, L. 244/2007);",
	"RF04": "Agricoltura e attività connesse e pesca (artt. 34 e 34-bis, D.P.R. 633/1972);",
	"RF05": "Vendita sali e tabacchi (art. 74, c.1, D.P.R. 633/1972);",
	"RF06": "Commercio dei fiammiferi (art. 74, c.1, D.P.R. 633/1972);",
	"RF07": "Editoria (art. 74, c.1, D.P.R. 633/1972);",
	"RF08": "Gestione di servizi di telefonia pubblica (art. 74, c.1, D.P.R. 633/1972);",
	"RF09": "Rivendita di documenti di trasporto pubblico e di sosta (art. 74, c.1, D.P.R. 633/1972);",
	"RF10": "Intrattenimenti, giochi e altre attività di cui alla tariffa allegata al D.P.R. n. 640/72 (art. 74, c.6, D.P.R. 633/1972);",
	"RF11": "Agenzie di viaggi e turismo (art. 74-ter, D.P.R. 633/1972);",
	"RF12": "633/1972); Agriturismo (art. 5, c.2, L. 413/1991);",
	"RF13": "Vendite a domicilio (art. 25-bis, c.6, D.P.R. 600/1973);",
	"RF14": "Rivendita di beni usati, di oggetti d’arte, d’antiquariato o da collezione (art. 36, D.L. 41/1995);",
	"RF15": "Agenzie di vendite all’asta di oggetti d’arte, antiquariato o da collezione (art. 40-bis, D.L. 41/1995);",
	"RF17": "IVA per cassa (art. 32-bis, D.L. 83/2012);",
	"RF18": "Altro;",
	"RF19": "Forfettario (art.1, c. 54-89, L. 190/2014)",
}

var SocioUnico = map[string]string{
	"SU": "La società è a socio unico",
	"SM": "La società non è a socio unico",
}

var StatoLiquidazione = map[string]string{
	"LS": "La società è in stato di liquidazione",
	"LN": "La società non è in stato di liquidazione",
}

var SoggettoEmittente = map[string]string{
	"CC": "Cessionario/committente",
	"TZ": "Soggetto terzo",
}

var TipoRitenuta = map[string]string{
	"RT01": "Ritenuta persone fisiche",
	"RT02": "Ritenuta persone giuridiche",
	"RT03": "Contributo INPS",
	"RT04": "Contributo ENASARCO",
	"RT05": "Contributo ENPAM",
	"RT06": "Altro contributo previdenziale",
}

var TipoCassa = map[string]string{

	"TC01": "Cassa Nazionale Previdenza e Assistenza Avvocati e Procuratori Legali",
	"TC02": "Cassa Previdenza Dottori Commercialisti",
	"TC03": "Cassa Previdenza e Assistenza Geometri",
	"TC04": "Cassa Nazionale Previdenza e Assistenza Ingegneri e Architetti Liberi Professionisti",
	"TC05": "Cassa Nazionale del Notariato",
	"TC06": "Cassa Nazionale Previdenza e Assistenza Ragionieri e Periti Commerciali",
	"TC07": "Ente Nazionale Assistenza Agenti e Rappresentanti di Commercio (ENASARCO)",
	"TC08": "Ente Nazionale Previdenza e Assistenza Consulenti del Lavoro (ENPACL)",
	"TC09": "Ente Nazionale Previdenza e Assistenza Medici (ENPAM)",
	"TC10": "Ente Nazionale Previdenza e Assistenza Farmacisti (ENPAF)",
	"TC11": "Ente Nazionale Previdenza e Assistenza Veterinari (ENPAV)",
	"TC12": "Ente Nazionale Previdenza e Assistenza Impiegati dell'Agricoltura (ENPAIA)",
	"TC13": "Fondo Previdenza Impiegati Imprese di Spedizione e Agenzie Marittime",
	"TC14": "Istituto Nazionale Previdenza Giornalisti Italiani (INPGI)",
	"TC15": "Opera Nazionale Assistenza Orfani Sanitari Italiani (ONAOSI)",
	"TC16": "Autonoma Assistenza Integrativa Giornalisti Italiani (CASAGIT)",
	"TC17": "Ente Previdenza Periti Industriali e Periti Industriali Laureati (EPPI)",
	"TC18": "Ente Previdenza e Assistenza Pluricategoriale (EPAP)",
	"TC19": "Ente Nazionale Previdenza e Assistenza Biologi (ENPAB)",
	"TC20": "Ente Nazionale Previdenza e Assistenza Professione Infermieristica (ENPAPI)",
	"TC21": "Ente Nazionale Previdenzae Assistenza Psicologi (ENPAP)",
	"TC22": "INPS",
}

var TipoScontoMaggiorazione = map[string]string{
	"SC": "Sconto",
	"MG": "Maggiorazione",
}

var EsigibilitaIVA = map[string]string{
	"I": "IVA ad esigibilità immediata",
	"D": "IVA ad esigibilità differita",
	"S": "scissione dei pagamenti",
}
