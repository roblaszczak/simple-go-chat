package chat

import (
	"errors"
	"log"
	"os"
)

// A Chat struct is a type with represents chat domain.
// It is fully independent from other components of app.
type Chat struct {
	clients          []*Client
	bot              *Client
	logger           *log.Logger
	outgoingMessages chan OutgoingMessage
}

// NewChat returns new Chat instance.
func NewChat() *Chat {
	logger := log.New(os.Stderr, "chat-websocket: ", log.Ldate|log.Ltime)

	chat := &Chat{logger: logger, outgoingMessages: make(chan OutgoingMessage)}
	err := chat.addBotClient()
	if err != nil {
		panic(err)
	}

	logger.Println("chat initialized")

	return chat
}

// ConnectClient connects new client to chat, and send him welcome message.
func (c *Chat) ConnectClient(client *Client) {
	c.addClient(client)
	c.sendHelloMessage(client)
}

func (c *Chat) addClient(client *Client) {
	c.clients = append(c.clients, client)

	c.logger.Print(client.Nick(), " connected")
	c.deliverClientList()
}

func (c Chat) sendHelloMessage(client *Client) {
	message := NewOutgoingMessage(c.bot, client, "hello, "+client.Nick())
	c.deliverOutgoingMessage(message)
}

// ReceiveMessage receive new message, and send it to all connected clients.
func (c Chat) ReceiveMessage(message IncomingMessage) {
	c.logger.Print("recieved message '", message.Content, "' from ", message.Author.Nick())

	for _, client := range c.clients {
		c.deliverOutgoingMessage(IncomingMessageToOutgoingMessage(message, client))
	}

	c.logger.Print("message ", message, "delivered")
}

func (c Chat) deliverOutgoingMessage(message OutgoingMessage) {
	if message.Receiver.CanReceiveMessage() {
		return
	}

	c.logger.Print("delivering message ", message)
	c.outgoingMessages <- message
}

func (c *Chat) addBotClient() error {
	if c.bot != nil {
		return errors.New("bot already connected")
	}

	botClient := &Client{nick: "Bot", isBot: true}
	c.addClient(botClient)
	c.bot = botClient

	return nil
}

// OutgoingMessages returns channel containing messages, that should be delivered to logged clients.
func (c *Chat) OutgoingMessages() chan OutgoingMessage {
	return c.outgoingMessages
}

func (c Chat) deliverClientList() {
	// TODO
}
