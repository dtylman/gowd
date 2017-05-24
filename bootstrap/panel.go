package bootstrap

import (
	"github.com/dtylman/gowd"
)

/*
<div class="panel panel-default">
  <div class="panel-heading">
    <h3 class="panel-title">Panel title</h3>
  </div>
  <div class="panel-body">
    Panel content
  </div>
</div>
*/

const (
	//PanelDefault bootstrap default paenl
	PanelDefault = "panel-default"
)

//Panel is a bootstrap panel
type Panel struct {
	*gowd.Element
	Heading *gowd.Element
	Body    *gowd.Element
}

//NewPanel creates a new bootstrap panel with the provided type.
func NewPanel(panelType string) *Panel {
	p := new(Panel)
	p.Element = NewElement("div", "panel "+panelType)
	p.Heading = NewElement("div", "panel-heading")
	p.Body = NewElement("div", "panel-body")
	p.AddElement(p.Heading)
	p.AddElement(p.Body)
	return p
}

//AddToBody  adds element to panel body
func (p *Panel) AddToBody(elem *gowd.Element) {
	p.Body.AddElement(elem)
}

//AddToHeading adds element to panel heading
func (p *Panel) AddToHeading(elem *gowd.Element) {
	p.Heading.AddElement(elem)
}

//AddTitle sets or adds the panel title.
func (p *Panel) AddTitle(title string) *gowd.Element {
	txt := gowd.NewStyledText(title, gowd.Heading3)
	txt.SetClass("panel-title")
	p.AddToHeading(txt)
	return txt
}
