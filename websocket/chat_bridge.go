package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/roblaszczak/go-chat/chat"
	"log"
	"os"
)

// A ChatBridge is responsible for connecting chat layer with websocket layer, and keep this layers separated.
type ChatBridge struct {
	controller  ChatController
	chatCore    *chat.Chat
	chatClients map[*WebsocketClient]*chat.Client
	logger      *log.Logger
}

// NewChatBridge create new instance of ChatBridge
func NewChatBridge(controller ChatController, chatCore *chat.Chat) *ChatBridge {
	logger := log.New(os.Stderr, "chat-bridge: ", log.Ldate|log.Ltime|log.Lshortfile)

	bridge := &ChatBridge{controller, chatCore, make(map[*WebsocketClient]*chat.Client), logger}
	return bridge
}

// Listen runs processes responsive for handling data sent to websocket, and sends it to chat layer.
func (b *ChatBridge) Listen() {
	go b.listenForNewClients()
	go b.listenForOutgoingMessages()
	go b.listenForIncomingMessages()
}

func (b *ChatBridge) listenForNewClients() {
	for {
		websocketClient := <-b.controller.Clients()
		chatClient := chat.NewClient(chat.RandomNick())
		b.chatCore.ConnectClient(chatClient)
		b.chatClients[websocketClient] = chatClient
	}
}

func (b *ChatBridge) listenForIncomingMessages() {
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

func (b *ChatBridge) listenForOutgoingMessages() {
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

func (b *ChatBridge) getChatClient(websocketClient *WebsocketClient) (*chat.Client, error) {
	chatClient, ok := b.chatClients[websocketClient]
	if !ok {
		return nil, errors.New("nie znaleziono klienta")
	}

	return chatClient, nil
}

func (b *ChatBridge) getWebsocketClient(chatClient *chat.Client) (*WebsocketClient, error) {
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
