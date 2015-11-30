package main

import (
	"fmt"
	"github.com/roblaszczak/go-chat/chat"
	"github.com/roblaszczak/go-chat/config"
	"github.com/roblaszczak/go-chat/adapter"
	"log"
	"net/http"
	"os"
)

func main() {
	RunServer(config.SERVER_HOST, config.SERVER_PORT)
}

func RunServer(host string, port int) {
	websocketController := adapter.NewWebsocketController()
	http.Handle("/chat", websocketController)

	chatCore := chat.NewChat()
	bridge := adapter.NewWebsocket(websocketController, chatCore)
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
