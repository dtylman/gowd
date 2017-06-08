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

func TestElement_ProcessEvent(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div><input id="input" type="text" value="lala"><btn id="button"></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	input := em["input"]
	assert.EqualValues(t, input.GetValue(), "lala")
	em["button"].OnEvent("onclick", func(sender *Element, event *EventElement) {
		assert.EqualValues(t, sender.GetID(), "button")
		assert.EqualValues(t, input.GetValue(), "shalom")
	})
	testOuput(t, elem, `<div><input id="input" type="text" value="lala"/><btn id="button" onclick="fire_event(&#39;onclick&#39;,this);"></btn></div>`)
	jsEvent := `
	{
	   "name":"onclick",
	   "sender":{
	      "properties":{
		 "id":"button"
	      }
	   },
	   "inputs":[
	      {
		 "properties":{
		    "id":"input",
		    "type":"text",
		    "value":"shalom"
		 }
	      },
	      {
		 "properties":{
		    "id":"_input14",
		    "type":"text",
		    "value":""EqualValues(t, )
		 }
	      },
	      {
		 "properties":{
		    "id":"button",
		    "class":"klass"
		 }
	      }
	   ]
	}`
	err = processEvents(elem, bytes.NewBufferString(jsEvent))
	if err != nil {
		t.Fatal(err)
	}
	class, found := elem.Find("button").GetAttribute("class")
	assert.EqualValues(t, class, "klass")
	assert.True(t, found)
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
