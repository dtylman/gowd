package gowd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	output := bytes.NewBuffer(make([]byte, 0))
	Order = 0
	elem := NewElement("div")
	render(elem, output)
	expected := `<div id="_div1"></div>`
	assert.Equal(t, expected+"\n", output.String())
}

func Test_ProcessEvent(t *testing.T) {
	elem := NewElement("input")
	elem.SetAttribute("value", "text")
	elem.SetID("_div12")
	elem.OnEvent("onclick", func(sender *Element, event *EventElement) {
		assert.EqualValues(t, sender.GetID(), "_div12")
		assert.EqualValues(t, event.GetID(), "_div12")
		assert.EqualValues(t, sender.GetValue(), "text")
		assert.EqualValues(t, event.GetValue(), "")
	})
	jsEvent := `{"name":"onclick","sender":{"properties":{"id":"_div12"}},"inputs":[]}`
	err := processEvents(elem, bytes.NewBufferString(jsEvent))
	assert.NoError(t, err)
}

func Test_Run(t *testing.T) {
	elem := NewElement("div")
	waitTime := time.Millisecond * 750
	startTime := time.Now()
	go func() {
		err := Run(elem)
		assert.EqualError(t, err, "EOF")
	}()
	time.Sleep(waitTime)
	totalTime := time.Now().Sub(startTime)
	assert.True(t, totalTime >= waitTime)
}

func TestError(t *testing.T) {
	elem := NewElement("div")
	elem.nodeType = 123
	err := Run(elem)
	assert.Error(t, err)
}
