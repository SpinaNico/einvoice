package header

import (
	"fmt"
	"testing"

	share "github.com/SpinaNico/go-struct-invoice/share"
)

func TestEmail(t *testing.T) {

	var s email
	s = email("")
	if err := s.Validate(); err == nil {
		t.Errorf("Email not check of empity string")
	}
	s = email("1")
	if err := s.Validate(); fmt.Sprintf("%s", err) != string(share.ErrorIncluded(7, 256)) {
		t.Errorf("Error not run included error string %s", err)
	}

	c := []rune{}
	for i := 0; i < 300; i++ {
		t := []rune("a")
		c = append(c, t[0])
	}
	s = email(c)
	if err := s.Validate(); fmt.Sprintf("%s", err) != string(share.ErrorIncluded(7, 256)) {
		t.Errorf("Error not run included error string %s", err)
	}

	s = email("this is not an intentional email")
	if err := s.Validate(); err == nil {
		t.Errorf("Error not control for invalid email")
	}

}
