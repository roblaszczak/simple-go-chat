package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

type Websocket struct {
	*js.Object
}

// interface check
var _ SenderDriver = (*Websocket)(nil)

type OnWebsocketMessage func(event *js.Object)

// NewWebSocket create new instance of WebSocket with provided callbacks.
func NewWebsocket(url string, onMessage OnWebsocketMessage) *Websocket {
	websocket := &Websocket{js.Global.Get("WebSocket").New(url)}
	websocket.setCallbacks(onMessage)

	return websocket
}

func (ws *Websocket) sendData(data string) {
	ws.Call("send", data)
}

func (w *Websocket) setCallbacks(onMessage OnWebsocketMessage) {
	w.Set("onopen", func() {
		// TODO: something else...
		fmt.Println("connected")
	})
	w.Set("onmessage", onMessage)
}
