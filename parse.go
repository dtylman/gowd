package gowd

import (
	"io"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"errors"
)

//ParseElements parse an html fragment and return a list of elements
func ParseElements(r io.Reader) ([]*Element, error) {
	nodes, err := html.ParseFragment(r, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		return nil, err
	}
	elems := make([]*Element, 0)
	for _, node := range nodes {
		elems = append(elems, NewElementFromNode(node))
	}
	return elems, nil
}

//NewElementFromNode creates an element from existing node
func NewElementFromNode(node*html.Node) *Element {
	elem := &Element{
		data:          strings.Trim(node.Data,"\n\r\t"),
		Attributes:    node.Attr,
		nodeType:      node.Type,
		Kids:          make([]*Element, 0),
		eventHandlers: make(map[string]EventHandler),
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		elem.AddElement(NewElementFromNode(c));
	}
	return elem
}

func ParseElement(innerHtml string) (*Element, error) {
	elems, err := ParseElements(strings.NewReader(innerHtml))
	if err != nil {
		return nil, err
	}
	if len(elems) != 1 {
		return nil, errors.New("The provided html must yield only one html element")
	}
	return elems[0], nil
}