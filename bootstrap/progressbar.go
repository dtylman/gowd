package bootstrap

import (
	"errors"
	"fmt"
	"github.com/dtylman/gowd"
	"strconv"
)

//<div class="progress">
//<div class="progress-bar" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">
//60%
//</div>
//</div>

//ProgressBar represents bootstrap progress-bar element
type ProgressBar struct {
	*gowd.Element
	bar *gowd.Element
	txt *gowd.Element
}

//NewProgressBar creates new bootstrap progress bar element
func NewProgressBar() *ProgressBar {
	progress := new(ProgressBar)
	progress.Element = NewElement("div", "progress")
	progress.bar = NewElement("div", "progress-bar")
	progress.bar.SetAttribute("role", "progressbar")
	progress.bar.SetAttribute("aria-valuemin", "0")
	progress.bar.SetAttribute("aria-valuemax", "100")
	progress.txt = gowd.NewText("")
	progress.bar.AddElement(progress.txt)
	progress.AddElement(progress.bar)
	return progress
}

//SetValue sets the progress value of the progress bar
func (pb *ProgressBar) SetValue(now, max int) error {
	if max == 0 {
		return pb.SetPercent(0)
	}
	return (pb.SetPercent(now * 100 / max))
}

//SetPercent sets the value of the progress bar as a percentage
func (pb *ProgressBar) SetPercent(percent int) error {
	if percent < 0 || percent > 100 {
		return errors.New("out of range")
	}
	pb.bar.SetAttribute("aria-valuenow", strconv.Itoa(percent))
	pb.bar.SetAttribute("style", fmt.Sprintf("width: %d%%;", percent))
	return nil
}

//SetText sets the progress bar's caption
func (pb *ProgressBar) SetText(caption string) {
	pb.txt.SetText(caption)
}
