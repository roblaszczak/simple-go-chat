TESTS_DIR = tests

run:
	go run go_chat.go

buildjs:
	gopherjs build js/*.go --output=public/app.js

test:
	cd $(TESTS_DIR) && go test

buildandtest: buildjs test
