package einvoice



var ErrorsMap = map[string]string{
	"00330": "l’indirizzo PEC indicato nel campo PECDestinatario  corrisponde ad una casella PEC del SdI",
	"00413": "AliquotaIVA è zero, ma non è stata definita la natura",
	"00415": "L'elemento Ritenuta è stato valorizzato con SI, ma non è presente il blocco DatiRitenuta.",
}



