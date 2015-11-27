package chat

import "time"

// An IncomingMessage represents a message sent to chat.
type IncomingMessage struct {
	Author  *Client
	Content string
	Time    time.Time
}

// NewIncomingMessage creates new instance of IncomingMessage
func NewIncomingMessage(author *Client, content string) IncomingMessage {
	return IncomingMessage{author, content, time.Now()}
}

// An OutgoingMessage represents a message with will be send to clients connected to chat.
type OutgoingMessage struct {
	Receiver *Client
	IncomingMessage
}

// NewOutgoingMessage creates new instance of an OutgoingMessage
func NewOutgoingMessage(author, receiver *Client, content string) OutgoingMessage {
	message := NewIncomingMessage(author, content)

	return IncomingMessageToOutgoingMessage(message, receiver)
}

// IncomingMessageToOutgoingMessage converts an IncomingMessage to OutgoingMessage.
// This function is used for sending IncomingMessage to chat clients.
func IncomingMessageToOutgoingMessage(incomingMessage IncomingMessage, receiver *Client) OutgoingMessage {
	message := OutgoingMessage{}
	message.Author = incomingMessage.Author
	message.Content = incomingMessage.Content
	message.Time = incomingMessage.Time
	message.Receiver = receiver

	return message
}

// A MessageData represents message received by websocket or sent to chat client.
type MessageData struct {
	Author, Content string
	Time            time.Time
}
