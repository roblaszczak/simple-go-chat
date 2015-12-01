![GoChat](https://github.com/roblaszczak/simple-go-chat/blob/master/public/logo.png?raw=true)

# Simple Go Chat (Alpha)

Simple chat written in almost pure Golang. It uses websockets to comunicate with client. This is my first Golang app, so
it is still far from perfection :) Any suggestions are welcome!

All frontend's JavaScript's are written in Golang and they are dumped into JavaScript using 
[gopherjs](https://github.com/gopherjs/gopherjs). Frontend is build using AngularJS.

Chat is still in **alpha**, so some features are missing. More info at TODO section.

![Screen](https://github.com/roblaszczak/simple-go-chat/blob/master/docs/img/screen.png?raw=true)

## How to set up?

### Standard way

    # installs GopherJS
    go get github.com/gopherjs/gopherjs
    
    # installs jQuery GopherJS's bindings
    go get github.com/gopherjs/jquery
    
    # installs Simple Go chat
    go get github.com/roblaszczak/simple-go-chat/cmd/gochat
    
    # build JavaScript for frontend
    cd "$GOPATH/src/github.com/roblaszczak/simple-go-chat"
    make buildjs 


### Docker
    
    git clone https://github.com/roblaszczak/simple-go-chat
    cd simple-go-chat/cmd/gochat
    docker build -t simple-go-chat .

## How to run

### Standard way

Just execute

    gochat
    
### Docker

    docker run -p 8080:8080 --rm --name chat simple-go-chat

## Scheme

![Structure](https://github.com/roblaszczak/simple-go-chat/blob/master/docs/img/scheme.png?raw=true)

## TODO

- Frontend
  - [ ] Connection errors
  - [ ] Sound on message
  - [ ] Avatars support
- Chat
  - [ ] Custom nickname support
  - [ ] Connected clients list
- [ ] Config from args
- [ ] Set channels directions
- [ ] Extended readme
  - [x] Screen (or gif) from app
  - [x] Scheme of app
- [ ] Continious integration
- [x] Docker image
- [ ] Better support of client disconnect in controller
- [ ] Encrypted websocket communication

## Contributing

I have no idea, why anyone would like to contribute it... but of course pull requests are welcome ;)

## License

Simple Go Chat is MIT-Licensed
