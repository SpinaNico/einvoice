	// TipoCassa: cassa di previdenza della categoria
	// professionale di appartenenza
	// Valori possibili per questo campo:
	//
	// + TC01 Cassa Nazionale Previdenza e Assistenza Avvocati e Procuratori Legali
	// + TC02 Cassa Previdenza Dottori Commercialisti
	// + TC03 Cassa Previdenza e Assistenza Geometri
	// + TC04 Cassa Nazionale Previdenza e Assistenza Ingegneri e Architetti Liberi Professionisti
	// + TC05 Cassa Nazionale del Notariato
	// + TC06 Cassa Nazionale Previdenza e Assistenza Ragionieri e Periti Commerciali
	// + TC07 Ente Nazionale Assistenza Agenti e Rappresentanti di Commercio (ENASARCO)
	// + TC08 Ente Nazionale Previdenza e Assistenza Consulenti del Lavoro (ENPACL)
	// + TC09 Ente Nazionale Previdenza e Assistenza Medici (ENPAM)
	// + TC10 Ente Nazionale Previdenza e Assistenza Farmacisti (ENPAF)
	// + TC11 Ente Nazionale Previdenza e Assistenza Veterinari (ENPAV)
	// + TC12 Ente Nazionale Previdenza e Assistenza Impiegati dell'Agricoltura (ENPAIA)
	// + TC13 Fondo Previdenza Impiegati Imprese di Spedizione e Agenzie Marittime
	// + TC14 Istituto Nazionale Previdenza Giornalisti Italiani (INPGI)
	// + TC15 Opera Nazionale Assistenza Orfani Sanitari Italiani (ONAOSI)
	// + TC16 Cassa Autonoma Assistenza Integrativa Giornalisti Italiani (CASAGIT)
	// + TC17 Ente Previdenza Periti Industriali e Periti Industriali Laureati (EPPI)
	// + TC18 Ente Previdenza e Assistenza Pluricategoriale (EPAP)
	// + TC19 Ente Nazionale Previdenza e Assistenza Biologi (ENPAB)
	// + TC20 Ente Nazionale Previdenza e Assistenza Professione Infermieristica (ENPAPI)
	// + TC21 Ente Nazionale Previdenza e Assistenza Psicologi (ENPAP)
	// + TC22 INPS

    	//AliquotaIVA: IVA applicata al contributo cassa
	//previdenziale. Va espressa in termini percentuali (es.: il
	//10% si esprime come 10.00 e non come 0.10),
	//altrimenti il file viene scartato con codice  e 00424.
	//Nel caso di non applicabilità, l’elemento deve essere
	//valorizzato a zero: se valorizzato a zero il sistema
	//verifica che sia presente l’elemento Natura; qualora
	//assente, il file viene scartato con codice  e 00413.



	// RiferimentoAmministrazione: eventuale riferimento
	// utile al destinatario pe automatizzare la gestione
	// amministrativa dell’operazione in fattura (capitolo di
	// spesa, conto economico ...)

	// Ritenuta: indica se il contributo cassa è soggetto a
	// ritenuta. Se soggetta (elemento valorizzato con SI) il
	// sistema controlla la presenza del blocco DatiRitenuta di
	// cui sopra: se questo blocco è assente, il file viene
	// scartato con codice  e 00415.