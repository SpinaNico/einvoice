package main

import (
	"github.com/spf13/cobra"
	"github.com/spinanico/einvoice/invoicer"
)

func main() {

	cmd := cobra.Command{}

	cmd.AddCommand(commandTable())
	cmd.Execute()
}

func NewInvoicer() invoicer.Invoicer {
	return invoicer.New()
}
