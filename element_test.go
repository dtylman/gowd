package gowd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestElement_SetAttributes(t *testing.T) {
	elem := NewElement("div")
	event := EventElement{}
	event.Properties = make(map[string]string)
	event.Properties["value"] = "123"
	event.Properties["id"] = "myID"
	elem.SetAttributes(&event)
	assert.EqualValues(t, event.GetID(), "myID")
	assert.EqualValues(t, event.GetValue(), "123")
	assert.EqualValues(t, event.GetID(), elem.GetID())
	assert.EqualValues(t, event.GetValue(), elem.GetValue())
}

func TestElement_Enable(t *testing.T) {
	elem := NewElement("div")
	elem.Disable()
	val, found := elem.GetAttribute("disabled")
	assert.EqualValues(t, val, "true")
	assert.True(t, found)
	elem.Enable()
	val, found = elem.GetAttribute("disabled")
	assert.EqualValues(t, val, "")
	assert.False(t, found)
}

func TestElement_SetClass(t *testing.T) {
	elem := NewElement("div")
	elem.SetClass("well sunken")
	class, _ := elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well sunken")
	elem.SetClass("upper")
	class, _ = elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well sunken upper")
	elem.UnsetClass("sunken")
	class, _ = elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well  upper")
}

func TestElement_Hide(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div><p id="text">text</p></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	p := em["text"]
	assert.False(t, p.Hidden)
	testOuput(t, elem, "<div><p id=\"text\">text</p></div>")
	p.Hide()
	assert.True(t, p.Hidden)
	testOuput(t, elem, "<div><!--p--></div>")
	p.SetText("Show me!!")
	p.Show()
	testOuput(t, elem, "<div><p id=\"text\">Show me!!</p></div>")
	elem.RemoveElement(p)
	testOuput(t, elem, "<div></div>")
}
