package share

import (
	"fmt"
)

// ErrorMaxLength : Errore della lunghezza
func ErrorMaxLength(n int) string {
	return fmt.Sprintf(MaxLengthErrorString, n)
}

// ErrorIncluded : Errore per valori string compresi tra due lunghezze
func ErrorIncluded(min int, max int) string {
	return fmt.Sprintf(MaxMinErrorString, min, max)
}

// ErrorEgual : Errore per lunghezza pari a...
func ErrorEgual(n int) string {
	return fmt.Sprintf(EgualStringLengthError, n)
}

// ErrorIncorrectValue ...
func ErrorIncorrectValue(value string) string {
	return fmt.Sprintf(IncorrectValueErrorString, value)
}

// ErrorMustBeInt ErrorMustBeInt
func ErrorMustBeInt() string {
	return fmt.Sprintf(MustBeInt)
}
