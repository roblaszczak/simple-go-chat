package main

import (
	"encoding/json"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"github.com/roblaszczak/go-chat/chat"
	"github.com/roblaszczak/go-chat/config"
	"time"
)

// NewApp creates new AngularJS chat application.
func NewApp(name string, modules AngularAppModules) *AngularApp {
	app := CreateAngularApp(name, modules)
	app.Call("controller", "ChatCtrl", js.S{"$scope", func(scope *js.Object) {
		newChatController(scope)
	}})

	return app
}

type chatController struct {
	scope *js.Object
	chat  *MessagesSender
}

func newChatController(scope *js.Object) *chatController {
	controller := &chatController{scope, &MessagesSender{createJsWebsocket(scope)}}
	controller.setDefaults()

	scope.Set("sendMessage", func() {
		controller.sendMessage()
	})
	scope.Set("showMessage", func(message *js.Object) {
		controller.showMessage(message)
	})

	return controller
}

func (c *chatController) setDefaults() {
	c.scope.Set("messages", []chat.MessageData{})
	c.scope.Set("nick", "anonymus1")
}

func (c chatController) sendMessage() {
	author := c.scope.Get("nick").String()
	messageContent := c.scope.Get("message").String()

	if len(messageContent) == 0 {
		return
	}

	message := chat.MessageData{author, messageContent, time.Now()}
	c.chat.send(message)
	c.scope.Set("message", "")
}

func (c chatController) showMessage(message *js.Object) {
	messages := &messages{c.scope.Get("messages")}
	messages.Add(message)
}

type messages struct {
	*js.Object
}

func (m *messages) Add(message *js.Object) {
	m.Call("push", message)
}

func createJsWebsocket(scope *js.Object) *Websocket {
	websocketUrl := fmt.Sprintf("ws://%s:%d/chat", config.SERVER_HOST, config.SERVER_PORT)
	websocket := NewWebsocket(websocketUrl, func(event *js.Object) {
		stringData := event.Get("data").String()

		message := chat.MessageData{}
		json.Unmarshal([]byte(stringData), &message)

		scope.Call("showMessage", message)
		// force refresh scope data for async data binding
		scope.Call("$apply")
		ScrollToChatBottom()
	})

	return websocket
}
