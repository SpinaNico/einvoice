package sdi

type CondizionePagamento string

const TP01 CondizionePagamento = "TP01"
const TP02 CondizionePagamento = "TP02"
const TP03 CondizionePagamento = "TP03"

var CondizioniPagamento = map[CondizionePagamento]string{
	TP01: "pagamento a rate",
	TP02: "pagamento completo",
	TP03: "anticipo",
}

func IsValidCondizionePagamento(condizione CondizionePagamento) bool {
	_, ok := CondizioniPagamento[condizione]
	return ok
}
