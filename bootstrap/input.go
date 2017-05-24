package bootstrap

import "github.com/dtylman/gowd"

//NewInputGroup creates new bootsrap input group from the given elements
func NewInputGroup(elems ...*gowd.Element) *gowd.Element {
	inputGroup := NewElement("div", "input-group")
	for _, elem := range elems {
		inputGroup.AddElement(elem)
	}
	return inputGroup
}

const (
	//InputTypeText is <input type=text>
	InputTypeText = "text"
	//InputTypeFile is <input type=file>
	InputTypeFile = "file"
	//InputTypeCheckbox is <input type=checkbox>
	InputTypeCheckbox = "checkbox"
)

//NewInput creates a new input with a provided type
func NewInput(inputType string) *gowd.Element {
	input := gowd.NewElement("input")
	input.SetAttribute("type", inputType)
	return input
}
