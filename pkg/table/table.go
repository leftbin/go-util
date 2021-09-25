package table

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTable(header table.Row, rows []table.Row) {
	println("")
	t := getDefaultTableWriter(header, rows)
	t.Render()
	println("")
}

func PrintTableWithSort(header table.Row, rows []table.Row, sortRules ...table.SortBy) {
	println("")
	t := getDefaultTableWriter(header, rows)
	if sortRules != nil && len(sortRules) > 0 {
		t.SortBy(sortRules)
	}
	t.Render()
	println("")
}

func PrintTableWithColConfig(header table.Row, rows []table.Row, colConfigs ...table.ColumnConfig) {
	println("")
	t := getDefaultTableWriter(header, rows)
	if colConfigs != nil && len(colConfigs) > 0 {
		t.SetColumnConfigs(colConfigs)
	}
	t.Render()
	println("")
}

func PrintTableWithSortAndColConfig(header table.Row, rows []table.Row, sortRules []table.SortBy, colConfigs []table.ColumnConfig) {
	println("")
	t := getDefaultTableWriter(header, rows)
	if sortRules != nil && len(sortRules) > 0 {
		t.SortBy(sortRules)
	}
	if colConfigs != nil && len(colConfigs) > 0 {
		t.SetColumnConfigs(colConfigs)
	}
	t.Render()
	println("")
}

func getDefaultTableWriter(header table.Row, rows []table.Row) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if header != nil {
		t.AppendHeader(header)
	}
	for _, r := range rows {
		t.AppendRow(r)
		t.AppendSeparator()
	}
	return t
}
