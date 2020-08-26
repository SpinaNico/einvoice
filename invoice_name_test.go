package einvoice

import "testing"

func mockF() *FatturaElettronica {
	return &FatturaElettronica{
		FatturaElettronicaHeader: FatturaElettronicaHeader{},
		FatturaElettronicaBody: []FatturaElettronicaBody{
			{},
		},
	}
}

func TestNameFile(t *testing.T) {

	result := SDIValidation("example_fake_error", mockF())
	if result == nil {
		t.Error("mi aspettavo un errore")
	}

	if !in("A0001", result.(*SDIError).GetCodes()) {
		t.Error("mi aspettavo il codice errore +0001")
	}

	if !in("A0002", result.(*SDIError).GetCodes()) {
		t.Error("mi aspettavo il codice errore +0002")
	}

	if !in("00001", result.(*SDIError).GetCodes()) {
		t.Error("mi aspettavo il codice errore 00001")
	}

	result = SDIValidation("IT01234567890_DF_FPA01.xml", mockF())
	if result != nil {
		t.Errorf("ho ricevuto errori: %v", result)
	}
}

func in(s string, ss []string) bool {
	for _, e := range ss {
		if e == s {
			return true
		}
	}
	return false
}
