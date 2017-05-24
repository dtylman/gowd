package bootstrap

import (
	"github.com/dtylman/gowd"
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

//NewAlert returns new bootstrap alert
func NewAlert(title string, caption string, alertType string, dismissible bool) *gowd.Element {
	alertClass := "alert " + alertType
	alert := NewElement("div", alertClass)
	alert.SetAttribute("role", "alert")
	if dismissible {
		span := gowd.NewElement("span")
		span.SetAttribute("aria-hidden", "true")
		span.SetText("X")
		btn := NewElement("button", "close")
		btn.SetAttribute("type", "button")
		btn.OnEvent(gowd.OnClick, alertClose)
		btn.AddElement(span)
		alert.AddElement(btn)
	}
	if title != "" {
		alert.AddElement(gowd.NewStyledText(title+" ", gowd.StrongText))
	}
	alert.AddElement(gowd.NewText(caption))
	return alert
}

func alertClose(sender *gowd.Element, event *gowd.EventElement) {
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
