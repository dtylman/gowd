package bootstrap

import (
	"github.com/dtylman/pictures/webkit"
	"fmt"
)

func NewElement(tag, class string) *webkit.Element {
	elem := webkit.NewElement(tag)
	elem.SetAttribute("class", class)
	return elem
}

func NewContainer(fluid bool) *webkit.Element {
	if fluid {
		return NewElement("div", "container-fluid")
	}
	return NewElement("div", "container")
}

func NewRow() *webkit.Element {
	return NewElement("div", "row")
}

const (
	ColumnLarge = "col-lg-"
	ColumnMedium = "col-md-"
	ColumnSmall = "col-sm-"
	ColumnXtraSmall = "col-xs"
)

func NewColumn(size string, count int) *webkit.Element {
	return NewElement("div", fmt.Sprintf("%s-%d", size, count))
}