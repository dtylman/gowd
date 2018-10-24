package gowd

import "golang.org/x/net/html"

const (
	//BoldText <b>
	BoldText = "b"
	//StrongText <strong>
	StrongText = "strong"
	//ItalicText <i>
	ItalicText = "i"
	//EmphasizedText <em>
	EmphasizedText = "em"
	//MarkedText <mark>
	MarkedText = "mark"
	//SmallText <small>
	SmallText = "small"
	//DeletedText <del>
	DeletedText = "del"
	//InsertedText <ins>
	InsertedText = "ins"
	//SubscriptText <sub>
	SubscriptText = "sub"
	//SuperscriptText <sup>
	SuperscriptText = "sup"
	//TitleText <title>
	TitleText = "title"
	//Paragraph <p>
	Paragraph = "p"
	//Heading1 <h1>
	Heading1 = "h1"
	//Heading2 <h2>
	Heading2 = "h2"
	//Heading3 <h3>
	Heading3 = "h3"
	//Heading4 <h4>
	Heading4 = "h4"
	//Heading5 <h5>
	Heading5 = "h5"
	//Heading6 <h6>
	Heading6 = "h6"
)

//NewText creates new text node (without HTML tag)
func NewText(text string) *Element {
	return &Element{nodeType: html.TextNode, data: elementText(text)}
}

//NewStyledText creates new text element using a specific style
func NewStyledText(text string, style string) *Element {
	txt := NewElement(style)
	txt.AddElement(NewText(text))
	return txt
}
