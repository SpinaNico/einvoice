package share

import (
	"fmt"
)

// ErrorLen : Errore della lunghezza
func ErrorLen(n int) string {
	return fmt.Sprintf("it cannot be longer than %d characters", n)
}

// ErrorIncluded : Errore per valori string compresi tra due lunghezze
func ErrorIncluded(min int, max int) string {
	return fmt.Sprintf("the value must be at least %d characters long %d  maximum", min, max)
}
