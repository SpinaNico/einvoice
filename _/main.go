package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	invoice "github.com/SpinaNico/einvoice"
)

func main() {

	var f invoice.FatturaElettronica
	if len(os.Args) == 0 {
		panic("name file of invoice (.xml) required")
	}
	name := os.Args[1]
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(fmt.Errorf(" ReadFIle: %s", err))
	}
	xml.Unmarshal(data, &f)

	if err := f.Validate(); err != nil {
		log.Fatal(fmt.Errorf("ValidateInvoice: %s", err))
	}

	jsonValue, _ := json.MarshalIndent(&f, "", "  ")
	ioutil.WriteFile("invoice.json", jsonValue, 0777)
}
