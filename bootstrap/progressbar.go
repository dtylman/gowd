package bootstrap

import (
	"errors"
	"fmt"
	"github.com/dtylman/pictures/webkit"
	"strconv"
)

//<div class="progress">
//<div class="progress-bar" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">
//60%
//</div>
//</div>

type ProgressBar struct {
	*webkit.Element
	bar *webkit.Element
	txt *webkit.Element
}

func NewProgressBar() *ProgressBar {
	progress := new(ProgressBar)
	progress.Element = NewElement("div", "progress")
	progress.bar = NewElement("div", "progress-bar")
	progress.bar.SetAttribute("role", "progressbar")
	progress.bar.SetAttribute("aria-valuemin", "0")
	progress.bar.SetAttribute("aria-valuemax", "100")
	progress.txt = webkit.NewText("")
	progress.bar.AddElement(progress.txt)
	progress.AddElement(progress.bar)
	return progress
}

func (pb *ProgressBar) SetValue(percent int) error {
	if percent < 0 || percent > 100 {
		return errors.New("out of range")
	}
	pb.bar.SetAttribute("aria-valuenow", strconv.Itoa(percent))
	pb.bar.SetAttribute("style", fmt.Sprintf("width: %d%%;", percent))
	return nil
}

func (pb *ProgressBar) SetText(caption string) {
	pb.txt.SetText(caption)
}
