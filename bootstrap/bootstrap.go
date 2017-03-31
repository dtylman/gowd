package bootstrap

import (
	"fmt"
	"github.com/dtylman/gowd"
)

func NewElement(tag, class string) *gowd.Element {
	elem := gowd.NewElement(tag)
	if class != "" {
		elem.SetAttribute("class", class)
	}
	return elem
}

func NewContainer(fluid bool) *gowd.Element {
	if fluid {
		return NewElement("div", "container-fluid")
	}
	return NewElement("div", "container")
}

func NewRow() *gowd.Element {
	return NewElement("div", "row")
}

const (
	ColumnLarge = "col-lg"
	ColumnMedium = "col-md"
	ColumnSmall = "col-sm"
	ColumnXtraSmall = "col-xs"
)

func NewColumn(size string, span int) *gowd.Element {
	return NewElement("div", fmt.Sprintf("%s-%d", size, span))
}