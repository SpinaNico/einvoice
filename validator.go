package invoice

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

func regimeFiscaleValidator(rf validator.FieldLevel) bool {
	RF := rf.Field().String()

	regimeFiscale := make(map[string]string)
	for i := 1; i < 20; i++ {
		regimeFiscale[fmt.Sprintf("RF%d", i)] = "true"
	}
	_, exists := regimeFiscale[string(RF)]
	//println(exists, RF)
	return exists == true
}

func trasmissioneValidate(d validator.StructLevel) {

}
