package bootstrap

import (
	"github.com/dtylman/pictures/webkit"
)

const (
	ButtonDefault = "btn-default"
	ButtonPrimary = "btn-primary"
)

func NewButton(buttontype string, caption string) *webkit.Element {
	btn := NewElement("button", "btn "+buttontype)
	btn.SetText(caption)
	return btn
}

func NewLinkButton(caption string) *webkit.Element {
	linkBtn := webkit.NewElement("a")
	linkBtn.SetAttribute("href", "#")
	linkBtn.AddElement(webkit.NewText(caption))
	return linkBtn
}
