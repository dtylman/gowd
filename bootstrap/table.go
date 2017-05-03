package bootstrap

import (
	"github.com/dtylman/gowd"
	"fmt"
)

type Table struct {
	*gowd.Element
	Head *gowd.Element
	Body *gowd.Element
	Rows []*gowd.Element
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
	t.Rows = make([]*gowd.Element, 0)
	return t
}

func (t *Table) AddRow() *gowd.Element {
	row := gowd.NewElement("tr")
	t.Rows = append(t.Rows, row)
	t.Body.AddElement(row)
	return row
}

func (t* Table) AddHeader(caption string) *gowd.Element{
	th:=gowd.NewElement("th")
	th.AddElement(gowd.NewText(caption))
	t.Head.AddElement(th)
	return th
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