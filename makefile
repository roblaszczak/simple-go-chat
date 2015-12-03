TESTS_DIR = tests

run:
	go run cmd/gochat/main.go

buildjs:
	gopherjs build js/*.go --output=public/app.js

test:
	cd $(TESTS_DIR) && go test -v

buildandtest: buildjs test
