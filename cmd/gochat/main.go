package main

import (
	"fmt"
	"github.com/roblaszczak/go-chat/adapter"
	"github.com/roblaszczak/go-chat/chat"
	"github.com/roblaszczak/go-chat/config"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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
	fs := http.FileServer(http.Dir(getPublicDir()))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

func assertPublicDirFiles() {
	appJsFile := getPublicDir() + "/app.js"
	if _, err := os.Stat(appJsFile); err != nil {
		panic(appJsFile + " doesn't exists. You need to dump JS files using 'make buildjs' command.")
	}
}

func getPublicDir() string {
	_, filename, _, _ := runtime.Caller(1)

	publicDir, err := filepath.Abs(filepath.Dir(filename) + "/../../" + config.PUBLIC_DIR)
	if err != nil {
		panic(err)
	}

	return publicDir
}
