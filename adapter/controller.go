package adapter

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

// A WebsocketController is responsible for handling data sent to websocket, and send data to connected clients.
type WebsocketController struct {
	logger   *log.Logger
	clients  chan *WebsocketClient
	received chan ReceivedData
}

// check interface implementation
var _ http.Handler = (*WebsocketController)(nil)

// NewWebsocketController creates new WebsocketController instance.
func NewWebsocketController() *WebsocketController {
	logger := log.New(os.Stderr, "chat-websocket: ", log.Ldate|log.Ltime)

	return &WebsocketController{logger, make(chan *WebsocketClient), make(chan ReceivedData)}
}

// ServeHTTP implements http.Handler interface
func (w *WebsocketController) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	connection := w.createWebsocketConnection(writer, req)
	defer connection.Close()

	client := newWebsocketClient()
	w.clients <- client
	defer w.logger.Println("client", client, "disconnected")

	disconnect := make(chan bool)

	go w.listenForOutcomingData(connection, client)
	go w.listenForIncomingData(connection, client, disconnect)

	<-disconnect
	// TODO: remove client from chat on disconnect
}

func (w *WebsocketController) createWebsocketConnection(writer http.ResponseWriter, req *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		panic(err)
	}

	w.logger.Println("connection opened")

	return connection
}

func (w *WebsocketController) listenForIncomingData(connection *websocket.Conn, client *WebsocketClient, disconnect chan bool) {
	for {
		mt, message, err := connection.ReadMessage()

		if mt == websocket.CloseMessage {
			disconnect <- true
		}
		if _, ok := err.(*websocket.CloseError); ok {
			disconnect <- true
		}

		if err != nil {
			w.logger.Println("warning: read message error:", err)
			continue
		}
		if mt != websocket.TextMessage && mt != websocket.CloseMessage {
			w.logger.Println("warning: recieved unsupported message: ", mt, message)
			continue
		}

		if mt == websocket.TextMessage {
			w.Received() <- ReceivedData{client, message}
		}
	}
}

func (w *WebsocketController) listenForOutcomingData(connection *websocket.Conn, client *WebsocketClient) {
	for {
		received := <-client.In()
		connection.WriteMessage(websocket.TextMessage, []byte(received))

		w.logger.Println("client", client, "recieved", received)
	}
}

func (d *WebsocketController) Received() chan ReceivedData {
	return d.received
}

func (d *WebsocketController) Clients() chan *WebsocketClient {
	return d.clients
}
