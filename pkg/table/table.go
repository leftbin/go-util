package table

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func PrintTable(header table.Row, rows []table.Row) {
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if header != nil {
		t.AppendHeader(header)
	}
	for _, r := range rows {
		t.AppendRow(r)
		t.AppendSeparator()
	}
	t.Render()
	println("")
}
