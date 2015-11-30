package adapter

// A ReceivedData struct represents data sent to websocket
type ReceivedData struct {
	Connection *WebsocketClient
	Data       []byte
}
