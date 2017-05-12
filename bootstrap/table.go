package bootstrap

import (
	"github.com/dtylman/gowd"
	"fmt"
)

type TableRow struct {
	*gowd.Element
}

type Table struct {
	*gowd.Element
	Head *gowd.Element
	Body *gowd.Element
	Rows []*TableRow
}

const TableStripped = "table-striped"

func NewTable(tableType string) *Table {
	t := new(Table)
	t.Element = NewElement("table", "table")
	if tableType != "" {
		t.Element.SetClass(tableType)
	}
	t.Head = gowd.NewElement("thead")
	t.AddElement(t.Head)
	t.Body = gowd.NewElement("tbody")
	t.AddElement(t.Body)
	t.Rows = make([]*TableRow, 0)
	return t
}

func (t *Table) AddRow() *TableRow {
	row := NewTableRow()
	t.Rows = append(t.Rows, row)
	t.Body.AddElement(row.Element)
	return row
}

func (t*Table) AddHeader(caption string) *gowd.Element {
	th := gowd.NewElement("th")
	th.AddElement(gowd.NewText(caption))
	t.Head.AddElement(th)
	return th
}

func NewCell(caption string) *gowd.Element {
	td := gowd.NewElement("td")
	td.AddElement(gowd.NewText(caption))
	return td
}

func QuickTable(tableType string, data map[string]interface{}) *Table {
	t := NewTable(tableType)
	for key, value := range data {
		row := t.AddRow()
		row.AddElement(NewCell(key))
		row.AddElement(NewCell(fmt.Sprintf("%v", value)))
	}
	return t
}

func NewTableRow() *TableRow {
	tr := new(TableRow)
	tr.Element = gowd.NewElement("tr")
	return tr
}

func (tr*TableRow) AddCells(cells...string) {
	for _, cell := range cells {
		tr.AddElement(NewCell(cell))
	}
}