package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spinanico/einvoice/invoicer"
)

func readerOfDir(dir string, format string) []string {

	var output = make([]string, 0)

	entry, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entry {
		if e.IsDir() {
			output = append(output, readerOfDir(path.Join(dir, e.Name()), format)...)
		} else {
			info, err := e.Info()
			if err != nil {
				log.Fatal(err)
			}
			if strings.HasSuffix(strings.ToLower(info.Name()), strings.ToLower(format)) {
				output = append(output, path.Join(dir, info.Name()))
			}
		}
	}

	return output
}

func CommandTable() *cobra.Command {

	c := cobra.Command{
		Aliases: []string{"list", "l"},
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				log.Fatal("add folder in argument ")
			}

			inv := invoicer.New()

			for _, inputdir := range args {

				table := tablewriter.NewWriter(os.Stdout)
				var files = readerOfDir(inputdir, ".xml")
				table.SetHeader([]string{"Data", "Num Doc", "File", "Total", "Iva"})
				fatturato := 0.0
				iva := 0.0
				for _, f := range files {

					row := make([]string, 0)
					fat, err := inv.FromFile(f)
					if err != nil {
						log.Println("error:", err)
						continue
					}
					date := fat.Date()

					row = append(row, date.Format("2006-01-02"))

					n := fat.Number()

					row = append(row, n)
					row = append(row, f)

					total := fat.Total()
					row = append(row, fmt.Sprintf("%.2f €", total))
					fatturato += total

					iva := fat.TotalTax()
					row = append(row, fmt.Sprintf("%.2f €", iva))

					table.Append(row)
				}
				table.Append([]string{"", "", "", "", ""})
				table.Append([]string{"", "", "", fmt.Sprintf("fatturato: %.2f €", fatturato), fmt.Sprintf("Iva: %.2f €", iva)})
				table.Render()
			}

		},
	}

	return &c
}
