package header

type contatti struct {
	Telefono string `xml:"Telefono" json:"Telefono" validate:"omitempty,min=5,max=12"`
	//todo: Validate Fax
	Fax   string `xml:"Fax" json:"Fax"`
	Email string `xml:"Email" json:"Email" validate:"omitempty,email,max=255"`
}

// func (c contatti) Validate() error {
// 	var err error

// 	if err = c.Email.Validate(); err != nil {
// 		return fmt.Errorf("Email %s", err)
// 	}

// 	if err = c.Telefono.Validate(); err != nil {
// 		return fmt.Errorf("Telefono %s", err)
// 	}
// 	return nil
// }
