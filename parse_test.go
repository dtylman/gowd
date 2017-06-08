package gowd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testOuput(t *testing.T, elem *Element, expected string) {
	output := bytes.NewBuffer(make([]byte, 0))
	Output = output
	Order = 0
	elem.Render()
	assert.Equal(t, expected+"\n", output.String())
}

func TestParseElement(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div id='div'><b id='text'>text</b><button id="btn" value="lala"/></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, elem, em["div"])
	assert.EqualValues(t, "lala", em["btn"].GetValue())
}

func TestParseElement2(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div id="myDiv"><h1>My First Heading</h1><p>My first paragraph.</p></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	testOuput(t, elem, "<div id=\"myDiv\"><h1>My First Heading</h1><p>My first paragraph.</p></div>")
	div := elem.Find("myDiv")
	assert.NotNil(t, div)
	assert.EqualValues(t, div.GetID(), em["myDiv"].GetID())
	p, err := div.AddHTML(`<p id="myP">another paragraph</p>`, em)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.EqualValues(t, em["myP"].data, p[0].data)

}
