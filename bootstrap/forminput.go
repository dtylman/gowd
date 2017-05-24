package bootstrap

import (
	"github.com/dtylman/gowd"
)

//<div class="form-group">
//<label for="exampleInputFile">File input</label>
//<input type="file" id="exampleInputFile">
//<p class="help-block">Example block-level help text here.</p>
//</div>

//FormInput is a bootstrap "form-group" input
type FormInput struct {
	*gowd.Element
	input   *gowd.Element
	txt     *gowd.Element
	helpTxt *gowd.Element
}

//NewFormInput creates a bootstrap "form-group" containing an input with a given type and caption
func NewFormInput(inputType string, caption string) *FormInput {
	i := new(FormInput)
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

//SetPlaceHolder sets the input placeholder text
func (i *FormInput) SetPlaceHolder(placeHolder string) {
	i.input.SetAttribute("placeHolder", placeHolder)
}

//SetHelpText sets the input help text
func (i *FormInput) SetHelpText(help string) {
	i.helpTxt.SetText(help)
	i.helpTxt.Hidden = false
}

//SetValue  sets the input value
func (i *FormInput) SetValue(value string) {
	i.input.SetAttribute("value", value)
}

//GetValue returns the input value
func (i *FormInput) GetValue() string {
	return i.input.GetValue()
}
