package main

import (
	"github.com/spf13/cobra"
	"github.com/spinanico/einvoice/cli"
)

func main() {

	cmd := cobra.Command{}

	cmd.AddCommand(cli.CommandTable())
	cmd.Execute()
}
