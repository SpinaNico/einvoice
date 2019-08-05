package share

import (
	"fmt"
)

// ErrorMaxLength : Errore della lunghezza
func ErrorMaxLength(n int) string {
	return fmt.Sprintf("it cannot be longer than %d characters", n)
}

// ErrorIncluded : Errore per valori string compresi tra due lunghezze
func ErrorIncluded(min int, max int) string {
	return fmt.Sprintf("the value must be at least %d characters long %d  maximum", min, max)
}

// ErrorEgual : Errore per lunghezza pari a...
func ErrorEgual(n int) string {
	return fmt.Sprintf("must be with a length of characters equal to %d", n)
}

// ErrorIncorrectValue ...
func ErrorIncorrectValue(value string) string {
	return fmt.Sprintf("The \"%s\" value is invalid", value)
}

// ErrorMustBeInt ErrorMustBeInt
func ErrorMustBeInt() string {
	return fmt.Sprintf("must be int")
}
