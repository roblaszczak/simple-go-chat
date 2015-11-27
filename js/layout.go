package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

const (
	CHAT_SELECTOR         = ".chat"
	CHAT_WRAPPER_SELECTOR = ".chat-wrapper"
	CHAT_PADDING          = 190
)

// CalculateChatHeight calculates and set chat div height based on window height.
func CalculateChatHeight() {
	chatHeight := getWindowHeight() - CHAT_PADDING
	setChatHeight(chatHeight)
}

// ScrollToChatBottom scroll chat messages overflow to bottom.
func ScrollToChatBottom() {
	scrollHeight := jQuery(CHAT_SELECTOR).Get(0).Get("scrollHeight")
	jQuery(CHAT_SELECTOR).Call("scrollTop", scrollHeight)
}

func getWindowHeight() int64 {
	return js.Global.Get("window").Get("innerHeight").Int64()
}

func setChatHeight(height int64) {
	jQuery(CHAT_WRAPPER_SELECTOR).Call("css", "height", fmt.Sprintf("%dpx", height))
}
