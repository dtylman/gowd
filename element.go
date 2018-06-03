package gowd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	//Order counts the number of elements rendered (for generating auto-ids)
	Order int
	//Output output render target (configurable for unit-tests)
	Output io.Writer = os.Stdout
)

//Element represents a DOM element and its state.
type Element struct {
	//Parent the parent element
	Parent *Element
	//eventHandlers holds event handlers for events
	eventHandlers map[string]EventHandler
	//Kids child elements
	Kids []*Element
	//Attributes element attributes...
	Attributes []html.Attribute
	//Object arbitrary user object that can be associated with element.
	Object interface{}
	//nodeType the element node type
	nodeType html.NodeType
	//data html tag or text
	data string
	//Hidden if true the element will not be rendered
	Hidden bool
	//renderHash after rendered get a hash to identify if should be rendered again.
	renderHash []byte
}

//NewElement creates a new HTML element
func NewElement(tag string) *Element {
	elem := &Element{
		data:          tag,
		Attributes:    make([]html.Attribute, 0),
		nodeType:      html.ElementNode,
		Kids:          make([]*Element, 0),
		eventHandlers: make(map[string]EventHandler),
	}
	Order++
	elem.SetID(fmt.Sprintf("_%s%d", elem.data, Order))
	return elem
}

//SetAttributes sets attributes from an event element.
func (e *Element) SetAttributes(event *EventElement) {
	e.Attributes = make([]html.Attribute, 0)
	for key, value := range event.Properties {
		e.Attributes = append(e.Attributes, html.Attribute{Key: key, Val: value})
	}
}

//AddElement adds a child element
func (e *Element) AddElement(elem *Element) *Element {
	if elem == nil {
		panic("Cannot add nil element")
	}
	elem.Parent = e
	e.Kids = append(e.Kids, elem)
	return elem
}

//AddHTML parses the provided element and adds it to the current element. Returns a list of root elements from `html`.
//If em is not nil, for each HTML tag that has the `id` attribute set the corresponding element will be stored in the
//given ElementMap.
func (e *Element) AddHTML(innerHTML string, em ElementsMap) ([]*Element, error) {
	elems, err := ParseElements(strings.NewReader(innerHTML), em)
	if err != nil {
		return nil, err
	}
	for _, elem := range elems {
		e.Kids = append(e.Kids, elem)
	}
	return elems, nil
}

//RemoveElements remove all kids.
func (e *Element) RemoveElements() {
	e.Kids = make([]*Element, 0)
}

//RemoveElement remove a specific kid
func (e *Element) RemoveElement(elem *Element) {
	before := e.Kids
	e.RemoveElements()
	for _, kid := range before {
		if kid.GetID() != elem.GetID() {
			e.AddElement(kid)
		}
	}
}

//SetText Sets the element to hold ONLY the provided text
func (e *Element) SetText(text string) {
	if e.nodeType == html.TextNode {
		e.data = text
	} else {
		e.RemoveElements()
		e.AddElement(NewText(text))
	}
}

//SetClass adds the given class name to the class attribute.
func (e *Element) SetClass(class string) {
	prev, exists := e.GetAttribute("class")
	if exists {
		if strings.Contains(prev, class) {
			return
		}
	}
	e.SetAttribute("class", prev+" "+class)
}

//UnsetClass removes the given class name from the class attribute
func (e *Element) UnsetClass(class string) {
	prev, exists := e.GetAttribute("class")
	if !exists {
		return
	}
	e.SetAttribute("class", strings.Replace(prev, class, "", -1))
}

//RemoveAttribute removes the provided attribute by name
func (e *Element) RemoveAttribute(key string) {
	attributes := e.Attributes
	e.Attributes = make([]html.Attribute, 0)
	for _, attrib := range attributes {
		if attrib.Key != key {
			e.Attributes = append(e.Attributes, attrib)
		}
	}
}

