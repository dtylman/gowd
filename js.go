package gowd

import "fmt"

//ExecJS executes JS code after a DOM update from gowd
func ExecJS(js string) {
	execJSBuffer = append(execJSBuffer, js)
}

//ExecJSNow Executes JS code in NWJS without waiting for a DOM update to be finished.
func ExecJSNow(js string) {
	fmt.Fprintf(Output, "$%s\n", stripchars(js, '\r', '\n'))
}

//Alert calls javascript alert after DOM update from gowd
func Alert(text string) {
	ExecJS(fmt.Sprintf(`alert("%v");`, text))
}
