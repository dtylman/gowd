package bootstrap

import "github.com/dtylman/gowd"

func NewInputGroup(elems ...*gowd.Element) *gowd.Element {
	inputGroup := NewElement("div", "input-group")
	for _, elem := range
		elems {
		inputGroup.AddElement(elem)
	}
	return inputGroup
}

const (
	InputTypeText = "text"
	InputTypeFile = "file"
)

func NewInput(inputType string) *gowd.Element {
	input := gowd.NewElement("input")
	input.SetAttribute("type", inputType)
	return input
}