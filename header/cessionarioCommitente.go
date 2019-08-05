package header

import (
	"fmt"
)

//TODO:
//Se io NON sono un residente, devo
//valorizzare i blocchi: StabileOrganizzazione, RappresentanteFiscale

type cessionarioCommittente struct {
	DatiAnagrafici datiAnagrafici `xml:"DatiAnagrafici" json:"DatiAnagrafici"`
	Sede           sede           `xml:"Sede" json:"Sede"`
	// Da valorizzare solo se il cessionario non è un residente
	// quindi per l'operazione deve obbligatoriamente indicare uno stabile
	// in questa struttura deve indicare lo stabile (che risiede nel territorio italiano)
	StabileOrganizzazione sede                  `xml:"StabileOrganizzazione" json:"StabileOrganizzazione"`
	RappresentanteFiscale rappresentanteFiscale `xml:"RappresentanteFiscale" json:"RappresentanteFiscale"`
}

func (c cessionarioCommittente) Validate() error {
	var err error
	err = c.Sede.Validate()
	if err != nil {
		return fmt.Errorf("CessionarioCommittente %s", err)
	}
	if c.Sede.Nazione != "IT" {
		err = c.StabileOrganizzazione.Validate()
		if err != nil {
			return fmt.Errorf("CessionarioCommittente \"StabileOrganizzazione\" is type %s", err)
		}
		err = c.RappresentanteFiscale.Validate()
		if err != nil {
			return fmt.Errorf("CessionarioCommittente %s", err)
		}
	}
	// todo: Se è italiana la presenza eventualmente di
	// RappresentanteFiscale | StabileOrganizzazione
	// produce un errore ?
	return nil
}
