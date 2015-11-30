package main

import (
	"encoding/json"
	"github.com/roblaszczak/simple-go-chat/chat"
)

// A MessagesSender is type responsible for sending messages to user via SenderDriver (for example websocket).
type MessagesSender struct {
	driver SenderDriver
}

// A SenderDriver is interface for objects with can handle sending messages to chat user (for example websocket).
type SenderDriver interface {
	sendData(data string)
}

func (chat *MessagesSender) send(message chat.MessageData) {
	json_message, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	chat.driver.sendData(string(json_message))
}
