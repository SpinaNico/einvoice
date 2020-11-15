package einvoice

import (
	"fmt"
	"strconv"
)

/*
Verifica partita IVA
Versione: 1.0
Data: 1/5/2017
Autore: Squeeze69
Licenza: LGPL
Porting basato sulle informazioni pubblicate da Umberto Salsi su Icosaedro:
sito web: http://www.icosaedro.it/cf-pi/index.html

package: https://github.com/squeeze69/partitaiva
con go: go get github.com/squeeze69/partitaiva
*/

//ItPartitaIva controlla se è una partita IVA valida nel formato
//Nota: se piva è vuota, viene considerata valida, per questo caso, il controllo dovrebbe essere altrove
//Ingresso: piva: stringa
//Uscita: bool:true (a posto)/false (problemi) e *PIVAError (nil (a posto)/puntatore all'errore (problemi)

func CheckPIVA(piva string) (bool, error) {

	if len(piva) != 11 {
		return false, fmt.Errorf(ErrorLen)
	}

	var primo, secondo, i int
	for i = 0; i <= 9; i += 2 {
		v, err := strconv.Atoi(string(piva[i]))
		if err != nil {
			return false, fmt.Errorf("Caratteri non validi")
		}
		primo += v
	}
	for i = 1; i <= 9; i += 2 {
		v, err := strconv.Atoi(string(piva[i]))
		if err != nil {
			return false, fmt.Errorf(ErrorCharacterWrong)
		}
		secondo = 2 * v
		if secondo > 9 {
			secondo = secondo - 9
		}
		primo += secondo
	}
	if v, _ := strconv.Atoi(string(piva[10])); (10-primo%10)%10 == v {
		return true, nil
	}

	return false, fmt.Errorf(ErrorCharacterController)

}
