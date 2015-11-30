// This package contains client-side scripts, with can be dumped into
// JavaScript using gopherjs.
//
// Dump of JavascriptFiles can be done by command
//
//     gopherjs build js/*.go --output=public/app.js
//
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
