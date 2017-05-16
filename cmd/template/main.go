package main

import (
	"github.com/dtylman/gowd"
)

func main() {
	body, err := gowd.ParseElement("<h1>Hello world</h1>", nil)
	if err != nil {
		panic(err)
	}
	gowd.Run(body)
}
