package main

import (
	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
)

func main() {
	panel := bootstrap.NewPanel(bootstrap.PanelDefault)
	gowd.Run(panel.Element)
}
