package chat

import (
	"math/rand"
	"time"
)

// A Client represents user connected to chat.
type Client struct {
	nick  string
	isBot bool
}

// NewClient create new instance of Client.
func NewClient(nick string) *Client {
	return &Client{nick: nick}
}

// Nick returns nickname of chat client.
func (c Client) Nick() string {
	return c.nick
}

// CanReceiveMessage return false if client cannot receive messages (for example he is bot).
func (c Client) CanReceiveMessage() bool {
	return c.isBot
}

// RandomNick generate and return random nick for Client.
func RandomNick() string {
	suffixLen := 3

	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789"
	result := make([]byte, suffixLen)
	for i := 0; i < suffixLen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return "anonymus_" + string(result)
}
