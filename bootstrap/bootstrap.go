package bootstrap

import (
	"fmt"
	"github.com/dtylman/gowd"
)

func NewElement(tag, class string, kids ...*gowd.Element) *gowd.Element {
	elem := gowd.NewElement(tag)
	if class != "" {
		elem.SetAttribute("class", class)
	}
	for _, kid := range kids {
		elem.AddElement(kid)
	}
	return elem
}

func NewContainer(fluid bool, kids ...*gowd.Element) *gowd.Element {
	if fluid {
		return NewElement("div", "container-fluid", kids...)
	}
	return NewElement("div", "container", kids...)
}

func NewFormGroup(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "form-group", elems...)
}

func NewRow(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "row", elems...)
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