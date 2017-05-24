package bootstrap

import (
	"github.com/dtylman/gowd"
)

const (
	//ButtonDefault default bootstrap button
	ButtonDefault = "btn-default"
	//ButtonPrimary primary bootstrap button
	ButtonPrimary = "btn-primary"
)

//NewButton creates a new bootstrap <button> element
func NewButton(buttontype string, caption string) *gowd.Element {
	btn := NewElement("button", "btn "+buttontype)
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

//NewLinkButton creates a new bootstrap link button (<a>)
func NewLinkButton(caption string) *gowd.Element {
	linkBtn := gowd.NewElement("a")
	linkBtn.SetAttribute("href", "#")
	if caption != "" {
		linkBtn.AddElement(gowd.NewText(caption))
	}
	return linkBtn
}
