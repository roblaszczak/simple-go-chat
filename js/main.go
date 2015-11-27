package main

import (
	"github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery

func main() {
	CalculateChatHeight()
	BindOnWindowResize(func() {
		CalculateChatHeight()
	})
	NewApp("chatApp", []string{})
}
