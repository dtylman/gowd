package bootstrap

import "github.com/dtylman/pictures/webkit"

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
	*webkit.Element
	heading *webkit.Element
	body    *webkit.Element
}

func NewPanel(panelType string) *Panel {
	p := new(Panel)
	p.Element = NewElement("div", "panel " + panelType)
	p.heading = NewElement("div", "panel-heading")
	p.body = NewElement("div", "panel-body")
	return p
}

func (p*Panel) AddElement(elem*webkit.Element) {
	p.AddBody(elem)
}

func (p*Panel)AddBody(elem *webkit.Element) {
	p.body.AddElement(elem)
}

func (p*Panel) AddHeading(elem*webkit.Element) {
	p.heading.AddElement(elem)
}

func (p*Panel) AddTitle(title string) {
	txt := webkit.NewStyledText(title, webkit.Heading3)
	txt.SetClass("panel-title")
	p.AddHeading(txt)
}