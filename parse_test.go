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
	elem, err := ParseElement(`<br/><p/><div></div>`, em)
	assert.EqualError(t, err, `The provided html must yield only one html element, I have: [3:'br', 3:'p', 3:'div', ]`)
	assert.Nil(t, elem)
	elem, err = ParseElement("", nil)
	assert.EqualError(t, err, "The provided html must yield only one html element, I have: []")
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
