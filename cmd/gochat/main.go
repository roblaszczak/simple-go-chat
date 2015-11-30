package main

import (
	"fmt"
	"github.com/roblaszczak/simple-go-chat/adapter"
	"github.com/roblaszczak/simple-go-chat/chat"
	"github.com/roblaszczak/simple-go-chat/config"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var logger = log.New(os.Stderr, "gochat:", log.Ldate|log.Ltime)

func main() {
	RunServer(config.SERVER_HOST, config.SERVER_PORT)
}

// TODO: consider move it to separated package
func RunServer(host string, port int) {
	handleWebsocket()
	handlePublicFiles()

	logger.Println(fmt.Sprintf("server started at http://%s:%d/", host, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

func handleWebsocket() {
	websocketController := adapter.NewWebsocketController()
	http.Handle("/chat", websocketController)

	chatCore := chat.NewChat()
	bridge := adapter.NewWebsocket(websocketController, chatCore)
	bridge.Listen()
}

func handlePublicFiles() {
	appJsFile := getPublicDir() + "/app.js"
	if _, err := os.Stat(appJsFile); err != nil {
		panic(appJsFile + " doesn't exists. You need to dump JS files using 'make buildjs' command.")
	}

	fs := http.FileServer(http.Dir(getPublicDir()))
	logger.Print(fmt.Sprintf("serving public files from %s", getPublicDir()))
	http.Handle("/", fs)
}

func getPublicDir() string {
	_, filename, _, _ := runtime.Caller(1)

	publicDir, err := filepath.Abs(filepath.Dir(filename) + "/../../" + config.PUBLIC_DIR)
	if err != nil {
		panic(err)
	}

	return publicDir
}