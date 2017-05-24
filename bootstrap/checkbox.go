package bootstrap

import (
	"github.com/dtylman/gowd"
)

//<div class="checkbox">
//<label>
//<input type="checkbox"> Check me out
//</label>
//</div>

//Checkbox represents a checkbox with label
type Checkbox struct {
	*gowd.Element
	chkbox *gowd.Element
	txt    *gowd.Element
}

//NewCheckBox creates a bootstrap checkbox with label
func NewCheckBox(caption string, checked bool) *Checkbox {
	cb := new(Checkbox)
	cb.Element = NewElement("div", "checkbox")
	lbl := gowd.NewElement("label")
	cb.chkbox = gowd.NewElement("input")
	cb.chkbox.SetAttribute("type", "checkbox")
	if checked {
		cb.chkbox.SetAttribute("checked", "")
	}
	lbl.AddElement(cb.chkbox)
	cb.txt = gowd.NewText(caption)
	lbl.AddElement(cb.txt)
	lbl.SetAttribute("for", cb.chkbox.GetID())
	cb.Element.AddElement(lbl)
	return cb
}

//Checked is true if the checkbox is checked
func (cb *Checkbox) Checked() bool {
	_, exists := cb.chkbox.GetAttribute("checked")
	return exists
}
