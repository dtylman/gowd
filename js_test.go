package gowd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecJSNow(t *testing.T) {
	var buf bytes.Buffer
	Output = io.Writer(&buf)
	ExecJSNow("alert('this  is all the testing possible without running NWJS, hope itl do')")
	assert.Equal(t, "$alert('this  is all the testing possible without running NWJS, hope itl do')"+"\n", buf.String())
}

func TestExecJS(t *testing.T) {
	var buf bytes.Buffer
	Output = io.Writer(&buf)
	ExecJS("alert('this  is all the testing possible without running NWJS, hope itl do')")
	elem := NewElement("hello world")
	elem.SetID("_hello world1")
	render(elem, Output)
	assert.Equal(t, `<hello world id="_hello world1"></hello world>`+"\n"+`$alert('this  is all the testing possible without running NWJS, hope itl do')`+"\n", buf.String())
	Output = os.Stdout
}

func TestAlert(t *testing.T) {
	var buf bytes.Buffer
	Output = io.Writer(&buf)
	Alert("this  is all the testing possible without running NWJS, hope itl do'")
	elem := NewElement("hello world")
	elem.SetID("_hello world1")
	render(elem, Output)
	assert.Equal(t, `<hello world id="_hello world1"></hello world>`+"\n"+`$alert("this  is all the testing possible without running NWJS, hope itl do'");`+"\n", buf.String())
	Output = os.Stdout
}
