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

	s = email(share.GenerateStringForLength(300))
	if err := s.Validate(); fmt.Sprintf("%s", err) != string(share.ErrorIncluded(7, 256)) {
		t.Errorf("Error not run included error string %s", err)
	}

	s = email("this is not an correct email")
	if err := s.Validate(); err == nil {
		t.Errorf("Error not control for invalid email")
	}

}
