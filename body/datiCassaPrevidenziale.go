package body

//DatiCassaPrevidenziale (i seguenti elementi vanno valorizzati
//nei casi in cui sia previsto il contributo cassa previdenziale)
type datiCassaPrevidenziale struct {

	//TipoCassa: cassa di previdenza della categoria
	//professionale di appartenenza.
	TipoCassa string `xml:"TipoCassa" json:"TipoCassa"`

	//AlCassa: aliquota contributiva (espressa in percentuale %) prevista per la cassa di previdenza.
	AlCassa int `xml:"AlCassa" json:"AlCassa"`

	// ImportoContributoCassa: importo del contributo relativo alla cassa
	// di previdenza della categoria professionale.
	ImportoContributoCassa string `xml:"ImportoContributoCassa" json:"ImportoContributoCassa"`

	//ImponibileCassa: importo totale del volume di affari
	//sul quale occorre applicare il contributo di cassa previdenziale.
	ImponibileCassa float32 `xml:"ImponibileCassa" json:"ImponibileCassa"`

	//AliquotaIVA: IVA applicata al contributo cassa
	//previdenziale. Va espressa in termini percentuali (es.: il
	//10% si esprime come 10.00 e non come 0.10),
	//altrimenti il file viene scartato con codice errore 00424.
	//Nel caso di non applicabilità, l’elemento deve essere
	//valorizzato a zero: se valorizzato a zero il sistema
	//verifica che sia presente l’elemento Natura; qualora
	//assente, il file viene scartato con codice errore 00413.
	AliquotaIVA float32 `xml:"AliquotaIVA" json:"AliquotaIVA"`

	// Ritenuta: indica se il contributo cassa è soggetto a
	// ritenuta. Se soggetta (elemento valorizzato con SI) il
	// sistema controlla la presenza del blocco DatiRitenuta di
	// cui sopra: se questo blocco è assente, il file viene
	// scartato con codice errore 00415.
	Ritenuta string `xml:"Ritenuta" json:"Ritenuta"`

	// Natura: codice che esprime la natura della non
	// imponibilità del contributo cassa. Deve essere presente
	// nel solo caso in cui l’elemento AliquotaIVA vale zero.
	// 40 Se è presente a fronte di un valore dell’elemento
	// AliquotaIVA diverso da zero, il file viene scartato con
	// codice errore 0414.
	Natura int `xml:"Natura" json:"Natura"`

	// RiferimentoAmministrazione: eventuale riferimento
	// utile al destinatario pe automatizzare la gestione
	// amministrativa dell’operazione in fattura (capitolo di
	// spesa, conto economico ...)
	RiferimentoAmministrazione string `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione"`
}

// Validate ...
func (f datiCassaPrevidenziale) Validate() error {
	return nil
}
