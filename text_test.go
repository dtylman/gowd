package gowd

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"testing"
)

func TestNewText(t *testing.T) {
	txt := NewText("hello world")
	assert.EqualValues(t, txt.data, "hello world")
	assert.EqualValues(t, txt.nodeType, html.TextNode)
}

func TestNewStyledText(t *testing.T) {
	txt := NewStyledText("hello", SmallText)
	assert.EqualValues(t, txt.Kids[0].data, "hello")
	assert.EqualValues(t, txt.data, "small")
}
