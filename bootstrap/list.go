package bootstrap

import "github.com/dtylman/gowd"

type ListItems []*gowd.Element
type List struct {
	*gowd.Element
	Items ListItems
}

const (
	ListUnordered = "ul"
	ListOrdered = "ol"
	DescriptionList = "dl"
)

func NewList(listType string, class string) *List {
	l := new(List)
	l.Element = NewElement(listType, class)
	l.Items = make(ListItems, 0)
	return l
}

//AddItem creates new LI, adds the elem to the li and returns the li to the caller.
func (l*List) AddItem(elem*gowd.Element) *gowd.Element {
	item := gowd.NewElement("li")
	item.AddElement(elem)
	l.AddElement(item)
	l.Items = append(l.Items, elem)
	return item
}
