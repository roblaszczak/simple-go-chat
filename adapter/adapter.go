// Package adapter is responsible for translating incoming data (for example
// from websocket) to application operations, and sending outcoming messages to
// logged chat clients.
//
package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/roblaszczak/simple-go-chat/chat"
	"log"
	"os"
)

// A Websocket is responsible for connecting chat layer with websocket layer, and keep this layers separated.
type Websocket struct {
	controller  ChatController
	chatCore    *chat.Chat
	chatClients map[*WebsocketClient]*chat.Client
	logger      *log.Logger
}

// NewWebsocket create new instance of ChatBridge
func NewWebsocket(controller ChatController, chatCore *chat.Chat) *Websocket {
	logger := log.New(os.Stderr, "chat-bridge: ", log.Ldate|log.Ltime|log.Lshortfile)

	bridge := &Websocket{controller, chatCore, make(map[*WebsocketClient]*chat.Client), logger}
	return bridge
}

// Listen runs processes responsive for handling data sent to websocket, and sends it to chat layer.
func (b *Websocket) Listen() {
	go b.listenForNewClients()
	go b.listenForOutgoingMessages()
	go b.listenForIncomingMessages()
}

func (b *Websocket) listenForNewClients() {
	for {
		websocketClient := <-b.controller.Clients()
		chatClient := chat.NewClient(chat.RandomNick())
		b.chatCore.ConnectClient(chatClient)
		b.chatClients[websocketClient] = chatClient
	}
}

func (b *Websocket) listenForIncomingMessages() {
	for {
		received := <-b.controller.Received()
		fmt.Println("received:", received.Data)

		message := chat.MessageData{}
		err := json.Unmarshal(received.Data, &message)
		if err != nil {
			b.logger.Println("error: received invalid data: ", err, string(received.Data))
		}

		chatClient, err := b.getChatClient(received.Connection)
		if err != nil {
			b.logger.Println("error: client not found, connection:", received.Connection)
			continue
		}

		incomingMessage := chat.IncomingMessage{chatClient, message.Content, message.Time}
		b.chatCore.ReceiveMessage(incomingMessage)
	}
}

func (b *Websocket) listenForOutgoingMessages() {
	for {
		message := <-b.chatCore.OutgoingMessages()

		websocketClient, err := b.getWebsocketClient(message.Receiver)
		if err != nil {
			b.logger.Println("error: client not found, reeiver:", message.Receiver)
			continue
		}

		serializedMessage, err := json.Marshal(chat.MessageData{message.Author.Nick(), message.Content, message.Time})
		if err != nil {
			b.logger.Println("error:", err)

		}

		websocketClient.In() <- string(serializedMessage)
	}
}

func (b *Websocket) getChatClient(websocketClient *WebsocketClient) (*chat.Client, error) {
	chatClient, ok := b.chatClients[websocketClient]
	if !ok {
		return nil, errors.New("nie znaleziono klienta")
	}

	return chatClient, nil
}

func (b *Websocket) getWebsocketClient(chatClient *chat.Client) (*WebsocketClient, error) {
	for mapWebsocketClient, mapChatClient := range b.chatClients {
		if chatClient == mapChatClient {
			return mapWebsocketClient, nil
		}
	}
	return nil, errors.New("nie znaleziono klienta")
}

// A ChatController interface is implemented by objects that can communicate with chat client (for example websocket).
type ChatController interface {
	Received() chan ReceivedData
	Clients() chan *WebsocketClient
}
