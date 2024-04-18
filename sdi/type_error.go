package sdi

type ErrorLevel string

const (
	Warning = "WARN"
	Error   = "ERROR"
)

type SDIError struct {
	Level       ErrorLevel
	Description string
	Stack       []string
	SDICode     ErrorSDICode
}
