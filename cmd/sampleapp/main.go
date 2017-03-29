package main

import (
	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
)

type body struct {
	*gowd.Element
}

func newBody() *body {
	b := new(body)
	b.Element = bootstrap.NewContainer(true)
	pnl := bootstrap.NewPanel(bootstrap.PanelDefault)
	pnl.AddTitle("Title")
	pnl.AddBody(gowd.NewStyledText("Hello world", gowd.Heading1))
	b.AddElement(pnl.Element)
	return b
}

func main() {
	gowd.Run(newBody().Element)
}
