package bootstrap

import "github.com/dtylman/gowd"

//ListItems a list of elements
type ListItems []*gowd.Element

//List is a struct for <ul>, <ol>, <dl>
type List struct {
	*gowd.Element
	Items ListItems
}

const (
	//ListUnordered is <ul>
	ListUnordered = "ul"
	//ListOrdered is <ol>
	ListOrdered = "ol"
	//DescriptionList is <dl>
	DescriptionList = "dl"
)

//NewList creates a new list
func NewList(listType string, class string) *List {
	l := new(List)
	l.Element = NewElement(listType, class)
	l.Items = make(ListItems, 0)
	return l
}

//AddItem creates new LI, adds the elem to the li and returns the li to the caller.
func (l *List) AddItem(elem *gowd.Element) *gowd.Element {
	item := gowd.NewElement("li")
	item.AddElement(elem)
	l.AddElement(item)
	l.Items = append(l.Items, elem)
	return item
}
