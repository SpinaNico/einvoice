package header

import "testing"
import "encoding/xml"
import "fmt"

func makeContattiTest(email string, telefono string) []byte {
	return []byte(fmt.Sprintf(`
	<Contatti>
		<Email>%s</Email>
		<Telefono>%s</Telefono>
	</Contatti>
	`, email, telefono))
}

func TestContatti(t *testing.T) {
	var c contatti
	s := makeContattiTest("prova@Email.com", "123456789")
	xml.Unmarshal(s, &c)
	if c.Telefono != "123456789" {
		t.Errorf("not egual %s!=123456789", c.Telefono)
	}
	if c.Email != "prova@Email.com" {
		t.Errorf("not egual %s!=\"prova@Email.com\"", c.Email)
	}
	if err := c.Validate(); err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestEmailWithoutSnail(t *testing.T) {
	s := makeContattiTest("provax", "123456789")
	var c contatti
	xml.Unmarshal(s, &c)
	if err := c.Validate(); err == nil {
		t.Errorf("contatti Validate not correct work")
	}
}
