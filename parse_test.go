package gowd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseElement(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div id='div'><b id='text'>text</b><button id="btn" value="lala"/></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, elem, em["div"])
	assert.EqualValues(t, "lala", em["btn"].GetValue())
}
