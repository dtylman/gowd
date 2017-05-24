package gowd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	output := bytes.NewBuffer(make([]byte, 0))
	Order = 0
	elem := NewElement("div")
	render(elem, output)
	expected := `<div id="_div1"></div>`
	assert.Equal(t, expected+"\n", output.String())
}
