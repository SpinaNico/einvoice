#### Body Invoices

**IT**

**Se vuoi usare questa libreria, ti segnalo che ho creato già le stretture in typescript che trovi [QUI]().**

**Questa libreria fa parte dell'applicazione [invoices.spinanico.com](https://invoices.spinanico.com)   dove è stata implementata e creata l'interfaccia web. Puoi usare liberamente l'applicazione,  se  vuoi vedere come l'ho fatta trovi il codice [QUI]()**

**Note**: *Questo pacchetto è un sotto pacchetto del pacchetto [structinvoice]()*



**EN**

....

---

**go dependencies:**

+ ***structinvoice/share*** `internal`

---

#### structure XML and JSON 

+ [FatturaElettronicaBody](#FatturaElettronicaBody) `[]array`
  + DatiGenerali
    + DatiGeneraliDocumento `{}object`
      + TipoDocumento `string` 
      + Divisa `string`
      + Data `string` `format: YYYY-MM-DD`
      + Numero `string`
      + Causale `[]string`
      + DatiRitenuta `{}object`
        + TipoRitenuta 
        + ImportoRitenuta
        + AliquotaRitenuta
        + CausalePagamento
      + DatiBollosia `{}object`
        + BolloVirtuale
        + ImportoBollo
      + DatiCassaPrevidenziale `{}object`
        + TipoCassa
        + AlCassa
        + ImportoContributoCassa
        + ImponibileCassa
        + AliquotaIVA
        + Ritenuta
        + Natura
        + RiferimentoAmministrazione
      + ScontoMaggiorazione `{}object`
        + Tipo
        + Percentuale
        + Importo
      + ImportoTotaleDocumento
      + Arrotondamento
      + Art73
    + DatiOrdineAcquisto
    + DatiContratto
    + DatiConvenzione
    + DatiRicezione
    + DatiTrasporto
  + DatiBeniServizi
  + DatiPagamento `{}object`
    + CondizioniPagamento
    + DettaglioPagamento
      + ModalitaPagamento `string`
      + DataScadenzaPagamento `string` `format:YYYY-MM-DD`
      + ImportoPagamento `number` 
  + DatiVeicoli
    + Data `string` `format:YYYY-MM-DD`
    + TotalePercorso `string` 
  + Allegati
    + NomeAttachment
    + AlgoritmoCompressione
    + FormatoAttachment
    + DescrizioneAttachment
    + Attachment





# Description every XML Tag 

### <span id="FatturaElettronicaBody">FatturaElettronicaBody </span>

**IT**

Il secondo tipo complesso, FatturaElettronicaBody, è obbligatorio e può
essere inserito anche N volte all’interno della fattura elettronica nel caso
in cui si intenda spedire un Lotto di fatture. Contiene i seguenti
macroblocchi di dati:

+ DatiGenerali
+ DatiBeniServizi
+ DatiVeicoli
+ DatiPagamento
+ Allegati

**EN**

...

---

### DatiGenerali

**IT**

