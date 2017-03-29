package gowd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"reflect"
	"sync"
	"time"
)

var renderMutex sync.Mutex

func render(e *Element) error {
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
	err = html.Render(os.Stdout, node)
	if err != nil {
		return err
	}
	_, err = fmt.Println()
	return err
}

func Run(body *Element) error {
	for true {
		body.Render()
		err := body.Render()
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(os.Stdin)
		var event Event
		err = decoder.Decode(&event)
		if err != nil {
			return err
		}
		body.ProcessEvent(&event)
	}
	return nil
}

func Error(err error) {
	fmt.Println("Error: ", err.Error())
	time.Sleep(3 * time.Second)
}