//SetAttribute adds or set the attribute
func (e *Element) SetAttribute(key, val string) {
	for i := range e.Attributes {
		if e.Attributes[i].Key == key {
			e.Attributes[i].Val = val
			return
		}
	}
	e.Attributes = append(e.Attributes, html.Attribute{Key: key, Val: val})
}

//GetAttribute returns value for attribute
func (e *Element) GetAttribute(key string) (string, bool) {
	if e.Attributes == nil {
		return "", false
	}
	for _, a := range e.Attributes {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

//GetValue returns the value of the `value` attribute
func (e *Element) GetValue() string {
	val, _ := e.GetAttribute("value")
	return val
}

//GetID returns the value of the `id` attribute
func (e *Element) GetID() string {
	val, _ := e.GetAttribute("id")
	return val
}

//SetID sets the `id` attribute
func (e *Element) SetID(id string) {
	e.SetAttribute("id", id)
}

//Disable  sets the `disabled` attribute
func (e *Element) Disable() {
	e.SetAttribute("disabled", "true")
}

//Enable unsets the `disabled` attribute
func (e *Element) Enable() {
	e.RemoveAttribute("disabled")
}

//Hide if set, will not render the element.
func (e *Element) Hide() {
	e.Hidden = true
}

//Show if set, will render the element.
func (e *Element) Show() {
	e.Hidden = false
}

//Find  returns the kid, or offspring with a specific `id` attribute value.
func (e *Element) Find(id string) *Element {
	if e.GetID() == id {
		return e
	}
	for i := range e.Kids {
		elem := e.Kids[i].Find(id)
		if elem != nil {
			return elem
		}
	}
	return nil
}

//OnKeyPressEvent register handler as an OnKeyPressed event.
func (e *Element) OnKeyPressEvent(event string, keyCode int, handler EventHandler) {
	e.SetAttribute(event, fmt.Sprintf(`fire_keypressed_event(event,%d,'%s',this);`, keyCode, event))
	e.eventHandlers[event] = handler
}

//OnEvent register an DOM element event.
func (e *Element) OnEvent(event string, handler EventHandler) {
	e.SetAttribute(event, fmt.Sprintf(`fire_event('%s',this);`, event))
	e.eventHandlers[event] = handler
}

func (e *Element) toNode() *html.Node {
	node := &html.Node{Attr: e.Attributes, Data: e.data, Type: e.nodeType}
	if e.Hidden {
		node.Type = html.CommentNode // to a void returning null which may crash the caller
		return node
	}
	for i := range e.Kids {
		node.AppendChild(e.Kids[i].toNode())
	}
	return node
}

//Render renders the element.
func (e *Element) Render() error {
	return render(e, Output)
}

//ProcessEvent fires the event provided
func (e *Element) ProcessEvent(event *Event) {
	for _, input := range event.Inputs {
		elementID := input.GetID()
		if elementID != "" {
			e.updateState(elementID, &input)
		}
	}
	e.fireEvent(event.Name, event.Sender.GetID(), &event.Sender)
}

func (e *Element) updateState(elementID string, input *EventElement) {
	for i := range e.Kids {
		e.Kids[i].updateState(elementID, input)
	}
	if e.GetID() == elementID {
		e.SetAttributes(input)
	}
	//special cases:
	if e.data == atom.Select.String() {
		for i := range e.Kids {
			if e.Kids[i].GetValue() == e.GetValue() {
				e.Kids[i].SetAttribute("selected", "true")
			} else {
				e.Kids[i].RemoveAttribute("selected")
			}
		}
	}
}

func (e *Element) fireEvent(eventName string, senderID string, sender *EventElement) {
	if senderID == "" {
		return
	}
	kids := e.Kids //events may change the kids container
	for i := range kids {
		kids[i].fireEvent(eventName, senderID, sender)
	}
	if e.GetID() == senderID {
		handler, exists := e.eventHandlers[eventName]
		if exists {
			handler(e, sender)
		}
	}
}
