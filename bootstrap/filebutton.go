package bootstrap

import (
	"fmt"
	"github.com/dtylman/gowd"
)

/*<div>
  <input type="file" id="lala" style="display: none;" onchange="alert(lala.value)" />
  <button onclick="lala.click();"> click </button></div>
*/

//FileButton is an file-input linked to a button
type FileButton struct {
	*gowd.Element
	btn   *gowd.Element
	input *gowd.Element
}

//NewFileButton creates new 'file' input with a button
func NewFileButton(buttontype string, caption string, foldersOnly bool) *FileButton {
	fb := new(FileButton)
	fb.Element = gowd.NewElement("div")
	fb.btn = NewButton(buttontype, caption)
	fb.input = gowd.NewElement("input")
	fb.input.SetAttribute("type", "file")
	fb.input.SetAttribute("style", "display:none;")
	if foldersOnly {
		fb.input.SetAttribute("nwdirectory", "")
	}
	fb.btn.SetAttribute(gowd.OnClick, fmt.Sprintf("%s.click()", fb.input.GetID()))
	fb.AddElement(fb.input)
	fb.AddElement(fb.btn)
	return fb

}

//OnChange registers the onchange event
func (fb *FileButton) OnChange(handler gowd.EventHandler) {
	fb.input.OnEvent(gowd.OnChange, handler)
}

//GetValue returns the selected file value
func (fb *FileButton) GetValue() string {
	return fb.input.GetValue()
}
