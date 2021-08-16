
## Cose mie
controlla il Blocca Anagrafica, e il gioco sulle denominazioni se funziona correttamente.


## Domande Acgenzia delle entrate

FormatoTrasmisione 
Nel documento Allegato delle specifiche indicate come unico valore per il blocco 
"FormatoTrasmisione" il valore: `FPR12` questo nella documentazione.

ma negli esempi usate quest'altro valode: `FPA12`
si possono usare entrambi? credo che uno rappresenti il privato, l'altro il pubblico
ma gradirei una conferma di questa mia supposizione.

---
CodiceDestinatario

Il file : `IT01234567890_FPA01.xml`
ha questo nel CodiceDestinatario
<CodiceDestinatario>AAAAAAA</CodiceDestinatario>

il varlore non è valido ho corretto il mio problema aggiugendo un "A" ma è forviante
perché è un elemento di 7 caratteri alfanumerici e non di soli 6.
---


FatturaElettronicaBody.DatiGenerali.DatiGeneraliDocumento.DatiCassaPrevidenziale.Ritenuta == "SI"
(FatturaElettronicaBody.DatiGenerali.DatiGeneraliDocumento.DatiRitenuta) si deve validare altrimenti no