package bootstrap

import (
	"fmt"
	"github.com/dtylman/gowd"
)

//NewElement returns new bootstrap element
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

//NewContainer returns new bootstrap container.
func NewContainer(fluid bool, kids ...*gowd.Element) *gowd.Element {
	if fluid {
		return NewElement("div", "container-fluid", kids...)
	}
	return NewElement("div", "container", kids...)
}

//NewFormGroup returns new bootsrap form group
func NewFormGroup(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "form-group", elems...)
}

//NewRow return new bootstrap row
func NewRow(elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", "row", elems...)
}

const (
	ColumnLarge     = "col-lg"
	ColumnMedium    = "col-md"
	ColumnSmall     = "col-sm"
	ColumnXtraSmall = "col-xs"
)

//NewColumn returns new bootstrap column
func NewColumn(size string, span int, elems ...*gowd.Element) *gowd.Element {
	return NewElement("div", fmt.Sprintf("%s-%d", size, span), elems...)
}
