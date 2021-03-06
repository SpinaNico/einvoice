package einvoice

import (
	"fmt"
	"strings"
)

/*
Verifica codice fiscale
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/codicefiscale
con go: go get github.com/squeeze69/codicefiscale
*/

//CodiceFiscale  controlla il codice fiscale, restituisce doppio valore, true/false e messaggio, ove opportuno
//se cfin è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
//Ingresso: cfin: stringa,non importa maiuscolo o minuscolo
//Uscita: bool:true (a posto)/false (problemi) e *CFError (nil (a posto)/puntatore all'errore (problemi)
func CheckCodiceFiscale(cfin string) (bool, error) {
	if len(cfin) == 0 {
		return true, nil
	}
	if len(cfin) != 16 {
		return false, fmt.Errorf(ErrorLen)
	}
	//decodifica carattere di controllo cf
	tcf := map[string]int{
		"0": 1, "1": 0, "2": 5, "3": 7, "4": 9, "5": 13, "6": 15, "7": 17, "8": 19,
		"9": 21, "A": 1, "B": 0, "C": 5, "D": 7, "E": 9, "F": 13, "G": 15, "H": 17,
		"I": 19, "J": 21, "K": 2, "L": 4, "M": 18, "N": 20, "O": 11, "P": 3, "Q": 6, "R": 8,
		"S": 12, "T": 14, "U": 16, "V": 10, "W": 22, "X": 25, "Y": 24, "Z": 23,
	}
	//map per simulare "ord" di altri linguaggi - più semplice di int(rune) - int('A') oppure int('0')
	ordv := map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5, "G": 6, "H": 7, "I": 8, "J": 9,
		"K": 10, "L": 11, "M": 12, "N": 13, "O": 14, "P": 15, "Q": 16, "R": 17, "S": 18,
		"T": 19, "U": 20, "V": 21, "W": 22, "X": 23, "Y": 24, "Z": 25,
	}
	cfin = strings.ToUpper(cfin)
	//verifica per simboli inattesi
	for _, v := range cfin {
		if _, ok := ordv[string(v)]; !ok {
			return false, fmt.Errorf(ErrorCharacterWrong)
		}
	}
	s := tcf[string(cfin[14])]
	for i := 0; i <= 13; i += 2 {
		s += ordv[string(cfin[i+1])]
		s += tcf[string(cfin[i])]
	}
	if s%26 != ordv[string(cfin[15])] {
		return false, fmt.Errorf(ErrorCharacterController)
	}
	return true, nil
}
