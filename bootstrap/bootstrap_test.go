package bootstrap

import (
	"bytes"
	"testing"

	"github.com/dtylman/gowd"
	"github.com/stretchr/testify/assert"
)

func testOuput(t *testing.T, elem *gowd.Element, expected string) {
	output := bytes.NewBuffer(make([]byte, 0))
	gowd.Output = output
	gowd.Order = 0
	elem.Render()
	assert.Equal(t, expected+"\n", output.String())
}

func TestNewAlert(t *testing.T) {
	alert := NewAlert("alert", "lala", AlertInfo, true)
	expected := `<div id="_div1" class="alert alert-info" role="alert"><button id="_button3" class="close" type="button" onclick="fire_event(&#39;onclick&#39;,this);"><span id="_span2" aria-hidden="true">X</span></button><strong id="_strong4">alert</strong>lala</div>`
	testOuput(t, alert, expected)
}

func TestNewButton(t *testing.T) {
	button := NewButton(ButtonPrimary, "click")
	expected := `<button id="_button1" class="btn btn-primary">click</button>`
	testOuput(t, button, expected)
}

func TestNewCell(t *testing.T) {
	cell := NewCell("hoho")
	expected := `<td id="_td1">hoho</td>`
	testOuput(t, cell, expected)
}

func TestNewCheckBox(t *testing.T) {
	elem := NewCheckBox("check", true).Element
	expected := `<div id="_div1" class="checkbox"><label id="_label2" for="_input3"><input id="_input3" type="checkbox" checked=""/>check</label></div>`
	testOuput(t, elem, expected)

}

func TestNewColumn(t *testing.T) {
	elem := NewColumn(ColumnSmall, 3, gowd.NewText("one"), gowd.NewText("two"))
	expected := `<div id="_div1" class="col-sm-3">onetwo</div>`
	testOuput(t, elem, expected)
}

func TestNewRow(t *testing.T) {
	elem := NewRow(NewContainer(false), NewContainer(true))
	expected := `<div id="_div3" class="row"><div id="_div1" class="container"></div><div id="_div2" class="container-fluid"></div></div>`
	testOuput(t, elem, expected)
}

func TestNewContainer(t *testing.T) {
	elem := NewContainer(true)
	expected := `<div id="_div1" class="container-fluid"></div>`
	testOuput(t, elem, expected)
}

func TestNewElement(t *testing.T) {
	elem := NewElement("div", "well", gowd.NewText("well done"))
	expected := `<div id="_div1" class="well">well done</div>`
	testOuput(t, elem, expected)
}

func TestNewFormGroup(t *testing.T) {
	elem := NewFormGroup(gowd.NewText("test"))
	expected := `<div id="_div1" class="form-group">test</div>`
	testOuput(t, elem, expected)
}

func TestNewFileButton(t *testing.T) {
	elem := NewFileButton(ButtonPrimary, "File", true).Element
	expected := `<div id="_div1"><input id="_input3" type="file" style="display:none;" nwdirectory=""/><button id="_button2" class="btn btn-primary" onclick="_input3.click()">File</button></div>`
	testOuput(t, elem, expected)
}
