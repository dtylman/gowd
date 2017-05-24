package bootstrap

import (
	"fmt"
	"github.com/dtylman/gowd"
)

//TableRow represents a <tr>
type TableRow struct {
	*gowd.Element
}

//Table represents <table>
type Table struct {
	*gowd.Element
	Head *gowd.Element
	Body *gowd.Element
	Rows []*TableRow
}

//TableStripped table is stripped
const TableStripped = "table-striped"

//NewTable creates a new table with type
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

//AddRow adds a row
func (t *Table) AddRow() *TableRow {
	row := NewTableRow()
	t.Rows = append(t.Rows, row)
	t.Body.AddElement(row.Element)
	return row
}

//AddHeader adds an header row
func (t *Table) AddHeader(caption string) *gowd.Element {
	th := gowd.NewElement("th")
	th.AddElement(gowd.NewText(caption))
	t.Head.AddElement(th)
	return th
}

//NewCell creates and adds new cell
func NewCell(caption string) *gowd.Element {
	td := gowd.NewElement("td")
	td.AddElement(gowd.NewText(caption))
	return td
}

//QuickTable creates a table from a given map with key-value pairs
func QuickTable(tableType string, data map[string]interface{}) *Table {
	t := NewTable(tableType)
	for key, value := range data {
		row := t.AddRow()
		row.AddElement(NewCell(key))
		row.AddElement(NewCell(fmt.Sprintf("%v", value)))
	}
	return t
}

//NewTableRow creates a new table row
func NewTableRow() *TableRow {
	tr := new(TableRow)
	tr.Element = gowd.NewElement("tr")
	return tr
}

//AddCells adds cells to a table row
func (tr *TableRow) AddCells(cells ...string) {
	for _, cell := range cells {
		tr.AddElement(NewCell(cell))
	}
}
