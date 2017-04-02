package gowd

import "golang.org/x/net/html"

const (
	BoldText = "b"
	StrongText = "strong"
	ItalicText = "i"
	EmphasizedText = "em"
	MarkedText = "mark"
	SmallText = "small"
	DeletedText = "del"
	InsertedText = "ins"
	SubscriptText = "sub"
	SuperscriptText = "sup"
	TitleText = "title"
	Paragraph = "p"
	Heading1 = "h1"
	Heading2 = "h2"
	Heading3 = "h3"
	Heading4 = "h4"
	Heading5 = "h5"
	Heading6 = "h6"
)

func NewText(text string) *Element {
	return &Element{nodeType: html.TextNode, data: text}
}

func NewStyledText(text string, style string) *Element {
	txt := NewElement(style)
	txt.AddElement(NewText(text))
	return txt
}

