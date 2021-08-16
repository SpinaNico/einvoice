
still under construction, there is nothing to see here


#### Body Invoices

> Allegati > Attachment > Controllare se è base64 

> **datiCassaPrevidenziale > AliquotaIva**
> Viene detto di settarlo con zero ma si intende: "0" intero o "0.0"

> dati Cassa previdenziale > ritenuta SI ma solo se ci sono i dati

> datiRitenuta > CausalePagamento Può avere i valori CU, ZO, M2 ma cosa indicano?
---

#### structure XML 

FatturaElettronicaBody

+ DatiGenerali
  + DatiGeneraliDocumento 
    + TipoDocumento
    + Divisa 
    + Data 
    + Numero 
    + Causale 
    + DatiRitenuta 
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

 