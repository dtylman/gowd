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
	PanelDefault = "panel-default"
)

type Panel struct {
	*gowd.Element
	Heading *gowd.Element
	Body    *gowd.Element
}

func NewPanel(panelType string) *Panel {
	p := new(Panel)
	p.Element = NewElement("div", "panel " + panelType)
	p.Heading = NewElement("div", "panel-heading")
	p.Body = NewElement("div", "panel-body")
	p.AddElement(p.Heading)
	p.AddElement(p.Body)
	return p
}

func (p*Panel)AddToBody(elem *gowd.Element) {
	p.Body.AddElement(elem)
}

func (p*Panel) AddToHeading(elem*gowd.Element) {
	p.Heading.AddElement(elem)
}

func (p*Panel) AddTitle(title string) *gowd.Element {
	txt := gowd.NewStyledText(title, gowd.Heading3)
	txt.SetClass("panel-title")
	p.AddToHeading(txt)
	return txt
}