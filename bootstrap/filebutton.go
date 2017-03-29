package bootstrap

import (
	"fmt"
	"github.com/dtylman/pictures/webkit"
)

/*<div>
  <input type="file" id="lala" style="display: none;" onchange="alert(lala.value)" />
  <button onclick="lala.click();"> click </button></div>
*/
type FileButton struct {
	*webkit.Element
	btn   *webkit.Element
	input *webkit.Element
}

func NewFileButton(buttontype string, caption string, foldersOnly bool) *FileButton {
	fb := new(FileButton)
	fb.Element = webkit.NewElement("div")
	fb.btn = NewButton(buttontype, caption)
	fb.input = webkit.NewElement("input")
	fb.input.SetAttribute("type", "file")
	fb.input.SetAttribute("style", "display:none;")
	if foldersOnly {
		fb.input.SetAttribute("nwdirectory", "")
	}
	fb.btn.SetAttribute(webkit.OnClick, fmt.Sprintf("%s.click()", fb.input.GetID()))
	fb.AddElement(fb.input)
	fb.AddElement(fb.btn)
	return fb

}

func (fb *FileButton) OnChange(handler webkit.EventHandler) {
	fb.input.OnEvent(webkit.OnChange, handler)
}

func (fb *FileButton) GetValue() string {
	return fb.input.GetValue()
}
