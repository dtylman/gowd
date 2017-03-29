package bootstrap

import "github.com/dtylman/pictures/webkit"

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
	*webkit.Element
	input   *webkit.Element
	txt     *webkit.Element
	helpTxt *webkit.Element
}

func NewInput(inputType string, caption string) *Input {
	i := new(Input)
	i.Element = NewElement("div", "form-group")
	lbl := webkit.NewElement("label")
	i.txt = webkit.NewText(caption)
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
