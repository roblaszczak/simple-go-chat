run:
	go run go_chat.go

buildjs:
	gopherjs build js/*.go --output=public/app.js

test:
	go test

buildandtest: buildjs test
