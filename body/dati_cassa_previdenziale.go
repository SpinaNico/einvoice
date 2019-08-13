package body

import (
	"fmt"
	"github.com/SpinaNico/go-struct-invoice/share"
	"strconv"
	"strings"
)

//DatiCassaPrevidenziale (i seguenti elementi vanno valorizzati
//nei casi in cui sia previsto il contributo cassa previdenziale)
type datiCassaPrevidenziale struct {

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
	TipoCassa string `xml:"TipoCassa" json:"TipoCassa"`

	//AlCassa: aliquota contributiva (espressa in percentuale %) prevista per la cassa di previdenza.
	AlCassa decimale2 `xml:"AlCassa" json:"AlCassa"`

	// ImportoContributoCassa: importo del contributo relativo alla cassa
	// di previdenza della categoria professionale.
	ImportoContributoCassa decimale2 `xml:"ImportoContributoCassa" json:"ImportoContributoCassa"`

	//ImponibileCassa: importo totale del volume di affari
	//sul quale occorre applicare il contributo di cassa previdenziale.
	ImponibileCassa decimale2 `xml:"ImponibileCassa" json:"ImponibileCassa"`

	//AliquotaIVA: IVA applicata al contributo cassa
	//previdenziale. Va espressa in termini percentuali (es.: il
	//10% si esprime come 10.00 e non come 0.10),
	//altrimenti il file viene scartato con codice errore 00424.
	//Nel caso di non applicabilità, l’elemento deve essere
	//valorizzato a zero: se valorizzato a zero il sistema
	//verifica che sia presente l’elemento Natura; qualora
	//assente, il file viene scartato con codice errore 00413.
	AliquotaIVA decimale2 `xml:"AliquotaIVA" json:"AliquotaIVA"`

	// Ritenuta: indica se il contributo cassa è soggetto a
	// ritenuta. Se soggetta (elemento valorizzato con SI) il
	// sistema controlla la presenza del blocco DatiRitenuta di
	// cui sopra: se questo blocco è assente, il file viene
	// scartato con codice errore 00415.
	Ritenuta string `xml:"Ritenuta" json:"Ritenuta"`

	Natura natura `xml:"Natura" json:"Natura"`

	// RiferimentoAmministrazione: eventuale riferimento
	// utile al destinatario pe automatizzare la gestione
	// amministrativa dell’operazione in fattura (capitolo di
	// spesa, conto economico ...)
	RiferimentoAmministrazione string `xml:"RiferimentoAmministrazione" json:"RiferimentoAmministrazione"`
}

// Validate ...
func (f datiCassaPrevidenziale) Validate() error {
	var err error
	if len(f.TipoCassa) != 4 {
		return fmt.Errorf("TipoCassa %s", share.ErrorEgual(4))
	}
	f.TipoCassa = strings.ToUpper(f.TipoCassa)
	values := make(map[string]string)
	for i := 1; i < 23; i++ {
		values[fmt.Sprintf("TC%d", i)] = "true"
	}
	if _, ok := values[f.TipoCassa]; ok == false {
		return fmt.Errorf("TipoCassa %s", share.ErrorIncorrectValue(f.TipoCassa))
	}

	if err = f.AlCassa.Validate(); err != nil {
		return fmt.Errorf("AlCassa %s", err)
	}

	if len(string(f.AlCassa)) > 6 && len(string(f.AlCassa)) < 4 {
		return fmt.Errorf("AlCassa %s", share.ErrorIncluded(4, 6))
	}

	if err = f.ImponibileCassa.Validate(); err != nil {
		return fmt.Errorf("ImponibileCassa %s", err)
	}
	if len(string(f.ImponibileCassa)) > 15 {
		return fmt.Errorf("ImponibileCassa %s", share.ErrorMaxLength(15))
	}

	if err = f.ImportoContributoCassa.Validate(); err != nil {
		return fmt.Errorf("ImportoContributoCassa %s", err)
	}
	if len(string(f.ImportoContributoCassa)) > 15 {
		return fmt.Errorf("ImportoContributoCassa %s", share.ErrorMaxLength(15))
	}

	if err = f.AliquotaIVA.Validate(); err != nil {
		return fmt.Errorf("AliquotaIVA %s", err)
	}

	if len(string(f.AliquotaIVA)) > 6 && len(string(f.AliquotaIVA)) < 4 {
		return fmt.Errorf("AliquotaIVA %s", share.ErrorIncluded(4, 6))
	}

	if val, _ := strconv.ParseFloat(string(f.AliquotaIVA), 32); val == 0.0 {
		// Non sono sicuro di questo punto
		f.AliquotaIVA = decimale2("0")
		if f.Natura == "" {
			return fmt.Errorf("Natura: is empity but AliquotaIVA is 0.0")
		}
	}

	if f.Natura != "" {
		if err = f.Natura.Validate(); err != nil {
			return fmt.Errorf("Natura %s", err)
		}
	}

	if len(f.RiferimentoAmministrazione) > 20 {
		return fmt.Errorf("RiferimentoAmministrazione %s", share.ErrorMaxLength(20))
	}

	if f.Ritenuta != "" && f.Ritenuta != "SI" {
		return fmt.Errorf("Ritenuta %s", share.ErrorIncorrectValue(f.Ritenuta))
	}

	return nil
}
