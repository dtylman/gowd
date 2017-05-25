package gowd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"reflect"
	"sync"
)

var renderMutex sync.Mutex

func render(e *Element, w io.Writer) error {
	node := e.toNode()
	h := md5.New()
	err := html.Render(h, node)
	if err != nil {
		return err
	}
	sum := h.Sum(nil)
	if reflect.DeepEqual(e.renderHash, sum) {
		return nil //already rendered
	}
	e.renderHash = sum
	renderMutex.Lock()
	defer renderMutex.Unlock()
	err = html.Render(w, node)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w)
	return err
}

func processEvents(e *Element, r io.Reader) error {
	decoder := json.NewDecoder(r)
	var event Event
	err := decoder.Decode(&event)
	if err != nil {
		return err
	}
	e.ProcessEvent(&event)
	return nil
}

//Run starts the message loop with body as the root element. This function never exits.
func Run(body *Element) error {
	for true {
		err := body.Render()
		if err != nil {
			return err
		}
		err = processEvents(body, os.Stdin)
		if err != nil {
			return err
		}
	}
	return nil
}
