package bootstrap

import (
	"github.com/dtylman/pictures/webkit"
)

/*
   <div class="alert alert-warning" role="alert">
  <button type="button" class="close" ><span aria-hidden="true">X</span></button>
  <strong>Warning!</strong> Better check yourself, you're not looking too good.
</div>
*/

const (
	AlertSuccess = "alert-success"
	AlertInfo    = "alert-info"
	AlertWarning = "alert-warning"
	AlertDanger  = "alert-danger"
)

func NewAlert(title string, caption string, alertType string, dismissible bool) *webkit.Element {
	alertClass := "alert " + alertType
	alert := NewElement("div", alertClass)
	alert.SetAttribute("role", "alert")
	if dismissible {
		span := webkit.NewElement("span")
		span.SetAttribute("aria-hidden", "true")
		span.SetText("X")
		btn := NewElement("button", "close")
		btn.SetAttribute("type", "button")
		btn.OnEvent(webkit.OnClick, alertClose)
		btn.AddElement(span)
		alert.AddElement(btn)
	}
	if title != "" {
		alert.AddElement(webkit.NewStyledText(title+" ", webkit.StrongText))
	}
	alert.AddElement(webkit.NewText(caption))
	return alert
}

func alertClose(sender *webkit.Element, event *webkit.EventElement) {
	//remove this alert
	if sender != nil {
		//the button
		if sender.Parent != nil {
			// the alert div
			if sender.Parent.Parent != nil {
				// whatever holds the alert
				sender.Parent.Parent.RemoveElement(sender.Parent)
			}
		}
	}
}
