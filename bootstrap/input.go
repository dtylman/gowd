package bootstrap

import (
	"github.com/dtylman/gowd"
)

//<div class="form-group">
//<label for="exampleInputFile">File input</label>
//<input type="file" id="exampleInputFile">
//<p class="help-block">Example block-level help text here.</p>
//</div>

const (
	InputTypeText = "text"
	InputTypeFile = "file"
)

type Input struct {
	*gowd.Element
	input   *gowd.Element
	txt     *gowd.Element
	helpTxt *gowd.Element
}

func NewInput(inputType string, caption string) *Input {
	i := new(Input)
	i.Element = NewElement("div", "form-group")
	lbl := gowd.NewElement("label")
	i.txt = gowd.NewText(caption)
	lbl.AddElement(i.txt)
	i.input = NewElement("input", "form-control")
	i.input.SetAttribute("type", inputType)
	i.helpTxt = NewElement("p", "help-block")
	i.AddElement(lbl)
	i.AddElement(i.input)
	i.AddElement(i.helpTxt)
	lbl.SetAttribute("for", i.input.GetID())
	i.helpTxt.Hidden = true
	return i
}

func (i *Input) SetPlaceHolder(placeHolder string) {
	i.input.SetAttribute("placeHolder", placeHolder)
}

func (i *Input) SetHelpText(help string) {
	i.helpTxt.SetText(help)
	i.helpTxt.Hidden = false
}

func (i *Input) SetValue(value string) {
	i.input.SetAttribute("value", value)
}

func (i *Input) GetValue() string {
	return i.input.GetValue()
}
