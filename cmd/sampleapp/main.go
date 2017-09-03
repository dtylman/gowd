package main

import (
	"fmt"
	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
	"math/rand"
	"time"
)

type body struct {
	*gowd.Element
	txt         *gowd.Element
	progressBar *bootstrap.ProgressBar
	btnStart    *gowd.Element
}

func newBody() *body {
	b := new(body)
	b.Element = bootstrap.NewContainer(true)
	b.AddElement(gowd.NewElement("hr"))
	pnl := bootstrap.NewPanel(bootstrap.PanelDefault)
	pnl.AddTitle("Title")
	b.txt = gowd.NewStyledText("Hello world", gowd.Heading1)
	pnl.AddToBody(b.txt)
	b.AddElement(pnl.Element)

	for i := 0; i < 3; i++ {
		input := bootstrap.NewFormInput(bootstrap.InputTypeText, fmt.Sprintf("Question %v:", i))
		input.OnEvent(gowd.OnChange, b.inputChanged)
		b.AddElement(input.Element)
	}

	row := bootstrap.NewRow()
	column := bootstrap.NewColumn(bootstrap.ColumnLarge, 6)
	b.btnStart = bootstrap.NewButton(bootstrap.ButtonPrimary, "Start")
	b.btnStart.OnEvent(gowd.OnClick, b.btnStartClick)
	column.AddElement(b.btnStart)
	row.AddElement(column)

	column = bootstrap.NewColumn(bootstrap.ColumnLarge, 6)
	btnClose := bootstrap.NewLinkButton("Close")
	btnClose.SetAttribute(gowd.OnClick, "window.close();")

	column.AddElement(btnClose)
	row.AddElement(column)

	b.AddElement(row)

	b.AddElement(gowd.NewElement("hr"))

	b.progressBar = bootstrap.NewProgressBar()
	b.progressBar.Hide()
	b.AddElement(b.progressBar.Element)

	b.AddElement(gowd.NewElement("hr"))

	return b
}

func (b *body) btnStartClick(sender *gowd.Element, event *gowd.EventElement) {
	b.AddElement(bootstrap.NewAlert("Started!", fmt.Sprintf("Started on %v", time.Now()), bootstrap.AlertInfo, true))
	go func() {
		b.btnStart.Disable()
		b.progressBar.Show()
		defer func() {
			b.progressBar.Hide()
			b.btnStart.Enable()
			b.Render()
		}()

		for i := 0; i <= 100; i++ {
			b.progressBar.SetText(fmt.Sprintf("Working on it (%v percent done)", i))
			b.progressBar.SetPercent(i)
			time.Sleep(time.Duration(rand.Uint32() / 30))
			b.Render()
		}
	}()
}

func (b *body) inputChanged(sender *gowd.Element, event *gowd.EventElement) {
	b.txt.SetText(fmt.Sprintf("Text from %v: %v", sender.GetID(), sender.Kids[1].GetValue()))
}

func main() {
	gowd.Run(newBody().Element)
}
