package bootstrap

import (
	"github.com/dtylman/gowd"
	"fmt"
)

type Table struct {
	*gowd.Element
	head *gowd.Element
	body *gowd.Element
	rows []*gowd.Element
}

const TableStripped = "table-striped"

func NewTable(tableType string) *Table {
	t := new(Table)
	t.Element = NewElement("table", "table")
	if tableType != "" {
		t.Element.SetClass(tableType)
	}
	t.head = gowd.NewElement("thead")
	t.AddElement(t.head)
	t.body = gowd.NewElement("tbody")
	t.AddElement(t.body)
	t.rows = make([]*gowd.Element, 0)
	return t
}

func (t *Table) AddRow() *gowd.Element {
	row := gowd.NewElement("tr")
	t.rows = append(t.rows, row)
	t.body.AddElement(row)
	return row
}

func (t*Table) NewCell(caption string) *gowd.Element {
	td := gowd.NewElement("td")
	td.AddElement(gowd.NewText(caption))
	return td
}

func QuickTable(tableType string, data map[string]interface{}) *Table {
	t := NewTable(tableType)
	for key, value := range data {
		row := t.AddRow()
		row.AddElement(t.NewCell(key))
		row.AddElement(t.NewCell(fmt.Sprintf("%v", value)))
	}
	return t
}