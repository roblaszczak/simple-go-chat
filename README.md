![GoChat](https://github.com/roblaszczak/simple-go-chat/blob/master/public/logo.png?raw=true)

# Simple Go Chat

Simple chat written in almost pure Golang. It uses websockets to comunicate with client. This is my first Golang app, so
it is still far from perfection :) Any suggestions are welcome!

All frontend's JavaScript's are written in Golang and they are dumped into JavaScript using 
[gopherjs](https://github.com/gopherjs/gopherjs). Frontend is build using AngularJS.

## How to set up?

### Standard way

    go get github.com/gopherjs/gopherjs  # installs GopherJS
    
    go get github.com/gopherjs/jquery  # installs jQuery GopherJS's bindings
    
    go get github.com/roblaszczak/simple-go-chat/cmd/gochat  # installs Simple Go chat
    
    cd "$GOPATH/src/github.com/roblaszczak/simple-go-chat"
    
    make buildjs # build JavaScript for frontend


### Docker
    
    go get github.com/roblaszczak/simple-go-chat/cmd/gochat
    
    cd "$GOPATH/src/github.com/roblaszczak/simple-go-chat/cmd/gochat"
    
    docker build -t simple-go-chat .

## How to run

### Standard way

Just execute

    gochat
    
### Docker

    docker run -p 8080:8080 -it --rm --name chat simple-go-chat

## TODO

- [ ] Config from args
- [ ] Set channels directions
- [ ] Extended readme
  - [ ] Screen (or gif) from app
  - [ ] Scheme of app
- [ ] Continious integration
- [x] Docker image
- [ ] Better support of client disconnect in controller
- [ ] Clients list
- [ ] [Frontend] Connection errors

## Contributing

I have no idea, why anyone would like to contribute it... but of course pull requests are welcome ;)

## License

Simple Go Chat is MIT-Licensed
