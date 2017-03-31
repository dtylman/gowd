package bootstrap

import (
	"github.com/dtylman/gowd"
)

const (
	ButtonDefault = "btn-default"
	ButtonPrimary = "btn-primary"
)

func NewButton(buttontype string, caption string) *gowd.Element {
	btn := NewElement("button", "btn " + buttontype)
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

func NewLinkButton(caption string) *gowd.Element {
	linkBtn := gowd.NewElement("a")
	linkBtn.SetAttribute("href", "#")
	if caption != "" {
		linkBtn.AddElement(gowd.NewText(caption))
	}
	return linkBtn
}
