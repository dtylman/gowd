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
	heading *gowd.Element
	body    *gowd.Element
}

func NewPanel(panelType string) *Panel {
	p := new(Panel)
	p.Element = NewElement("div", "panel " + panelType)
	p.heading = NewElement("div", "panel-heading")
	p.body = NewElement("div", "panel-body")
	p.AddElement(p.heading)
	p.AddElement(p.body)
	return p
}

func (p*Panel)AddToBody(elem *gowd.Element) {
	p.body.AddElement(elem)
}

func (p*Panel) AddToHeading(elem*gowd.Element) {
	p.heading.AddElement(elem)
}

func (p*Panel) AddTitle(title string) *gowd.Element {
	txt := gowd.NewStyledText(title, gowd.Heading3)
	txt.SetClass("panel-title")
	p.AddToHeading(txt)
	return txt
}