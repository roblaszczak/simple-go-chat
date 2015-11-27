package websocket

// A WebsocketClient represents client connected to websocket, and provides acces for client input channel.
type WebsocketClient struct {
	in chan string
}

func newWebsocketClient() *WebsocketClient {
	return &WebsocketClient{make(chan string)}
}

// In returns channel with allow to send data to connected client.
func (c WebsocketClient) In() chan string {
	return c.in
}
