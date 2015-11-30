package main

import (
	"fmt"
	"github.com/roblaszczak/go-chat/chat"
	"github.com/roblaszczak/go-chat/config"
	"github.com/roblaszczak/go-chat/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	RunServer(config.SERVER_HOST, config.SERVER_PORT)
}

func RunServer(host string, port int) {
	websocketController := websocket.NewWebsocketController()
	http.Handle("/chat", websocketController)

	chatCore := chat.NewChat()
	bridge := websocket.NewChatBridge(websocketController, chatCore)
	bridge.Listen()

	assertPublicDirFiles()
	fs := http.FileServer(http.Dir(config.PUBLIC_DIR))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

func assertPublicDirFiles() {
	appJsFile := config.PUBLIC_DIR + "/app.js"
	if _, err := os.Stat(appJsFile); err != nil {
		panic(appJsFile + " doesn't exists. You need to dump JS files using 'make buildjs' command.")
	}
}
