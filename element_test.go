package gowd

import (
	"github.com/stretchr/testify/assert"
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
